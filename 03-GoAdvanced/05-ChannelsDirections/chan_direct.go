package main

import "fmt"

// ------------------------------------------------------------
// 1. Bidirectional channel (default chan T)
// ------------------------------------------------------------
func bidirectionalDemo() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i // send
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("bidirectional received:", v)
	}
}

// ------------------------------------------------------------
// 2. Send‑only channel (chan<- T)
// ------------------------------------------------------------
func producer(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i // allowed: send only
	}
	close(out)
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("send‑only received:", v)
	}
}

func sendOnlyDemo() {
	ch := make(chan int) // underlying channel is bidirectional
	go producer(ch)      // passed as chan<- int
	consumer(ch)         // passed as <-chan int
}

// ------------------------------------------------------------
// 3. Receive‑only channel (<-chan T)
// ------------------------------------------------------------
func generator() <-chan string {
	ch := make(chan string)
	go func() {
		for _, s := range []string{"alpha", "beta", "gamma"} {
			ch <- s
		}
		close(ch)
	}()
	return ch // caller sees only receive‑only channel
}

func receiveOnlyDemo() {
	msgs := generator()
	for m := range msgs {
		fmt.Println("receive‑only got:", m)
	}
}

// ------------------------------------------------------------
// 4. Pipe: receive‑only -> send‑only
// ------------------------------------------------------------
func pipe(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * 2 // transform while forwarding
	}
	close(out)
}

func pipeDemo() {
	src := make(chan int)
	dst := make(chan int)

	go func() {
		for i := 1; i <= 4; i++ {
			src <- i
		}
		close(src)
	}()

	go pipe(src, dst)

	for v := range dst {
		fmt.Println("pipe output:", v)
	}
}

// ------------------------------------------------------------
// main – runs all demos sequentially
// ------------------------------------------------------------
func main() {
	fmt.Println("--- Bidirectional Demo ---")
	bidirectionalDemo()

	fmt.Println("--- Send‑Only Demo ---")
	sendOnlyDemo()

	fmt.Println("--- Receive‑Only Demo ---")
	receiveOnlyDemo()

	fmt.Println("--- Pipe Demo (receive‑only → send‑only) ---")
	pipeDemo()
}
