// Prints its command-line arguments using strings join

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], "---"))
}

/*
Input:
go run main.go test1 test2 test3

Output:
test1---test2---test3
*/
