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

	// Handy routes
	s.HandleFunc("/ping", PingHandler).Methods("GET")

	// Deploy routes
	s.HandleFunc("/repos/{owner}/{repo}/deployments", DeployListHandler).Methods("GET")
	s.HandleFunc("/repos/{owner}/{repo}/deployments", DeployCreateHandler).Methods("POST")

	// Status routes
	s.HandleFunc("/repos/{owner}/{repo}/deployments/{id}/statuses", StatusListHandler).Methods("GET")
	s.HandleFunc("/repos/{owner}/{repo}/deployments/{id}/statuses", StatusCreateHandler).Methods("POST")

	return router
}
