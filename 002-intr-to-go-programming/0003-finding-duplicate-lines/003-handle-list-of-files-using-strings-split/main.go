// Prints the count and text of lines that appear more than once in the input.
// First,it only reads named files, not the standard input, since ReadFile
// requires a file name argument. Second, we moved the counting of the lines
// back into main, since it is now needed in only one place.

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	// Iterate through command-line arguments (file names)
	for _, filename := range os.Args[1:] {
		// Read the content of the file
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Error reading file %s: %v\n", filename, err)
			continue
		}

		// Split the file content into lines and count them
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if line != "" { // Skip empty lines
				counts[line]++
			}
		}
	}

	// Print lines that appear more than once
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

/*
Input:
go run the-go-programming-language/002-intr-to-go-programming/0003-finding-duplicate-lines/003-handle-list-of-files-usi
ng-strings-split/main.go ./the-go-programming-language/002-intr-to-go-programming/0003-finding-duplicate-lines/003-handle-list-of-files-using-strings-split/file1.txt ./the-go-progra
mming-language/002-intr-to-go-programming/0003-finding-duplicate-lines/003-handle-list-of-files-using-strings-split/file2.txt

Output:
4       one
2       two
2       three
4       ten
2       nine
*/
