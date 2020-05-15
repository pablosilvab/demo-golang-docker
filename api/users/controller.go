package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pablosilvab/elastic-lib"
)

// AppName -> index to Elasticsearch
const AppName = "demo-golang-docker"

// GetClients return json clients
func GetClients(w http.ResponseWriter, r *http.Request) {
	go elastic.Log(AppName, Log{r.RequestURI, time.Now()})
	// Dummy users
	clients := []Client{
		{
			Person: User{Name: "Pablo", Age: 26}, Enabled: true,
		},
		{
			Person: User{Name: "Maria", Age: 26}, Enabled: false,
		},
	}
	log.Println(clients)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(clients)
}

// GetUsers return json users
func GetUsers(w http.ResponseWriter, r *http.Request) {

	// TODO: create cloud queue
	// go elastic.Log(AppName, Log{r.RequestURI, time.Now()})

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
	_ = json.NewEncoder(w).Encode(users)
}

// GetUser return user data in a encoding data example
func GetUser(w http.ResponseWriter, r *http.Request) {

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
	_, _ = w.Write(userData)
}

// GetLastUser return encoding data. also there is an example of decoding.
func GetLastUser(w http.ResponseWriter, r *http.Request) {
	// Decoding data
	var user User
	b := []byte(`{"Name":"Pablo","Age":26}`)
	log.Printf("Encoding Data: %v", b)

	err := json.Unmarshal(b, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Printf("Decoding Data: %v", user)
	_, _ = w.Write(b)
}
