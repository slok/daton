package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/slok/daton/data"
	"github.com/slok/daton/utils"
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

func writeJson(w http.ResponseWriter, data []byte, headers map[string]string, status int) (int, error) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return w.Write(data)
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
	writeJson(w, b, nil, http.StatusOK)

	logHandler(r, "pingHandler", params, nil)
}

func DeployListHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logHandler(r, "DeployListHandler", params, r.URL.Query())
	namespace := strings.Join([]string{params["owner"], params["repo"]}, "/")

	deployments, err := data.ListDeployments(namespace)
	if err != nil {
		log.Error("Error processing deployments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(deployments)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Error Marshalling json")
		return
	}

	writeJson(w, b, nil, http.StatusOK)
}
func DeployCreateHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	namespace := strings.Join([]string{params["owner"], params["repo"]}, "/")
	logHandler(r, "DeploycreateHandler", params, nil)

	// Get POST data
	decoder := json.NewDecoder(r.Body)
	d := data.Deployment{}
	err := decoder.Decode(&d)

	// Error on the json we received
	if err != nil {
		log.Errorf("Error decoding json: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid json"}`))
		return
	}

	// Check required params are present
	if len(d.Ref) <= 0 {
		log.Warning("Can't create deployment, ref missing")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Ref field is missing"}`))
		return
	}

	// Set basic data
	d.CreatedAt = time.Now().UTC()
	d.UpdatedAt = d.CreatedAt
	d.Namespace = namespace
	d.Save()

	// Set Urls for the response
	d.Url = utils.ApiUrl(fmt.Sprintf("repos/%s/deployments/%d", namespace, d.Id))
	d.StatusesUrl = fmt.Sprintf("%s/statuses", d.Url)
	//d.RepositoryUrl

	// Write final response
	b, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error("Error Marshalling json")
		return
	}

	//w.Header().Set("Location", d.Url)
	writeJson(w, b, map[string]string{"Location": d.Url}, http.StatusCreated)
}
func StatusListHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logHandler(r, "StatusListHandler", params, nil)
	w.WriteHeader(http.StatusNotImplemented)
}
func StatusCreateHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	logHandler(r, "StatusCreateHandler", params, nil)
	w.WriteHeader(http.StatusNotImplemented)
}
