package main

type ResGetAllUser struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
}

type ResGetOneUser struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User
}
