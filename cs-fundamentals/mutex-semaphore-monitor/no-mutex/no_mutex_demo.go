package main

import (
	"fmt"
	"sync"
	"time"
)

// Shared Resource
var sharedCounter int = 0

// Function to simulate the task without Mutex
func incrementCounter(id int, wg *sync.WaitGroup) {
	// Decrease the WaitGroup counter when the function returns
	defer wg.Done()

	fmt.Printf("[Goroutine-%d] Starting task (No Lock)...\n", id)

	// --- CRITICAL SECTION (Unprotected) ---
	// Read the current value of the shared counter
	localCopy := sharedCounter

	// Simulate processing time (Context Switch usually happens here)
	// All goroutines might read the same initial value '0' before writing back
	time.Sleep(2 * time.Second)

	// Update the local copy and write back to the shared variable
	sharedCounter = localCopy + 1
	// --------------------------------------

	fmt.Printf("[Goroutine-%d] Finished. Local read: %d, Wrote: %d\n", id, localCopy, sharedCounter)
}

func main() {
	var wg sync.WaitGroup
	fmt.Println("--- Starting Simulation WITHOUT Mutex (Race Condition) ---")

	// Spawn 5 Goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go incrementCounter(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("------------------------------")
	fmt.Printf("Expected result: 5\n")
	fmt.Printf("Actual result:   %d\n", sharedCounter)
	fmt.Println("------------------------------")
}
