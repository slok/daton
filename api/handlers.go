package api

import (
	"encoding/json"
	"net/http"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

func logHandler(r *http.Request, handler string, params map[string]string, queryParams map[string][]string) {
	log.WithFields(log.Fields{
		"url":     r.RequestURI,
		"handler": handler,
		"method":  r.Method,
		"params":  params,
		"Query":   queryParams,
	}).Debug("Calling handler")
}

// Return ping
func PingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//body
	data := map[string]interface{}{
		"status":    "ok",
		"timestamp": time.Now().UTC().Unix(),
	}
	b, _ := json.Marshal(data)
	w.Write(b)

	logHandler(r, "pingHandler", params, nil)
}

func DeployListHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logHandler(r, "DeployListHandler", params, r.URL.Query())
}
func DeploycreateHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logHandler(r, "DeploycreateHandler", params, nil)
}
func StatusListHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logHandler(r, "StatusListHandler", params, nil)
}
func StatusCreatePingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logHandler(r, "StatusCreatePingHandler", params, nil)
}
