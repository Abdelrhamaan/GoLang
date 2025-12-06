package main

import (
	"fmt"
	"time"
)

// func main() {
// 	// created unbeffered channel which needs immediate reciever
// 	// you cannot save data in channel
// 	// make(chan Type)

// 	// creating buffered channels which not need immediate reciever
// 	// can save data in channel until i use it
// 	// make(chan Type, capacity)

// 	// blocked ubuffered channel will get dead lock error if not using goroutine and make immediate receiver
// 	// ch := make(chan int)

// 	// non blocked buffered channel will block in two cases:
// 	//  1- when channel is full
// 	//  2- when channed becomes empty
// 	ch := make(chan int, 3)

// 	ch <- 1
// 	ch <- 2
// 	ch <- 10
// 	// this will cause deadlock error
// 	// if we consumed from it before adding this part no error will happens
// 	// ch <- 10

// 	for range len(ch) {
// 		fmt.Println("Value is: ", <-ch)
// 	}

// 	fmt.Println("Buffered msg is in channel now")
// }

func main() {

	// in buffered channels two errors can occur:
	//  1- first one if you tried to add to full channel
	//  2- second one if you tried to consume from empty channel
	// ch := make(chan int, 3)
	// ch <- 1
	// ch <- 2
	// ch <- 3
	// // this wil cause an error you should consume first then refill
	// // ch <- 3

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("first rec value: ", <-ch)
	// }()

	// // add new value in chan
	// ch <- 4
	// fmt.Println("Rec value is: ", <-ch)
	// fmt.Println("Rec value is: ", <-ch)
	// fmt.Println("Rec value is: ", <-ch)

	// // consume from empty channel
	// ch2 := make(chan int32, 3)

	// // this will not cause error
	// // go func() {
	// // 	fmt.Println("Consuming from empty chan: ", <-ch2)
	// // }()
	// // this will cause error
	// fmt.Println("Consuming from empty chan: ", <-ch2)

	// fmt.Println("finished")
	// =====================================================

	// important weired observation
	ch3 := make(chan int, 2)

	ch3 <- 1
	ch3 <- 2

	fmt.Println("Starting weiring recieved")
	go func() {
		// when you adding print statement here second print will happens
		// but when you remove it second print sometimes happens sometimes no
		// so when <-ch3 happens it remove from channel and returns fast to main thread
		fmt.Println("First Print ")
		time.Sleep(2 * time.Second)
		fmt.Println("Second Print: ", <-ch3)
	}()

	fmt.Println("Adding new values in the channel: ")
	ch3 <- 3
	fmt.Println("Finishing")
}

// WHY does adding fmt.Println("First Print") change behavior?

// Because printing to stdout is slow (syscall), and this gives the scheduler time to start the goroutine before main blocks.

// So when you add:
// fmt.Println("First Print")

// What happens?

// goroutine prints immediately

// scheduler gives it CPU time

// goroutine sleeps

// main is still running

// eventually goroutine wakes up → reads from ch3 → unblocks main

// main continues

// you see “Second Print”

// ❗ When you remove that print:

// The goroutine may start late.

// Timeline:

// main fills channel

// main launches goroutine

// main IMMEDIATELY sends 3 → but the goroutine has NOT run yet

// main blocks

// program prints nothing else for 2 seconds

// sometimes scheduler kills main before other goroutine runs → so no output

// sometimes goroutine wakes → prints output

// This creates the “sometimes works, sometimes doesn’t” effect.

// 🧠 KEY REASON:
// goroutines DO NOT start immediately after go func()

// They start when the scheduler decides, which varies run-to-run.

// Even tiny things (like print statements, logging, GC, CPU load) change the scheduling order.
