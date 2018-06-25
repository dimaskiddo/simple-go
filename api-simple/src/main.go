package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ResponseGetAll struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Users
}

type ResponseGet struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Users
}

type ResponseGlobal struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func main() {
	// Initialize Router
	router := mux.NewRouter()

	// Set Router to Handle Endpoint
	router.HandleFunc("/users", returnUserAll).Methods("GET")
	router.HandleFunc("/users", returnUserAdd).Methods("POST")
	router.HandleFunc("/users/{id}", returnUserOne).Methods("GET")
	router.HandleFunc("/users/{id}", returnUserUpdate).Methods("PUT", "PATCH")
	router.HandleFunc("/users/{id}", returnUserDelete).Methods("DELETE")

	// Set Server Listener
	fmt.Println("Serve and Listen at Port :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
