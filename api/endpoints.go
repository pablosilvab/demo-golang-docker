package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pablosilvab/elastic-lib"
)

const APP_NAME = "demo-golang-docker"

// User -> struct to represent users
type User struct {
	Name string
	Age  int
}

// Log -> struct to write logs in Elasticsearch
type Log struct {
	Endpoint  string
	Timestamp time.Time
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received request for %s\n", name)

	_, err := w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))

	if err != nil {
		os.Exit(0)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	go elastic.Log(APP_NAME, Log{r.RequestURI, time.Now()})
	// Dummy users
	users := []User{
		{
			Name: "Pablo",
			Age:  26,
		},
		{
			Name: "Maria",
			Age:  26,
		},
	}
	log.Println(users)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {

	// Dummy user
	user := User{
		Name: "Pablo",
		Age:  26,
	}
	log.Printf("Decoding Data: %v", user)

	// Encoding data
	userData, _ := json.Marshal(user)
	log.Printf("Encoding Data: %v", userData)

	w.Header().Set("Content-Type", "application/json")
	w.Write(userData)
}

func getLastUser(w http.ResponseWriter, r *http.Request) {
	// Decoding data
	var user User
	b := []byte(`{"Name":"Pablo","Age":26}`)
	log.Printf("Encoding Data: %v", b)

	err := json.Unmarshal(b, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Printf("Decoding Data: %v", user)
	w.Write(b)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
