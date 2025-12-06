package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch := make(chan struct{})

// 	go func() {
// 		fmt.Println("Starting working in channels")
// 		time.Sleep(2 * time.Second)
// 		ch <- struct{}{}
// 	}()

// 	<-ch
// 	fmt.Println("Consuming unbeffered channels")

// }

// func main() {
// 	goRoutines := 3
// 	ch := make(chan int, 3)
// 	for i := range goRoutines {
// 		time.Sleep(time.Second)
// 		go func(id int) {
// 			fmt.Printf("Working goroutine is %d\n", i)
// 			ch <- i // sending signal of completion
// 		}(i)
// 	}

// 	// if we have many goroutines we should made chan. synch.
// 	// by iterating over capacity of channel
// 	for range goRoutines {
// 		<-ch // wait for all goroutines are complete
// 	}
// 	// main rule if goroutine executin finished main thread will wait for it if not finished main thread will not wait for it
// 	// this will not recieve all values from channel
// 	// no receiver for third goroutine so out main progoram will not wait it
// 	// and memory leak will happens should we should handle it
// 	// for range 2 {
// 	// 	<-ch
// 	// }

// 	fmt.Println("Goroutines finished")
// }

// synchronization data excahnging
func main() {
	dataCh := make(chan string)

	go func() {
		fmt.Println("Adding data to channel")
		for i := range 5 {
			dataCh <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		// before closing channel for new reciever
		// if we commit it an error will happen
		// the error will happens because second for loop will continue consume from chan
		//  while chan not containing new valus
		// close(dataCh)
	}()
	// close(dataCh) // this will prevent receiving any thing because channel is closed and have no values

	fmt.Println("len of data channel is: ", len(dataCh))

	for value := range dataCh {
		fmt.Println("Recieved value from channel is: ", value, time.Now())
	}
}
