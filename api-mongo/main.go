package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dimaskiddo/simple-go/api-mongo/configs"
	"github.com/dimaskiddo/simple-go/api-mongo/controllers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize Configuration
	configs.Initialize()

	// Initialize Router
	router := mux.NewRouter()

	// Set Router to Handle Endpoint
	router.HandleFunc("/", controllers.GetIndex).Methods("GET")
	router.HandleFunc("/users", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.AddUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.PutUserById).Methods("PUT", "PATCH")
	router.HandleFunc("/users/{id}", controllers.DelUserById).Methods("DELETE")

	// Add Handler For CORS
	handler := cors.Default().Handler(router)

	// Set Server Listener
	fmt.Println("Application Serving at Port", configs.SvcPort)
	log.Fatal(http.ListenAndServe(configs.SvcPort, handler))
}
