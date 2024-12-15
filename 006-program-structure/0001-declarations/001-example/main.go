// This Go program calculates and prints the boiling point of water
// in both Fahrenheit and Celsius.
package main

import "fmt"

// The constant boilingF stores the boiling point of water in Fahrenheit.
const boilingF = 212.0

func main() {
	// Declare and initialize the variable 'f' with the value of boilingF.
	var f = boilingF

	// Calculate the equivalent temperature in Celsius and store it in 'c'.
	var c = (f - 32) * 5 / 9

	// Print the boiling point in both Fahrenheit and Celsius.
	fmt.Printf("Boiling point = %g°F or %g°C\n", f, c)
}

/*
Input:
go run the-go-programming-language/003-program-structure/0001-declarations/001-example/main.go

Output:
Boiling point = 212°F or 100°C
*/
