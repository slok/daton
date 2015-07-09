package api

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/slok/daton/configuration"
)

var (
	apiVersion = configuration.ApiVersion
)

func BindApiRoutes(router *mux.Router) *mux.Router {
	prefix := fmt.Sprintf("/api/v%d", apiVersion)

	// Create the router if not new
	if router == nil {
		router = mux.NewRouter()
	}

	s := router.PathPrefix(prefix).Subrouter()
	s.HandleFunc("/ping", PingHandler).Methods("GET")

	return router
}
