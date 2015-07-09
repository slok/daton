package api

import (
	"fmt"
	"net/http"
	"time"
)

// Return ping
func PingHandler(w http.ResponseWriter, r *http.Request) {
	//body
	m := fmt.Sprintf("{\"status\": \"ok\", \"timestamp\": %d}",
		time.Now().Unix())
	fmt.Fprintf(w, m)
	// Headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
