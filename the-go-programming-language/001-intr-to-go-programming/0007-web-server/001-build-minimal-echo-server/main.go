// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Register a handler function for the root ("/") path.
	http.HandleFunc("/", handler)

	// Start the HTTP server on port 8000 and handle requests.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler is a function that handles HTTP requests.
// It echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	// Write the URL.Path to the HTTP response writer.
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

/*
Input:
go build the-go-programming-language/001-intr-to-go-programming/0007-web-server/001-build-minimal-echo-server/main.go

main exe file will be generated in root folder. Move that particular exe to
./the-go-programming-language/001-intr-to-go-programming/0007-web-server/001-build-minimal-echo-server/ path.

Run Server
./the-go-programming-language/001-intr-to-go-programming/0007-web-server/001-build-minimal-echo-server/main http://localhost:8000

Curl Command - curl http://localhost:8000/help

Output:
URL.Path = "/help"

*/
