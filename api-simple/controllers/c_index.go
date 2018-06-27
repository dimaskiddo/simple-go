package controllers

import (
	"net/http"

	"github.com/dimaskiddo/simple-go/api-simple/helpers"
	"github.com/dimaskiddo/simple-go/api-simple/routers"
)

// Function to Show API Information
func GetIndex(w http.ResponseWriter, r *http.Request) {
	var response routers.Response

	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Simple Go Programming Example (API Simple)"

	routers.ResponseWrite(w, response.Code, response)
}

// Function to Get JWT Authentication
func GetAuthenticationJWT(w http.ResponseWriter, r *http.Request) {
	// Get Username And Password From HTTP Request
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Make Sure Username and Passwor is Not Empty
	if len(username) == 0 || len(password) == 0 {
		routers.ResponseBadRequest(w)
		return
	}

	// Some Business Logic Here to Match The Username and Password
	if username == "user" && password == "password" {
		// Get JWT Token From Pre-Defined Function
		token, err := helpers.GetJWTToken(username)
		if err != nil {
			routers.ResponseInternalError(w)
		} else {
			var response helpers.ResponseJWT

			response.Status = true
			response.Code = http.StatusOK
			response.Token = "Bearer " + token

			routers.ResponseWrite(w, response.Code, response)
		}
	} else {
		routers.ResponseUnauthorized(w)
	}
}
