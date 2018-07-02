package controllers

import (
	"net/http"

	"github.com/dimaskiddo/simple-go/api-mongo/routers"
)

// Function to Show Information at Index
func GetIndex(w http.ResponseWriter, r *http.Request) {
	var response routers.Response

	response.Status = true
	response.Code = http.StatusOK
	response.Message = "Simple Go Programming Example (API MongoDB)"

	routers.ResponseWrite(w, response.Code, response)
}
