package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Function to Get All User Data
func returnUserAll(w http.ResponseWriter, r *http.Request) {
	var response ResGetUser

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
			var user []User
			var response ResGetUser

			// Convert Users Array to Single Data User Array
			user = append(user, users[userID-1])

			// Set Response Data
			response.Status = http.StatusOK
			response.Message = "Success"
			response.Data = user

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

	// Set User ID to Current Users Array Length + 1
	user.ID = len(users) + 1

	// Insert User Data to Users Array
	users = append(users, user)

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

			// Update User Data to Users Array
			users[userID-1].Name = user.Name
			users[userID-1].Email = user.Email

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

			// Delete User Data from Users Array
			users = append(users[:userID-1], users[userID:]...)

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
