package data

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/spf13/viper"
	. "gopkg.in/check.v1"

	"github.com/slok/daton/configuration"
)

// Hook up gocheck into the "go test" runner.
func TestData(t *testing.T) { TestingT(t) }

type DataDeploymentSuite struct {
}

var _ = Suite(&DataDeploymentSuite{})

func (s *DataDeploymentSuite) TestDeploymentModel(c *C) {

	t, _ := time.Parse(time.RFC3339Nano, "2015-07-12T12:47:20.059465681Z")

	d := Deployment{
		Url:  "http://test.com/repos/slok/daton/deployments/1",
		Id:   1,
		Sha:  "d583d658d6da0b2f95ab3bcd27cd7d4bd93c3fc0",
		Ref:  "master",
		Task: "deploy",
		Payload: map[string]interface{}{
			"user":    "slok",
			"room_id": 123456,
		},
		Environment: "production",
		Description: "Deploy request from hubot",
		Creator: User{
			Login: "slok",
			Id:    1,
		},
		CreatedAt:     t,
		UpdatedAt:     t,
		StatusesUrl:   "http://test.com/repos/slok/daton/deployments/1/statuses",
		RepositoryUrl: "http://test.com/repos/slok/daton",
		StatusId:      1,
		Namespace:     "slok/daton",
	}

	jsonOk := `{"url":"http://test.com/repos/slok/daton/deployments/1","id":1,"sha":"d583d658d6da0b2f95ab3bcd27cd7d4bd93c3fc0","ref":"master","task":"deploy","payload":{"room_id":123456,"user":"slok"},"environment":"production","description":"Deploy request from hubot","creator":{"login":"slok","id":1},"created_at":"2015-07-12T12:47:20.059465681Z","updated_at":"2015-07-12T12:47:20.059465681Z","statuses_url":"http://test.com/repos/slok/daton/deployments/1/statuses","repository_url":"http://test.com/repos/slok/daton","status_id":1}`
	j, err := json.Marshal(d)
	if err != nil {
		c.Error(err)
	}

	c.Assert(string(j), Equals, jsonOk)
}

func (s *DataDeploymentSuite) TestStatusModel(c *C) {

	t, _ := time.Parse(time.RFC3339Nano, "2015-07-12T12:47:20.059465681Z")

	status := Status{
		Url:   "http://test.com/repos/slok/daton/deployments/1/statuses",
		Id:    1,
		State: StatusSuccess,
		Creator: User{
			Login: "slok",
			Id:    1,
		},
		Description:   "Deployment finished successfully.",
		TargetUrl:     "https://test.com/deployment/1/output",
		CreatedAt:     t,
		UpdatedAt:     t,
		DeploymentUrl: "http://test.com/repos/slok/daton/deployments/1",
		RepositoryUrl: "http://test.com/repos/slok/daton",
		DeploymentId:  1,
		Namespace:     "slok/daton",
	}

	jsonOk := `{"url":"http://test.com/repos/slok/daton/deployments/1/statuses","id":1,"state":"success","creator":{"login":"slok","id":1},"description":"Deployment finished successfully.","target_url":"https://test.com/deployment/1/output","created_at":"2015-07-12T12:47:20.059465681Z","updated_at":"2015-07-12T12:47:20.059465681Z","deployment_url":"http://test.com/repos/slok/daton/deployments/1","repository_url":"http://test.com/repos/slok/daton","deploy_id":1}`
	j, err := json.Marshal(status)
	if err != nil {
		c.Error(err)
	}

	c.Assert(string(j), Equals, jsonOk)
}

//-------------------------

type DatabaseModelDeploymentSuite struct {
	deployments []*Deployment
	db          *BoltDb
}

var _ = Suite(&DatabaseModelDeploymentSuite{})

func (s *DatabaseModelDeploymentSuite) SetUpSuite(c *C) {
	configuration.LoadSettingsFromFile()
	viper.Set("BoltdbName", "/tmp/datontest.db")
}

func (s *DatabaseModelDeploymentSuite) SetUpTest(c *C) {
	// Set test instances
	s.deployments = []*Deployment{
		&Deployment{
			Sha:         "d583d658d6da0b2f95ab3bcd27cd7d4bd93c3fc0",
			Ref:         "master",
			Task:        "deploy",
			Environment: "production",
			Payload: map[string]interface{}{
				"user":    "slok",
				"room_id": 123456,
			},
			Namespace: "slok/daton",
		}, &Deployment{
			Sha:         "980670afcebfd86727505b3061d8667195234816",
			Ref:         "fix-#4213",
			Task:        "deploy",
			Environment: "preproduction",
			Payload: map[string]interface{}{
				"user":    "slok",
				"room_id": 123456,
			},
			Namespace: "slok/daton",
		},
	}

	// Set database
	s.db, _ = GetBoltDb()

}

