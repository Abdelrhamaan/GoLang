package main

import (
	"fmt"
	"time"
)

func main() {
	// ------------------None blocking receiving channels
	// ch := make(chan int)

	// select {
	// case msg := <-ch:
	// 	fmt.Println("Recived message is ", msg)
	// default:
	// 	fmt.Println("No messages to recieve")
	// }

	// ------------------Non blocking sender msgs
	// select {
	// case ch <- 30:
	// 	fmt.Println("Channel is sending msgs")
	// default:
	// 	fmt.Println("No reviever founds to recieve sended msgs")
	// }

	// ------------------Non blocking real time

	data := make(chan int)
	isExit := make(chan bool)

	go func() {
		for {
			select {
			case comingData := <-data:
				fmt.Println("Coming Data: ", comingData)
			case <-isExit:
				fmt.Println("Stop Recving data")
				return
			default:
				fmt.Println("Waiting to receive data")
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}
	isExit <- true

}
