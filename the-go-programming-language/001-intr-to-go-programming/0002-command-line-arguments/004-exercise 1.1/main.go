//Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}

/*
Input:
go run main.go

Output:
/var/folders/bt/qmww6y011q3cftzhvs0hwql00000gn/T/go-build2381457517/b001/exe/main
*/
