package users

import "time"

// User -> struct to represent users
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Client -> struct to represent clients
type Client struct {
	Person  User `json:"user"`
	Enabled bool `json:"enabled"`
}

// Log -> struct to write logs in Elasticsearch
type Log struct {
	Endpoint  string
	Timestamp time.Time
}
