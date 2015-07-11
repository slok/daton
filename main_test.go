package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/spf13/viper"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestMain(t *testing.T) { TestingT(t) }

type MainSuite struct {
	url string
}

var _ = Suite(&MainSuite{})

func (s *MainSuite) SetUpSuite(c *C) {
	// Overwrite settings for the tests
	viper.Set("Port", 7357)
	s.url = fmt.Sprintf("http://127.0.0.1:%d", viper.GetInt("Port"))
}

func (s *MainSuite) TestServerStart(c *C) {
	// Call the server
	go main()

	// Wait to start
	time.Sleep(time.Second)

	// Check server is up
	_, err := http.Get(s.url)

	if err != nil {
		c.Errorf("Server should be responding...: %v", err)
	}
}
