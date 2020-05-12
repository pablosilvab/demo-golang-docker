package api

import (
	"github.com/gorilla/mux"
	"github.com/pablosilvab/demo-golang-docker/healthcheck"
)

// InitRouter : Return a router instance with paths.
func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/user", getUser).Methods("GET")
	r.HandleFunc("/api/user/last", getLastUser).Methods("GET")
	r.HandleFunc("/health", healthcheck.HealthHandler).Methods("GET")
	r.HandleFunc("/ready", readinessHandler).Methods("GET")
	return r
}
