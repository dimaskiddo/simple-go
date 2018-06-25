package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Function to Get All User Data
func returnUserAll(w http.ResponseWriter, r *http.Request) {
	var response ResponseGetAll

	// Set Response Data
	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = UsersData

	// Set Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Function to Add User Data
func returnUserAdd(w http.ResponseWriter, r *http.Request) {
	var user Users
	var response ResponseGet

	// Decode JSON User Data
	// Use _ As Temporary Variable
	_ = json.NewDecoder(r.Body).Decode(&user)

	// Get Latest User ID
	user.ID = len(UsersData) + 1

	// Insert User to UsersData
	UsersData = append(UsersData, user)

	// Set Response Data
	response.Status = http.StatusCreated
	response.Message = "Success"
	response.Data = UsersData[len(UsersData)-1]

	// Set Response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Function to Get One User By ID
func returnUserOne(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID Is Empty
	if len(params["id"]) == 0 {
		var response ResponseGlobal

		// Set Response Data
		response.Status = http.StatusBadRequest
		response.Message = "Error, ID Not Found"

		// Set Response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		// Get User ID As Integer Value
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var response ResponseGet

			// Set Response Data
			response.Status = http.StatusOK
			response.Message = "Success, User Created"
			response.Data = UsersData[userID-1]

			// Set Response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			var response ResponseGlobal

			// Set Response Data
			response.Status = http.StatusInternalServerError
			response.Message = "Error, Coverting ID to Int"

			// Set Response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}
}

// Function to Update User Data By ID
func returnUserUpdate(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID Is Empty
	if len(params["id"]) == 0 {
		var response ResponseGlobal

		// Set Response Data
		response.Status = http.StatusBadRequest
		response.Message = "Error, ID Not Found"

		// Set Response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		// Get User ID As Integer Value
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var user Users
			var response ResponseGet

			// Decode JSON User Data
			// Use _ As Temporary Variable
			_ = json.NewDecoder(r.Body).Decode(&user)

			// Update User Data
			UsersData[userID-1].Name = user.Name
			UsersData[userID-1].Email = user.Email

			// Set Response Data
			response.Status = http.StatusCreated
			response.Message = "Success, User Updated"
			response.Data = UsersData[userID-1]

			// Set Response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			var response ResponseGlobal

			// Set Response Data
			response.Status = http.StatusInternalServerError
			response.Message = "Error, Coverting ID to Int"

			// Set Response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}
}

// Function to Update User Data By ID
func returnUserDelete(w http.ResponseWriter, r *http.Request) {
	// Get Parameters From URI
	params := mux.Vars(r)

	// Handle Error If Parameters ID Is Empty
	if len(params["id"]) == 0 {
		var response ResponseGlobal

		// Set Response Data
		response.Status = http.StatusBadRequest
		response.Message = "Error, ID Not Found"

		// Set Response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		// Get User ID As Integer Value
		userID, err := strconv.Atoi(params["id"])
		if err == nil {
			var response ResponseGlobal

			// Delete User Data
			UsersData = append(UsersData[userID-1], UsersData[userID]...)

			// Set Response Data
			response.Status = http.StatusOK
			response.Message = "Success, User Deleted"

			// Set Response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			var response ResponseGlobal

			// Set Response Data
			response.Status = http.StatusInternalServerError
			response.Message = "Error, Coverting ID to Int"

			// Set Response
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}
}
