// The following Go program, Ftoc, performs Fahrenheit-to-Celsius conversions
// and prints the results.
package main

import "fmt"

// Constants for freezing and boiling points in Fahrenheit.
const freezingF, boilingF = 32.0, 212.0

// fToC converts a temperature in Fahrenheit to Celsius.
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	// Print Fahrenheit-to-Celsius conversions for freezing and boiling points.
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

/*
Input:
go run the-go-programming-language/003-program-structure/0001-declarations/002-temp-conversion/main.go

Output:
32°F = 0°C
212°F = 100°C
*/
