package main

import (
	"fmt"
	"time"
)

func main() {
	// this approach will case an error and we will see how to fix it
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// rec1 := <-ch1
	// fmt.Println("RECIEVED MSG1: ", rec1)
	// rec2 := <-ch2
	// fmt.Println("RECIEVED MSG2: ", rec2)

	//  we will fix it usig multiplexing(select - case - default)
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 1
	}()
	go func() {
		// time.Sleep(2 * time.Second)
		time.Sleep(1 * time.Second)
		ch2 <- 1
	}()

	// try to wait main thread to allow goroutines to finish
	time.Sleep(3 * time.Second)

	select {
	case rec := <-ch1:
		fmt.Println("RECIEVED from ch1: ", rec)
	case rec := <-ch2:
		fmt.Println("RECIEVED from ch2: ", rec)
	// if not default case an error may happens if no channel is ready
	// default is blockin
	// if we removed default and thers is no channels ready yet he will wait them
	// if no sender to channels and no defaults a deadlock error will happens
	default:
		fmt.Println("Now channel is ready")
	}

	// how to receive two messages from two channels

	go func() {
		ch1 <- 5
	}()
	go func() {
		ch2 <- 6
	}()

	time.Sleep(2 * time.Second)
	for range 2 {
		select {
		case msg := <-ch1:
			fmt.Println("REC FROM CH1", msg)
		case msg := <-ch2:
			fmt.Println("REC FROM CH2", msg)

		}
	}

	// handle time out
	ch3 := make(chan int32)

	go func() {
		time.Sleep(time.Second)
		ch3 <- 20
		close(ch3)
	}()

	select {
	case msg := <-ch3:
		fmt.Println("Rec from ch3", msg)
	// case <-time.After(3 * time.Second):
	case <-time.After(time.Second):
		fmt.Println("Time out")
	}

	// handle closing channels in while loop
	ch6 := make(chan int32)

	go func() {
		ch6 <- 100
		close(ch6)
	}()

	for {
		select {
		case msg, ok := <-ch6:
			fmt.Println("msg first is: ", msg)
			fmt.Println("ok first is: ", ok)
			if !ok {
				// go out of while loop
				fmt.Println("WHile closing channel")
				fmt.Println("msg is: ", msg)
				fmt.Println("ok is: ", ok)
				return
			}
			fmt.Println("We RECIEVD MSG AND CAHNNEL CLOSED", msg)
		}
	}
}