func (s *DatabaseModelDeploymentSuite) TearDownTest(c *C) {
	// Delete the config file (if present)
	s.db.Disconnect()
	err := os.RemoveAll(viper.GetString("BoltdbName"))
	if err != nil {
		panic(err)
	}
}

func (s *DatabaseModelDeploymentSuite) TestDeploymentSave(c *C) {
	for _, i := range s.deployments {
		i.Save()
	}

	// retrieve the deployments
	db.Conn.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("deployments"))
		c.Assert(10, Equals, b.Stats().KeyN)
		// 10 = 2(deploy data) + 1(production) + 1(preproduction) + 1(master) + 1(fix-#4213) + 1(deploy) + 1(d583d658d6da0b2f95ab3bcd27cd7d4bd93c3fc0) + 1(980670afcebfd86727505b3061d8667195234816) + 1 (counter)
		c.Assert("[1]", Equals, string(b.Get([]byte("slok/daton:query:d583d658d6da0b2f95ab3bcd27cd7d4bd93c3fc0"))))
		c.Assert("[2]", Equals, string(b.Get([]byte("slok/daton:query:980670afcebfd86727505b3061d8667195234816"))))

		c.Assert("[1]", Equals, string(b.Get([]byte("slok/daton:query:master"))))
		c.Assert("[2]", Equals, string(b.Get([]byte("slok/daton:query:fix-#4213"))))

		c.Assert("[1,2]", Equals, string(b.Get([]byte("slok/daton:query:deploy"))))

		c.Assert("[1]", Equals, string(b.Get([]byte("slok/daton:query:production"))))
		c.Assert("[2]", Equals, string(b.Get([]byte("slok/daton:query:preproduction"))))

		d1, _ := json.Marshal(s.deployments[0])
		d2, _ := json.Marshal(s.deployments[1])
		c.Assert(string(d1), Equals, string(b.Get([]byte("slok/daton:data:1"))))
		c.Assert(string(d2), Equals, string(b.Get([]byte("slok/daton:data:2"))))

		return nil
	})
}

func (s *DatabaseModelDeploymentSuite) TestDeploymentListAsJson(c *C) {
	for _, i := range s.deployments {
		i.Save()
	}
	ds, _ := ListDeploymentsAsJson(s.deployments[0].Namespace)

	c.Assert(len(ds), Equals, len(s.deployments))

	dj1 := Deployment{}
	dj2 := Deployment{}
	json.Unmarshal(ds[0], &dj1)
	json.Unmarshal(ds[1], &dj2)

	c.Assert(dj1.Id, Equals, int64(1))
	c.Assert(dj2.Id, Equals, int64(2))

	c.Assert(dj1.Sha, Equals, "d583d658d6da0b2f95ab3bcd27cd7d4bd93c3fc0")
	c.Assert(dj2.Sha, Equals, "980670afcebfd86727505b3061d8667195234816")

	c.Assert(dj1.Ref, Equals, "master")
	c.Assert(dj2.Ref, Equals, "fix-#4213")

	c.Assert(dj1.Environment, Equals, "production")
	c.Assert(dj2.Environment, Equals, "preproduction")

	c.Assert(dj1.Task, Equals, "deploy")
	c.Assert(dj2.Task, Equals, "deploy")
}

func (s *DatabaseModelDeploymentSuite) TestDeploymentList(c *C) {
	for _, i := range s.deployments {
		i.Save()
	}
	ds, _ := ListDeployments(s.deployments[0].Namespace)

	c.Assert(len(ds), Equals, len(s.deployments))

	c.Assert(ds[0].Id, Equals, int64(1))
	c.Assert(ds[1].Id, Equals, int64(2))

	c.Assert(ds[0].Sha, Equals, "d583d658d6da0b2f95ab3bcd27cd7d4bd93c3fc0")
	c.Assert(ds[1].Sha, Equals, "980670afcebfd86727505b3061d8667195234816")

	c.Assert(ds[0].Ref, Equals, "master")
	c.Assert(ds[1].Ref, Equals, "fix-#4213")

	c.Assert(ds[0].Environment, Equals, "production")
	c.Assert(ds[1].Environment, Equals, "preproduction")

	c.Assert(ds[0].Task, Equals, "deploy")
	c.Assert(ds[1].Task, Equals, "deploy")
}
