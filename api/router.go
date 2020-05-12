package api

import "github.com/gorilla/mux"

func InitRouter() *mux.Router {
	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/health", HealthHandler).Methods("GET")
	r.HandleFunc("/ready", readinessHandler).Methods("GET")
	return r
}
