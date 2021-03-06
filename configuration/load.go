package configuration

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// Configuration keys
// 		- BoltdbName
// 		- Debug
// 		- DefaultEnvironment
// 		- DefaultTask
// 		- EnableAutomerge
// 		- Port
// 		- AppHost
//--------------------------------

func setConfigLocation() {
	viper.SetConfigName(ConfigName)
	for _, i := range ConfigPaths {
		viper.AddConfigPath(i)
	}
}

func LoadSettingsFromFile() {
	viper.SetConfigType(ConfigType)
	setConfigLocation()

	// Set defaults
	viper.SetDefault("Port", Port)
	viper.SetDefault("EnableAutomerge", EnableAutomerge)
	viper.SetDefault("Debug", Debug)
	viper.SetDefault("DefaultEnvironment", DefaultEnvironment)
	viper.SetDefault("DefaultTask", DefaultTask)
	viper.SetDefault("BoltdbName", BoltdbName)

	// Load configuration
	err := viper.ReadInConfig()
	if err != nil {
		log.Warning("Configuration not found, loading defaults")
	}

	// set log level
	if viper.GetBool("Debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Debugf("Loaded setting values: %v", viper.AllSettings())
	log.Info("Settings loaded")
}
