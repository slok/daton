package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/spf13/viper"

	"github.com/slok/daton/api"
	"github.com/slok/daton/configuration"
)

func main() {
	log.Info("Starting Daton...")

	// Load configuration
	configuration.LoadSettingsFromFile()

	// Bind routing with handlers
	router := api.BindApiRoutes(nil)

	// serve
	n := negroni.Classic()
	n.UseHandler(router)

	listenAddress := fmt.Sprintf(":%d", viper.GetInt("Port"))
	n.Run(listenAddress)
}
