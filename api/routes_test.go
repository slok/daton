package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	. "gopkg.in/check.v1"

	"github.com/slok/daton/configuration"
)

func TestApiRoutes(t *testing.T) { TestingT(t) }

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

	c.Assert(strings.HasPrefix(body, "{\"status\": \"ok\", \"timestamp\":"),
		Equals, true)
}
