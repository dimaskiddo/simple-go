package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/dimaskiddo/simple-go/api-mysql/configs"
	"github.com/dimaskiddo/simple-go/api-mysql/controllers"
	"github.com/dimaskiddo/simple-go/api-mysql/helpers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize Channel for OS Signal
	signalExit := make(chan int)
	signalOS := make(chan os.Signal, 1)

	// Initialize Configuration
	configs.Initialize()

	// Initialize Database
	helpers.DB = helpers.MySQLConnect()
	defer helpers.DB.Close()

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

	// Catch OS Signal from Channel
	signal.Notify(signalOS, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalOS
		fmt.Println("Terminating Application Gracefully")
		helpers.DB.Close()
	}()

	// Return Exit Code
	<-signalExit
}
