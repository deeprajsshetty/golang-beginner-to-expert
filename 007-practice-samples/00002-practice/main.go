package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main()  {
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, dir := range pathSplit {
		fmt.Println(dir)
	}
}