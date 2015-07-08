package configuration

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/spf13/viper"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ConfigTestSuite struct {
	configName    string
	configPath    string
	configAbsPath string
}

var _ = Suite(&ConfigTestSuite{})

func (s *ConfigTestSuite) SetUpSuite(c *C) {
	// Create the config file
	s.configName = "daton.json"
	s.configAbsPath = "/tmp/datontest"
	s.configPath = path.Join(s.configAbsPath, s.configName)

	err := os.Mkdir(s.configAbsPath, 0744)
	if err != nil {
		panic(err)
	}

	// set config location to viper
	viper.AddConfigPath(s.configAbsPath)

	configData := []byte("{}")
	err = ioutil.WriteFile(s.configPath, configData, 0644)
	if err != nil {
		panic(err)
	}
}

func (s *ConfigTestSuite) TearDownSuite(c *C) {
	// Delete the config file
	err := os.RemoveAll(s.configAbsPath)
	if err != nil {
		panic(err)
	}
}

func (s *ConfigTestSuite) TestLoadFromFileExists(c *C) {
	LoadSettingsFromFile()
}
