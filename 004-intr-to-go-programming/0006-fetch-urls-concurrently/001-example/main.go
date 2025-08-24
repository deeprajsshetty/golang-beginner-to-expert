// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	// Launch a goroutine for each URL
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	// Receive and print results from the goroutines
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	// Calculate and print the total elapsed time
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch fetches a URL and reports the elapsed time and the size of the response.
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Send an error message to the channel
		return
	}
	defer resp.Body.Close() // Ensure the response body is closed when done

	// Discard the response body and count the bytes
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	// Calculate the elapsed time and send the result to the channel
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

/*
Input:
go run the-go-programming-language/002-intr-to-go-programming/0006-fetch-urls-concurrently/001-example/main.go https://
www.rust-lang.org/ https://go.dev/

Output:
0.65s    61870  https://go.dev/
2.19s    19737  https://www.rust-lang.org/
2.19s elapsed
*/
