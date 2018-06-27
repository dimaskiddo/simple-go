package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dimaskiddo/simple-go/api-mongo/configs"
	"github.com/dimaskiddo/simple-go/api-mongo/controllers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize Configuration
	configs.Initialize()

	// Initialize CORS Configurataion
	corsAllowedHeaders := handlers.AllowedHeaders(configs.CORSAllowedHeaders)
	corsAllowedOrigins := handlers.AllowedHeaders(configs.CORSAllowedOrigins)
	corsAllowedMethods := handlers.AllowedHeaders(configs.CORSAllowedMethods)

	// Initialize Router
	router := mux.NewRouter()

	// Initialize Router Endpoint
	router.HandleFunc("/", controllers.GetIndex).Methods("GET")
	router.HandleFunc("/users", controllers.GetUser).Methods("GET")
	router.HandleFunc("/users", controllers.AddUser).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.PutUserById).Methods("PUT", "PATCH")
	router.HandleFunc("/users/{id}", controllers.DelUserById).Methods("DELETE")

	// Set Router Handler with Logging & CORS Support
	routerHandler := handlers.LoggingHandler(os.Stdout, handlers.CORS(corsAllowedHeaders, corsAllowedOrigins, corsAllowedMethods)(router))

	// Start The HTTP Web Server
	fmt.Println("Application Serving at Port", configs.SvcPort)
	log.Fatal(http.ListenAndServe(configs.SvcPort, routerHandler))
}
