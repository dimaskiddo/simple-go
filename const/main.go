// -----------------------------
// This is a sample code in Go:
// - Constant Declaration
// - Print Constant Variable Value
// -----------------------------

package main

import (
	"fmt"
)

// Declare Global Constant Variable
const a string = "This is a constant variable"

func main() {
	// Print Global Constant Variable
	fmt.Println("The a as constant variable has value:", a)

	// Declare Local Constant Variable
	const b int = 10

	// Print Local Constant Variable
	fmt.Println("The b as constant variable has value:", b)

	// Combine with Dynamic Variable with Aritmatic Operator
	z := 100 / b

	// Print Dynamic Variable
	fmt.Println("The z variable has value:", z)
}
