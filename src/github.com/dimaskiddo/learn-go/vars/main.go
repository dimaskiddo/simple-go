// -----------------------------
// This is a sample code in Go:
// - Variable Declaration
// - Dynamic Variable Declaration
// - Set Variable Value
// - Print Variable Value
// -----------------------------

package main

import (
	"fmt"
)

func main() {
	// Declare Variables
	var x int
	x = 10

	var y float64
	y = 5.5

	// Print Variable Value
	fmt.Printf("The x variable has data type %T\n", x)
	fmt.Printf("The y variable has data type %T\n", y)

	// Print Varible as Standard
	fmt.Println("The x varible has value:", x)
	fmt.Println("The y varible has value:", y)

	// Line Break
	fmt.Println("--------------------------------------------")

	// Declare Dynamic Variable
	z := "This is string variable"
	w := 50

	// Print Dynamic Variable Value
	fmt.Printf("The z variable has data type %T\n", z)
	fmt.Printf("The w variable has data type %T\n", w)

	// Print Dynamic Varible as Standard
	fmt.Println("The z varible has value:", z)
	fmt.Println("The w varible has value:", w)
}
