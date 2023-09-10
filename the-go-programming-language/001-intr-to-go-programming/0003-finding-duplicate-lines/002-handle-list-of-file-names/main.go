// Prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Initialize a map to store line counts
	counts := make(map[string]int)

	// Get command line arguments (file names)
	files := os.Args[1:]

	// If no file names provided as arguments, read from stdin
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		// Loop through each file provided as an argument
		for _, arg := range files {
			// Open the file
			f, err := os.Open(arg)
			if err != nil {
				// Print an error message if the file couldn't be opened
				fmt.Fprintf(os.Stdout, "Error opening file: %v\n", err)
				continue // Continue to the next file
			}
			countLines(f, counts)
			f.Close()
		}
	}

	// Print lines that appear more than once
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// countLines reads from a file and counts the occurrences of each line.
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// Increment the count for the current line
		counts[input.Text()]++
	}

	// NOTE: ignoring potential errors from input.Err() for simplicity
}

/*
Input:
go run the-go-programming-language/001-intr-to-go-programming/0003-finding-duplicate-lines/002-handle-list-of-file-name
s/main.go ./the-go-programming-language/001-intr-to-go-programming/0003-finding-duplicate-lines/002-handle-list-of-file-names/file1.txt ./the-go-programming-language/001-intr-to-go-programming/0003-finding-duplicate-lines/002-handle-list-of-file-names/file2.txt

Output:
4       one
2       two
2       three
4       ten
2       nine
*/
