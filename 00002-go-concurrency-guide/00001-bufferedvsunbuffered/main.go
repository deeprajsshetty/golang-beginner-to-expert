package main

import (
	"fmt"
	"sync"
	"time"
)

// Entry point of the program
func main() {
	now := time.Now() // Record current time to measure execution duration

	var wg sync.WaitGroup // WaitGroup is used to wait for multiple goroutines to finish
	userID := 10 // Simulating operations for a user with ID 10

	// === Using Buffered Channel ===
	respch := make(chan string, 3) // Buffered channel with size 3 allows 3 goroutines to send without blocking

	// Add 3 to the WaitGroup counter, one for each goroutine
	wg.Add(1)
	go fetchUserDataUsingBufferedChannel(userID, respch, &wg) // Start goroutine for user data
	wg.Add(1)
	go fetchUserRecommendationsUsingBufferedChannel(userID, respch, &wg) // Start goroutine for recommendations
	wg.Add(1)
	go fetchUserLikesUsingBufferedChannel(userID, respch, &wg) // Start goroutine for likes

	// Wait for all three goroutines to finish
	wg.Wait()

	// Close the channel since no more data will be sent.
	// Required to end the range loop below.
	close(respch)

	// Read all values from the buffered channel
	for resp := range respch {
		fmt.Println(resp)
	}

	// === Using Unbuffered Channels ===
	// Each goroutine gets its own dedicated unbuffered channel
	userDataCh := make(chan string)
	userRecmCh := make(chan string)
	userLikesCh := make(chan string)

	// Launch goroutines for unbuffered channels
	go fetchUserDataUsingUnbufferedChannel(userID, userDataCh)
	go fetchUserRecommendationsUsingUnbufferedChannel(userID, userRecmCh)
	go fetchUserLikesUsingUnbufferedChannel(userID, userLikesCh)

	// Read from each unbuffered channel.
	// These reads will block until the corresponding goroutine sends data.
	fmt.Println(<-userDataCh)
	fmt.Println(<-userRecmCh)
	fmt.Println(<-userLikesCh)

	// Print total execution time
	fmt.Println(time.Since(now))
}

// Buffered Channel Goroutines
func fetchUserDataUsingBufferedChannel(userId int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	time.Sleep(80 * time.Millisecond) // Simulate latency
	respch <- fmt.Sprintf("Buffered: user data of %d", userId) // Send response to buffered channel
	// wg.Done() // Signal that this goroutine is done
}

func fetchUserRecommendationsUsingBufferedChannel(userId int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(120 * time.Millisecond)
	respch <- fmt.Sprintf("Buffered: user recommendations for %d", userId)
	// wg.Done()
}

func fetchUserLikesUsingBufferedChannel(userId int, respch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(50 * time.Millisecond)
	respch <- fmt.Sprintf("Buffered: user likes for %d", userId)
	// wg.Done()
}

// Unbuffered Channel Goroutines
func fetchUserDataUsingUnbufferedChannel(userId int, userDataCh chan string) {
	time.Sleep(80 * time.Millisecond)
	userDataCh <- fmt.Sprintf("Unbuffered: user data of %d", userId) // Blocks until main goroutine receives
}

func fetchUserRecommendationsUsingUnbufferedChannel(userId int, userRecmCh chan string) {
	time.Sleep(120 * time.Millisecond)
	userRecmCh <- fmt.Sprintf("Unbuffered: user recommendations for %d", userId)
}

func fetchUserLikesUsingUnbufferedChannel(userId int, userLikesCh chan string) {
	time.Sleep(50 * time.Millisecond)
	userLikesCh <- fmt.Sprintf("Unbuffered: user likes for %d", userId)
}