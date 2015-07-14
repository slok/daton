package configuration

import (
	"os"
	"path"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	. "gopkg.in/check.v1"

	"github.com/slok/daton/utils"
)

func TestConfigurationLoad(t *testing.T) { TestingT(t) }

type ConfigTestSuite struct {
	configPath string
}

var _ = Suite(&ConfigTestSuite{})

func (s *ConfigTestSuite) SetUpTest(c *C) {
	viper.Reset()
	// prepare config file
	s.configPath = "/tmp/daton-test/daton.json"
	viper.AddConfigPath(path.Dir(s.configPath))
}

func (s *ConfigTestSuite) TearDownTest(c *C) {
	// Delete the config file (if present)
	err := os.RemoveAll(path.Dir(s.configPath))
	if err != nil {
		panic(err)
	}
}

func (s *ConfigTestSuite) TestLoadFromFileExists(c *C) {
	utils.WriteStringFile([]byte("{}"), s.configPath)
	LoadSettingsFromFile()
}

func (s *ConfigTestSuite) TestLoadWithoutFile(c *C) {
	LoadSettingsFromFile()
}

func (s *ConfigTestSuite) TestLoadDefaults(c *C) {
	LoadSettingsFromFile()

	c.Assert(viper.GetInt("Port"), Equals, Port)
	c.Assert(viper.GetBool("EnableAutomerge"), Equals, EnableAutomerge)
	c.Assert(viper.GetBool("Debug"), Equals, Debug)
}

func (s *ConfigTestSuite) TestDebugMode(c *C) {
	type AnonConfig struct {
		Debug bool `json:"Debug"`
	}
	utils.WriteJsonFile(AnonConfig{Debug: true}, s.configPath)
	LoadSettingsFromFile()
	c.Assert(log.GetLevel(), Equals, log.DebugLevel)

	utils.WriteJsonFile(AnonConfig{Debug: false}, s.configPath)
	// Not need to reset because the conf file exists and loading it will overwrite
	LoadSettingsFromFile()
	c.Assert(log.GetLevel(), Equals, log.InfoLevel)
}
