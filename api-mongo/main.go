package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dimaskiddo/simple-go/api-mongo/configs"
	"github.com/dimaskiddo/simple-go/api-mongo/controllers"
	"github.com/dimaskiddo/simple-go/api-mongo/helpers"

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
	router.HandleFunc("/auth", controllers.GetAuthentication).Methods("POST")

	// Initialize Router Endpoint Secured With Basic Auth
	router.Handle("/basic-auth", helpers.AuthBasic(controllers.GetAuthentication)).Methods("GET", "POST")

	// Initialize Router Endpoint Secured With Authorization
	router.Handle("/users", helpers.AuthJWT(controllers.GetUser)).Methods("GET")
	router.Handle("/users", helpers.AuthJWT(controllers.AddUser)).Methods("POST")
	router.Handle("/users/{id}", helpers.AuthJWT(controllers.GetUserById)).Methods("GET")
	router.Handle("/users/{id}", helpers.AuthJWT(controllers.PutUserById)).Methods("PUT", "PATCH")
	router.Handle("/users/{id}", helpers.AuthJWT(controllers.DelUserById)).Methods("DELETE")

	// Set Router Handler with Logging & CORS Support
	routerHandler := handlers.LoggingHandler(os.Stdout, handlers.CORS(corsAllowedHeaders, corsAllowedOrigins, corsAllowedMethods)(router))

	// Start The HTTP Web Server
	fmt.Println("Application Serving at", configs.SvcIP+":"+configs.SvcPort)
	log.Fatal(http.ListenAndServe(configs.SvcIP+":"+configs.SvcPort, routerHandler))
}
