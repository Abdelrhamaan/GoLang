package main

import (
	"fmt"
	"sync"
)

// type Counter struct {
// 	mu    sync.Mutex
// 	count int
// }

// func (c *Counter) increment() {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	c.count++
// }

// func (c *Counter) getValue() int {
// 	c.mu.Lock()
// 	defer c.mu.Unlock()
// 	return c.count
// }

// func main() {
// 	// prevent multiple goroutines access shared resources in the same time
// 	// this avoid race conditions and data corruption
// 	// mutext will make just one goroutine access the value

// 	var wg sync.WaitGroup
// 	counter := &Counter{}
// 	numOfRoutines := 10

// 	for range numOfRoutines {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for range 1000 {
// 				counter.increment()
// 				// if we not using mutext sometimes value will change and will not result 1000
// 				// counter.count++
// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("Final value: ", counter.getValue())
// }

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int
	var numOfRoutines = 5

	wg.Add(5)
	increment := func() {
		defer wg.Done()
		for range 1000 {
			mu.Lock()
			// any var here will be locked
			counter++
			mu.Unlock()
		}
	}

	for range numOfRoutines {
		go increment()
	}

	wg.Wait()

	fmt.Println("Final value: ", counter)

}
