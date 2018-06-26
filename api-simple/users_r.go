package main

type ResGetUser struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}
