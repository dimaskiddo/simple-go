// -----------------------------
// This is a sample code in Go:
// - Declaration Type Struct
// - Declaration Array of Struct
// - Print Array of Struct
// - Insert Struct to Array of Struct
// - Remove Struct from Array of Struct
// - Update Struct to Array of Struct
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

	// Assign Value to Struct User
	user.ID = 1
	user.Name = "Dimas"

	// Insert Struct User to Array of Struct User
	users = append(users, user)

	// Assign Value to Struct User
	user.ID = 2
	user.Name = "Aas"

	// Insert More Struct User to Array of Struct User
	users = append(users, user)

	// Print Current Array of Struct User
	fmt.Println(users)

	// Assign Value to Struct User
	user.ID = 3
	user.Name = "Krisna"

	// Insert More Struct User to Array of Struct User
	users = append(users, user)

	// Print Current Array of Struct User
	fmt.Println(users)

	// Assign Value to Struct User
	user.ID = 4
	user.Name = "Donni"

	// Insert More Struct User to Array of Struct User
	users = append(users, user)

	// Print Current Array of Struct User
	fmt.Println(users)

	// Remove 4th Struct from Array of Struct
	users = append(users[:3], users[4:]...)

	// Print Current Array of Struct User
	fmt.Println(users)

	// Remove 2nd Struct from Array of Struct
	users = append(users[:1], users[2:]...)

	// Print Current Array of Struct User
	fmt.Println(users)

	// Update 2nd Struct to Array of Struct By Directly
	users[1].ID = 3
	users[1].Name = "Jajang"

	// Print Current Array of Struct User
	fmt.Println(users)

	// Assign Value to Struct User
	user.ID = 3
	user.Name = "Arif"

	// Update 2nd Struct to Array of Struct By Append
	users = append(users[:1], user)

	// Print Current Array of Struct User
	fmt.Println(users)
}
