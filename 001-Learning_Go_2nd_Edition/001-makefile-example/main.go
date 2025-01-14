package main

import "fmt"

func main() {
	fmt.Println("Hello, Makefile example")
}

/*
Input:
make

Output: [makefile_example exe should be available in the directory]
go fmt ./...
go vet ./...
go build
*/
