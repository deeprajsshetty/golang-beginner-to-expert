// Prints the text of each line that appears more than once in the
// standard input, preceded by its count

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Initialize a map to store line counts
	counts := make(map[string]int)

	// Create a Scanner to read input from stdin
	input := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter multiple lines (or type 'exit' to quit):\n")

	// Use the Scanner to read lines of text until "exit" is entered
	for input.Scan() {
		text := input.Text()

		// Check if the user entered "exit" to quit the loop
		if text == "exit" {
			break
		}

		// Increment the count for the current line
		counts[text]++
	}

	// Print lines that appear more than once
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}

/*
Input:
go run main.go
Enter multiple lines (or press Enter to exit):
one
one
Two
Three
Two
exit

Output:
2       one
2       Two
*/
