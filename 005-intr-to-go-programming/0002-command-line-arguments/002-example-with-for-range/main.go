// Prints its command-line arguments using for range

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "/"
	}
	fmt.Println(s)
}

/*
Input:
go run main.go test1 test2 test3

Output:
test1/test2/test3
*/
