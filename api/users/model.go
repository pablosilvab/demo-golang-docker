package users

import (
	"time"
)

// User -> struct to represent users
type User struct {
	Id            string    `json:"id"`
	Name          string    `json:"name"`
	Age           int       `json:"age"`
	InsertionDate time.Time `json:"insertion_date"`
}

// Log -> struct to write logs in Elasticsearch
type Log struct {
	App         string
	Host        string
	Method      string
	Endpoint    string
	Miliseconds int64
	Response    int
	Timestamp   time.Time
}
