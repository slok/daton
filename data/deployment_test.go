package data

import (
	"encoding/json"
	"testing"
	"time"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestDataDeployment(t *testing.T) { TestingT(t) }

type DataDeploymentSuite struct {
}

var _ = Suite(&DataDeploymentSuite{})

func (s *DataDeploymentSuite) testDeploymentModel(c *C) {

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
	}

	jsonOk := `{"url":"http://test.com/repos/slok/daton/deployments/1","id":1,"sha":"d583d658d6da0b2f95ab3bcd27cd7d4bd93c3fc0","ref":"master","task":"deploy","payload":{"room_id":123456,"user":"slok"},"environment":"production","description":"Deploy request from hubot","creator":{"login":"slok","id":1},"created_at":"2015-07-12T12:47:20.059465681Z","updated_at":"2015-07-12T12:47:20.059465681Z","statuses_url":"http://test.com/repos/slok/daton/deployments/1/statuses","repository_url":"http://test.com/repos/slok/daton","status_id":1}`
	j, err := json.Marshal(d)
	if err != nil {
		c.Error(err)
	}

	c.Assert(string(j), Equals, jsonOk)
}

func (s *DataDeploymentSuite) testStatusModel(c *C) {

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
	}

	jsonOk := `{"url":"http://test.com/repos/slok/daton/deployments/1/statuses","id":1,"state":"success","creator":{"login":"slok","id":1},"description":"Deployment finished successfully.","target_url":"https://test.com/deployment/1/output","created_at":"2015-07-12T12:47:20.059465681Z","updated_at":"2015-07-12T12:47:20.059465681Z","deployment_url":"http://test.com/repos/slok/daton/deployments/1","repository_url":"http://test.com/repos/slok/daton","deploy_id":1}`
	j, err := json.Marshal(status)
	if err != nil {
		c.Error(err)
	}

	c.Assert(string(j), Equals, jsonOk)
}
