package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	. "gopkg.in/check.v1"

	"github.com/slok/daton/configuration"
	"github.com/slok/daton/data"
)

// Hook up gocheck into the "go test" runner.
func TestApi(t *testing.T) { TestingT(t) }

type ApiRoutesTestSuite struct {
	apiPrefix string
	router    *mux.Router
	response  *httptest.ResponseRecorder
}

var _ = Suite(&ApiRoutesTestSuite{})

func (s *ApiRoutesTestSuite) SetUpSuite(c *C) {
	s.apiPrefix = fmt.Sprintf("/api/v%d", configuration.ApiVersion)
}

func (s *ApiRoutesTestSuite) SetUpTest(c *C) {
	s.router = BindApiRoutes(nil)
	s.response = httptest.NewRecorder()
}

func (s *ApiRoutesTestSuite) TestNewRouter(c *C) {
	if s.router == nil {
		c.Error("Expected a valid mux instance")
	}

}

func (s *ApiRoutesTestSuite) TestApiGetBind(c *C) {
	url := path.Join(s.apiPrefix, "ping")
	method := "GET"

	request, err := http.NewRequest(method, url, nil)

	if err != nil {
		c.Fatal(err)
	}

	// Make the call
	s.router.ServeHTTP(s.response, request)

	c.Assert(s.response.Code, Equals, http.StatusOK)
	body := s.response.Body.String()

	c.Assert(strings.HasPrefix(body, `{"status":"ok","timestamp":`),
		Equals, true)
}

//----------------------------------------------------------------------------

type DeploymentsApiRoutesTestSuite struct {
	apiPrefix string
	router    *mux.Router
	response  *httptest.ResponseRecorder

	namespace   string
	deployments []*data.Deployment
}

var _ = Suite(&DeploymentsApiRoutesTestSuite{})

func (s *DeploymentsApiRoutesTestSuite) SetUpSuite(c *C) {
	s.apiPrefix = fmt.Sprintf("/api/v%d", configuration.ApiVersion)
	s.namespace = "slok/daton"

	// Set configuration
	configuration.LoadSettingsFromFile()
	viper.Set("BoltdbName", "/tmp/datontest.db")
}

func (s *DeploymentsApiRoutesTestSuite) SetUpTest(c *C) {
	s.router = BindApiRoutes(nil)
	s.response = httptest.NewRecorder()

	s.deployments = []*data.Deployment{
		&data.Deployment{
			Sha:         "d583d658d6da0b2f95ab3bcd27cd7d4bd93c3fc0",
			Ref:         "master",
			Task:        "deploy",
			Environment: "production",
			Payload: map[string]interface{}{
				"room_id": 123456,
				"user":    "slok",
			},
			Namespace: "slok/daton",
		}, &data.Deployment{
			Sha:         "980670afcebfd86727505b3061d8667195234816",
			Ref:         "fix-#4213",
			Task:        "deploy",
			Environment: "preproduction",
			Payload: map[string]interface{}{
				"room_id": 123456,
				"user":    "slok",
			},
			Namespace: "slok/daton",
		},
	}
}

func (s *DeploymentsApiRoutesTestSuite) TearDownTest(c *C) {
	db, _ := data.GetBoltDb()
	db.Disconnect()

	// Delete the database (if present)
	err := os.RemoveAll(viper.GetString("BoltdbName"))
	if err != nil {
		panic(err)
	}
}

func (s *DeploymentsApiRoutesTestSuite) TestListDeployments(c *C) {
	for _, i := range s.deployments {
		i.Save()
	}

	url := path.Join(s.apiPrefix, "repos", s.namespace, "deployments")
	method := "GET"

	request, err := http.NewRequest(method, url, nil)

	if err != nil {
		c.Fatal(err)
	}

	// Make the call
	s.router.ServeHTTP(s.response, request)

	c.Assert(s.response.Code, Equals, http.StatusOK)
	body := s.response.Body.Bytes()

	deploys := []data.Deployment{}
	json.Unmarshal(body, &deploys)

	c.Assert(len(deploys), Equals, len(s.deployments))

	for k, i := range s.deployments {
		c.Assert(deploys[k].Sha, Equals, i.Sha)
		c.Assert(deploys[k].Ref, Equals, i.Ref)
		c.Assert(deploys[k].Task, Equals, i.Task)
		c.Assert(deploys[k].Environment, Equals, i.Environment)
	}
}

func (s *DeploymentsApiRoutesTestSuite) TestCreateDeployments(c *C) {
	url := path.Join(s.apiPrefix, "repos", s.namespace, "deployments")
	method := "POST"

	for _, i := range s.deployments {
		s.response = httptest.NewRecorder()
		d, _ := json.Marshal(i)
		request, err := http.NewRequest(method, url, bytes.NewBuffer(d))

		if err != nil {
			c.Fatal(err)
		}

		// Make the call
		s.router.ServeHTTP(s.response, request)
		c.Assert(s.response.Code, Equals, http.StatusCreated)
		body := s.response.Body.Bytes()

		deploy := data.Deployment{}
		json.Unmarshal(body, &deploy)

		c.Assert(deploy.Sha, Equals, i.Sha)
		c.Assert(deploy.Ref, Equals, i.Ref)
		c.Assert(deploy.Task, Equals, i.Task)
		c.Assert(deploy.Environment, Equals, i.Environment)
	}

	deps, _ := data.ListDeployments(s.deployments[0].Namespace)
	c.Assert(len(deps), Equals, len(s.deployments))

}

func (s *DeploymentsApiRoutesTestSuite) TestCreateDeploymentsRefMissing(c *C) {
	url := path.Join(s.apiPrefix, "repos", s.namespace, "deployments")
	method := "POST"

	request, err := http.NewRequest(method, url, bytes.NewBufferString("{}"))

	if err != nil {
		c.Fatal(err)
	}

	// Make the call
	s.router.ServeHTTP(s.response, request)
	c.Assert(s.response.Code, Equals, http.StatusBadRequest)

	body := s.response.Body.Bytes()
	returnError := map[string]string{}
	json.Unmarshal(body, &returnError)
	c.Assert(returnError["error"], Equals, "Ref field is missing")
}

func (s *DeploymentsApiRoutesTestSuite) TestCreateDeploymentsWrongJson(c *C) {
	url := path.Join(s.apiPrefix, "repos", s.namespace, "deployments")
	method := "POST"

	request, err := http.NewRequest(method, url, bytes.NewBufferString("worng"))

	if err != nil {
		c.Fatal(err)
	}

	// Make the call
	s.router.ServeHTTP(s.response, request)
	c.Assert(s.response.Code, Equals, http.StatusBadRequest)

	body := s.response.Body.Bytes()
	returnError := map[string]string{}
	json.Unmarshal(body, &returnError)
	c.Assert(returnError["error"], Equals, "Invalid json")
}
