package main

import (
	"fmt"
	"sync"
	"time"
)

// Shared Resource
var sharedCounter int = 0

// Initialize a Mutex to protect the critical section
var mu sync.Mutex

func incrementCounter(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("[Goroutine-%d] Waiting for Mutex...\n", id)

	// --- CRITICAL SECTION (Protected) ---
	// Lock the mutex before accessing the shared resource
	// Other goroutines must wait if the mutex is already locked
	mu.Lock()
	
	fmt.Printf("[Goroutine-%d] Acquired Mutex. Processing...\n", id)

	// Safe to access sharedCounter now
	localCopy := sharedCounter
	time.Sleep(2 * time.Second) 
	sharedCounter = localCopy + 1

	fmt.Printf("[Goroutine-%d] Updated counter to %d\n", id, sharedCounter)
	
	// Unlock the mutex to allow other waiting goroutines to enter
	mu.Unlock()
	// ------------------------------------
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("--- Starting Simulation WITH Mutex ---")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go incrementCounter(i, &wg)
	}

	wg.Wait()

	fmt.Printf("Final result: %d\n", sharedCounter)
}