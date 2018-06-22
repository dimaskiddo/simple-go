package main

import (
	"fmt"
)

func callByValue(param string) string {
	return param + " By Value"
}

func callByRefTrue(param *string) {
	*param += " By Ref"
}

func callByRefFalse(param string) {
	param += " By Ref"

	fmt.Println("Nilai byRefFalse Dari Dalam Fungsi callByRefFalse: ")
	fmt.Println(param)

	fmt.Println("")
}

func main() {
	var byValue, byRefTrue, byRefFalse string = "Hello", "Hello", "Hello"

	// ---------------------------------------------------------------
	fmt.Println("Nilai byValue Sebelum Memasuki Fungsi callByValue: ")
	fmt.Println(byValue)

	fmt.Println("")

	byValue = callByValue(byValue)

	fmt.Println("Nilai byValue Sesudah Memasuki Fungsi callByValue: ")
	fmt.Println(byValue)

	fmt.Println("")
	// ---------------------------------------------------------------

	// ---------------------------------------------------------------
	fmt.Println("Nilai byRefTrue Sebelum Memasuki Fungsi callByRefTrue: ")
	fmt.Println(byRefTrue)

	fmt.Println("")

	callByRefTrue(&byRefTrue)

	fmt.Println("Nilai byRefTrue Sesudah Memasuki Fungsi callByRefTrue: ")
	fmt.Println(byRefTrue)

	fmt.Println("")
	// ---------------------------------------------------------------

	// ---------------------------------------------------------------
	fmt.Println("Nilai byRefFalse Sebelum Memasuki Fungsi callByRefFalse: ")
	fmt.Println(byRefFalse)

	fmt.Println("")

	callByRefFalse(byRefFalse)

	fmt.Println("Nilai byRefFalse Sesudah Memasuki Fungsi callByRefFalse: ")
	fmt.Println(byRefFalse)

	fmt.Println("")
	// ---------------------------------------------------------------
}
