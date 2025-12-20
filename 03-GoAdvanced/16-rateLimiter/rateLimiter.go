package main

import (
	"fmt"
	"time"
)

type TokenBucket struct {
	Tokens     chan struct{}
	RefillTime time.Duration
}

func (tb *TokenBucket) StartRefill() {
	ticker := time.NewTicker(tb.RefillTime)
	defer ticker.Stop()

	for range ticker.C {
		select {
		// If channel has space → add token
		case tb.Tokens <- struct{}{}: // this line is sending empty struct (token) to channel
		// If channel is full → silently drop token
		default:
		}
	}
}

func (tb *TokenBucket) AllowRequests() bool {
	select {
	case <-tb.Tokens:
		return true
	default:
		return false
	}
}

func NewRateLimiter(refillTime time.Duration, tokenCapacity int) *TokenBucket {
	newTokenBucket := &TokenBucket{
		Tokens:     make(chan struct{}, tokenCapacity),
		RefillTime: refillTime,
	}

	for range tokenCapacity {
		newTokenBucket.Tokens <- struct{}{}
	}
	go newTokenBucket.StartRefill()
	return newTokenBucket

}

func main() {

	tokenBucket := NewRateLimiter(time.Second, 3)
	testToken(tokenBucket)
	time.Sleep(time.Second)
	testToken(tokenBucket)

}

func testToken(tb *TokenBucket) {
	for range 10 {
		if tb.AllowRequests() {
			fmt.Println("req allowed")
		} else {
			fmt.Println("req denied")
		}
	}
}
