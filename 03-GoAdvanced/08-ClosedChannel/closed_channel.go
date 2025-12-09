package main

import (
	"fmt"
	"sync"
)

// Demonstrates correct channel closing patterns:
// 1. Only the sender closes the channel.
// 2. The channel is closed exactly once.
// 3. Closing is performed by a single goroutine (the sender).
func main() {
	// Example 1: Simple producer‑consumer where the producer closes.
	fmt.Println("--- Example 1: Sender closes channel ---")
	data := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Producer sends values then closes the channel exactly once.
		for i := 0; i < 5; i++ {
			data <- i
		}
		// Only the sender is allowed to close – this is safe.
		close(data)
	}()

	// Consumer reads until the channel is closed.
	for v := range data {
		fmt.Printf("received %d\n", v)
	}
	wg.Wait()

	// Example 2: Prevent double close using sync.Once.
	fmt.Println("--- Example 2: Prevent double close ---")
	ch := make(chan struct{})
	var once sync.Once
	closeOnce := func() {
		once.Do(func() { close(ch) })
	}
	var wg2 sync.WaitGroup
	wg2.Add(2)
	go func() { defer wg2.Done(); closeOnce() }()
	go func() { defer wg2.Done(); closeOnce() }()
	wg2.Wait()
	fmt.Println("channel closed exactly once without panic")

	// Example 3: Multiple producers, single closer.
	fmt.Println("--- Example 3: Multiple producers, single closer ---")
	multi := make(chan int)
	var wg3 sync.WaitGroup
	producer := func(id int) {
		defer wg3.Done()
		for i := 0; i < 3; i++ {
			multi <- id*10 + i
		}
	}
	wg3.Add(3)
	go producer(1)
	go producer(2)
	go producer(3)
	// Close the channel after all producers are done.
	go func() {
		wg3.Wait()
		close(multi)
	}()
	// Consumer reads until closed.
	for v := range multi {
		fmt.Printf("multi received %d\n", v)
	}
}
