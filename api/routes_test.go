package api

import (
	"fmt"
	"testing"

	"github.com/slok/daton/configuration"

	. "gopkg.in/check.v1"
)

func TestApiRoutes(t *testing.T) { TestingT(t) }

type ApiRoutesTestSuite struct {
	apiPrefix string
}

var _ = Suite(&ApiRoutesTestSuite{})

func (s *ApiRoutesTestSuite) SetUpSuite(c *C) {
	s.apiPrefix = fmt.Sprintf("/api/v%d", configuration.ApiVersion)
}

func (s *ApiRoutesTestSuite) TestApiBind(c *C) {
	c.Skip("Not implemented")
}
