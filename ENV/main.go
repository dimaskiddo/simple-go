// -----------------------------
// This is a sample code in Go:
// - Read Enviroment Variable
// - Set Environment Variable
// -----------------------------

package main

import (
	"fmt"
	"os"
)

func main() {
	// Read Environment Variable
	config := os.Getenv("CONFIG")

	// If Environment Variable is Empty Then
	//    Set The Environment Variable with Your Default Value and
	//    Re-Read The Environment Variable
	if len(config) == 0 {
		// Set Environment Variable
		os.Setenv("CONFIG", "This is my configuration")

		// Re-Read Environment Variable
		config = os.Getenv("CONFIG")
	}

	// Print Environment Variable
	fmt.Println(config)
}
