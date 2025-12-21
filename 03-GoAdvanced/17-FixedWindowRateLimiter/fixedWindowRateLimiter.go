package main

import (
	"fmt"
	"sync"
	"time"
)

type rateLimiter struct {
	mu          sync.Mutex
	limit       int
	count       int
	window      time.Duration
	resetWindow time.Time
}

func newLimiter(limit int, window time.Duration) *rateLimiter {
	return &rateLimiter{
		limit:  limit,
		window: window,
	}
}

func (rl *rateLimiter) isAllowed() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()

	if now.After(rl.resetWindow) {
		rl.resetWindow = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false

}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	for i := range 2 {
		go func(id int) {
			defer wg.Done()
			rateLimiter := newLimiter(3, time.Second)
			for range 10 {
				if rateLimiter.isAllowed() {
					fmt.Printf("Request allowed %d\n", id)
				} else {
					fmt.Printf("Request Denied %d\n", id)
				}
			}
		}(i)
	}
	wg.Wait()

}
