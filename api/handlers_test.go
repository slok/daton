package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/slok/daton/configuration"

	. "gopkg.in/check.v1"
)

func TestApiHandlers(t *testing.T) { TestingT(t) }

type ApiHandlersTestSuite struct {
	apiPrefix string
}

var _ = Suite(&ApiHandlersTestSuite{})

func (s *ApiHandlersTestSuite) SetUpSuite(c *C) {
	s.apiPrefix = fmt.Sprintf("/api/v%d", configuration.ApiVersion)
}

func (s *ApiHandlersTestSuite) TestPingApiHandler(c *C) {
	request, _ := http.NewRequest("GET", path.Join(s.apiPrefix, "ping"), nil)
	response := httptest.NewRecorder()

	PingHandler(response, request)
	c.Assert(response.Code, Equals, http.StatusOK)
	body := response.Body.String()

	c.Assert(strings.HasPrefix(body, "{\"status\": \"ok\", \"timestamp\":"),
		Equals, true)
}
