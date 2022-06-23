package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/pablosilvab/elastic"
)

// AppName -> index to Elasticsearch
const AppName = "demo-golang-docker"

type allUsers []User

var users = allUsers{
	{
		Id:            "1",
		Name:          "Pablo",
		Age:           26,
		InsertionDate: time.Now(),
	},
	{
		Id:            "2",
		Name:          "Maria",
		Age:           26,
		InsertionDate: time.Now(),
	},
	{
		Id:            "3",
		Name:          "Tony",
		Age:           46,
		InsertionDate: time.Now(),
	},
}

// GetUsers return json users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	initTime := time.Now()

	elastic.Log(AppName, Log{
		AppName,
		r.RemoteAddr,
		r.RequestURI,
		r.Method,
		time.Since(initTime).Milliseconds(),
		http.StatusOK,
		time.Now()})

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users)

}

// GetUser return user data in a encoding data example
func GetUser(w http.ResponseWriter, r *http.Request, id string) {
	initTime := time.Now()

	w.Header().Set("Content-Type", "application/json")

	userId := id

	for _, user := range users {
		if user.Id == userId {
			json.NewEncoder(w).Encode(user)

			elastic.Log(AppName, Log{
				AppName,
				r.RemoteAddr,
				r.RequestURI,
				r.Method,
				time.Since(initTime).Milliseconds(),
				http.StatusOK,
				time.Now()})
			return
		}
	}

	elastic.Log(AppName, Log{
		AppName,
		r.RemoteAddr,
		r.RequestURI,
		r.Method,
		time.Since(initTime).Milliseconds(),
		http.StatusNotFound,
		time.Now()})

	sendResponse(w, "User Not Found", http.StatusNotFound)
}

func sendResponse(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResp)
}
