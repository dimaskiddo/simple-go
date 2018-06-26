package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Function to Get All User Data
func returnUserAll(w http.ResponseWriter, r *http.Request) {
	var user User
	var users []User
	var response ResGetUser

	// Connection Handle to Database
	db := connectMySQL()
	defer db.Close()

	// Database Query
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Print(err)
	}

	// Populate Data
	for rows.Next() {
		// Match / Binding Database Field with Struct
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Fatal(err.Error())
		} else {
			// Append User Struct to Users Array of Struct
			users = append(users, user)
		}
	}

	// Set Response Data
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = users

	// Set Response Data to HTTP
	resWrite(w, response.Status, response)
}

// Function to Get One User By ID
func returnUserOne(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID is Empty
	if len(params["id"]) == 0 {
		resBadRequest(w)
	} else {
		// Get ID Parameters From URI Then Convert it to Integer
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var user User
			var users []User
			var response ResGetUser

			// Connection Handle to Database
			db := connectMySQL()
			defer db.Close()

			// Database Query
			rows, err := db.Query("SELECT * FROM users WHERE id=? LIMIT 1", userID)
			if err != nil {
				log.Print(err)
			}

			// Populate Data
			for rows.Next() {
				// Match / Binding Database Field with Struct
				if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
					log.Fatal(err.Error())
				} else {
					// Append User Struct to Users Array of Struct
					users = append(users, user)
				}
			}

			// Set Response Data
			response.Status = http.StatusOK
			response.Message = "Success"
			response.Data = users

			// Set Response Data to HTTP
			resWrite(w, response.Status, response)
		} else {
			resInternalError(w)
		}
	}
}

// Function to Add User Data
func returnUserAdd(w http.ResponseWriter, r *http.Request) {
	var user User
	var response ResGlobal

	// Decode JSON from Request Body to User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Connection Handle to Database
	db := connectMySQL()
	defer db.Close()

	// Database Query
	_, err := db.Exec("INSERT INTO users (name, email) VALUE (?, ?)", user.Name, user.Email)
	if err != nil {
		log.Print(err)
	}

	// Set Response Data
	response.Status = http.StatusCreated
	response.Message = "Success, User Created"

	// Set Response Data to HTTP
	resWrite(w, response.Status, response)

	// Write Log
	fmt.Println("User Created ID:", user.ID)
}

// Function to Update User Data By ID
func returnUserUpdate(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID is Empty
	if len(params["id"]) == 0 {
		resBadRequest(w)
	} else {
		// Get ID Parameters From URI Then Convert it to Integer
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var user User
			var response ResGlobal

			// Decode JSON from Request Body to User Data
			// Use _ As Temporary Variable
			_ = json.NewDecoder(r.Body).Decode(&user)

			// Connection Handle to Database
			db := connectMySQL()
			defer db.Close()

			// Database Query
			_, err := db.Exec("UPDATE users SET name=?, email=? WHERE id=? LIMIT 1", user.Name, user.Email, userID)
			if err != nil {
				log.Print(err)
			}

			// Set Response Data
			response.Status = http.StatusCreated
			response.Message = "Success, User Updated"

			// Set Response Data to HTTP
			resWrite(w, response.Status, response)

			// Write Log
			fmt.Println("User Updated ID:", userID)
		} else {
			resInternalError(w)
		}
	}
}

// Function to Update User Data By ID
func returnUserDelete(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID is Empty
	if len(params["id"]) == 0 {
		resBadRequest(w)
	} else {
		// Get ID Parameters From URI Then Convert it to Integer
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var response ResGlobal

			// Connection Handle to Database
			db := connectMySQL()
			defer db.Close()

			// Database Query
			_, err := db.Query("DELETE FROM users WHERE id=? LIMIT 1", userID)
			if err != nil {
				log.Print(err)
			}

			// Set Response Data
			response.Status = http.StatusOK
			response.Message = "Success, User Deleted"

			// Set Response Data to HTTP
			resWrite(w, response.Status, response)

			// Write Log
			fmt.Println("User Deleted ID:", userID)
		} else {
			resInternalError(w)
		}
	}
}
