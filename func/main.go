// -----------------------------
// This is a sample code in Go:
// - Declare Procedure
// - Declare Function
// - Function Call By Value
// - Function Call By Reference
// -----------------------------

package main

import (
	"fmt"
)

// Declare Function With Call By Value
// So It's Need To Return Some Data Type
func funcByValue(param string) string {
	return param + " This Value Has Been Changed By The funcByValue Function"
}

// Declare Function With Call By Reference
// So It's Need To Get Some Variable Memory Address
// And Return It To The Memory
func funcByReference(param *string) {
	*param = *param + " This Value Has Been Changed By The funcByReference Function"
}

// Declare Procedure
func procHello(param string) {
	param = param + " This Value Has Been Changed By The procHello Procedure"

	fmt.Println("")
	fmt.Println("Value of Variable Printed From procHello Procedure: ")
	fmt.Println(param)
}

func main() {
	var a, b, c string = "Hello", "Hello", "Hello"

	// ---------------------------------------------------------------
	fmt.Println("Value of Variable 'a' Before Go To The funcByValue Function: ")
	fmt.Println(a)

	// By Value Function Need a Variable To Return a Value
	// So We Need a Variable To Get It
	a = funcByValue(a)
	fmt.Println("")

	fmt.Println("Value of Variable 'a' After Go To The funcByValue Function: ")
	fmt.Println(a)
	// ---------------------------------------------------------------
	fmt.Println("---------------------------------------------------")
	// ---------------------------------------------------------------
	fmt.Println("Value of Variable 'b' Before Go To The funcByReference Function: ")
	fmt.Println(b)

	// By Reference Function Need To Get a Variable Memory to Return a Value
	// So We Need a Variable Memory To Get It
	funcByReference(&b)
	fmt.Println("")

	fmt.Println("Value of Variable 'b' After Go To The funcByReference Function: ")
	fmt.Println(b)
	// ---------------------------------------------------------------
	fmt.Println("---------------------------------------------------")
	// ---------------------------------------------------------------
	fmt.Println("Value of Variable 'c' Before Go To The procHello Procedure: ")
	fmt.Println(c)

	// Call Procedure
	procHello(c)
	fmt.Println("")

	fmt.Println("Value of Variable 'c' After Go To The procHello Procedure: ")
	fmt.Println(c)
	// ---------------------------------------------------------------
}
