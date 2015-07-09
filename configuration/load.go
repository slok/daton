package configuration

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

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

	// Load configuration
	err := viper.ReadInConfig()
	if err != nil {
		errorMessage := fmt.Sprintf("Configuration not found, loading defaults")
		log.Warning(errorMessage)
	}

	// set log level
	if viper.GetBool("Debug") {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.Info("Settings loaded")
}
