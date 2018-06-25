package main

import (
	"encoding/json"
	"net/http"
)

type ResGlobal struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func resWrite(w http.ResponseWriter, status int, res interface{}) {
	// Write Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Write JSON to Response
	json.NewEncoder(w).Encode(res)
}

func resBadRequest(w http.ResponseWriter) {
	var response ResGlobal

	// Set Response Data
	response.Status = http.StatusBadRequest
	response.Message = "Bad Request"

	// Set Response Data to HTTP
	resWrite(w, response.Status, response)
}

func resInternalError(w http.ResponseWriter) {
	var response ResGlobal

	// Set Response Data
	response.Status = http.StatusInternalServerError
	response.Message = "Internal Server Error"

	// Set Response Data to HTTP
	resWrite(w, response.Status, response)
}
