package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pablosilvab/demo-golang-docker/api/users"
	"github.com/pablosilvab/demo-golang-docker/healthcheck"
)

// InitRouter : Return a router instance with paths.
func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/api/users", users.GetUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}", getUsers).Methods("GET")
	r.HandleFunc("/health", healthcheck.HealthHandler).Methods("GET")
	r.HandleFunc("/ready", readinessHandler).Methods("GET")
	return r
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	users.GetUser(w, r, userId)
}
