package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var servPort string = ":3000"

func main() {
	// Initialize Router
	router := mux.NewRouter()

	// Set Router to Handle Endpoint
	router.HandleFunc("/users", returnUserAll).Methods("GET")
	router.HandleFunc("/users", returnUserAdd).Methods("POST")
	router.HandleFunc("/users/{id}", returnUserOne).Methods("GET")
	router.HandleFunc("/users/{id}", returnUserUpdate).Methods("PUT", "PATCH")
	router.HandleFunc("/users/{id}", returnUserDelete).Methods("DELETE")

	// Add Handler For CORS
	handler := cors.Default().Handler(router)

	// Set Server Listener
	fmt.Println("Application Running and Serving at Port", servPort)
	log.Fatal(http.ListenAndServe(servPort, handler))
}
