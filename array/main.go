// -----------------------------
// This is a sample code in Go:
// - Type Struct Declaration
// - Array of Struct Declaration
// - Inserting Struct to Array of Struct
// - Removing Struct from Array of Struct
// -----------------------------

package main

import (
	"fmt"
)

// Declare Struct of User
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Declare Array of Struct User
var users []User

func main() {
	// Declare Variable from Struct User
	var user User

	user.ID = 1
	user.Name = "Dimas"

	users = append(users, user)

	user.ID = 2
	user.Name = "Aas"

	users = append(users, user)

	fmt.Println(users)

	user.ID = 3
	user.Name = "Krisna"

	users = append(users, user)

	fmt.Println(users)

	user.ID = 4
	user.Name = "Donni"

	users = append(users, user)

	fmt.Println(users)

	users = append(users[:3], users[4:]...)

	fmt.Println(users)

	users = append(users[:1], users[2:]...)

	fmt.Println(users)
}
