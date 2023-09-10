// Server is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex // Mutex for protecting the 'count' variable
var count int     // Counter to keep track of requests

func main() {
	// Register handlers for the root ("/") and "/count" paths
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)

	// Start the HTTP server on port 8000 and handle requests.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler is a function that handles HTTP requests for the root ("/") path.
// It increments the request count and echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()   // Lock the mutex to protect the 'count' variable
	count++     // Increment the request count
	mu.Unlock() // Unlock the mutex

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter is a function that handles HTTP requests for the "/count" path.
// It displays the current request count.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock() // Lock the mutex to access the 'count' variable safely
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock() // Unlock the mutex
}

/*
Input:
go build the-go-programming-language/001-intr-to-go-programming/0007-web-server/002-build-minimal-echo-and-counter-server/main.go

main exe file will be generated in root folder. Move that particular exe to
./the-go-programming-language/001-intr-to-go-programming/0007-web-server/002-build-minimal-echo-and-counter-server/ path.

Run Server
./the-go-programming-language/001-intr-to-go-programming/0007-web-server/002-build-minimal-echo-and-counter-server/main http://localhost:8000

curl http://localhost:8000/counter1
curl http://localhost:8000/counter2
curl http://localhost:8000/counter3
curl http://localhost:8000/count

Output:
URL.Path = "/counter1"
URL.Path = "/counter2"
URL.Path = "/counter3"

Count 3

*/
