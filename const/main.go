// -----------------------------
// This is a sample code in Go:
// - Constant Declaration
// - Print Constant Variable Value
// -----------------------------

package main

import (
	"fmt"
)

// Declare Constant Outside Function
const a string = "This is a constant variable"

func main() {
	// Print Constant Variable Outside Function
	fmt.Println("The a as constant variable has value:", a)

	// Declare Constant Variable Inside Function
	const b int = 10

	// Print Constant Variable Inside Function
	fmt.Println("The b as constant variable has value:", b)

	// Combine with Dynamic Variable with Aritmatic Operator
	z := 100 / b

	// Print Dynamic Variable
	fmt.Println("The z variable has value:", z)
}
