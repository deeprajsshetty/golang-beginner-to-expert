// Prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*
Input:
go run main.go test1 test2 test3

Output:
test1 test2 test3
*/
