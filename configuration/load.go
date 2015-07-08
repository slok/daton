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

	// Load configuration
	err := viper.ReadInConfig()
	if err != nil {
		errorMessage := fmt.Sprintf("Cant load config file: %s", err)
		log.Fatal(errorMessage)
	}
	log.Info("Settings loaded")
}
