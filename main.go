package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/spf13/viper"

	"github.com/slok/daton/api"
	"github.com/slok/daton/configuration"
	"github.com/slok/daton/data"
)

func main() {
	log.Info("Starting Daton...")

	// Load configuration
	configuration.LoadSettingsFromFile()

	// Init database
	db, err := data.GetBoltDb()
	defer db.Disconnect()
	if err != nil {
		log.Panic("Couldn't connect to bolt database")
	}

	// Bind routing with handlers
	router := api.BindApiRoutes(nil)

	// serve
	n := negroni.Classic()
	n.UseHandler(router)

	listenAddress := fmt.Sprintf(":%d", viper.GetInt("Port"))
	n.Run(listenAddress)
}
