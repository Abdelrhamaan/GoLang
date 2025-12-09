package main

import (
	"fmt"
	"time"
)

// ============ basic timer ==================
// func main() {
// 	fmt.Println("Starting New time")
// 	// time.sleep(time.second) is blocking by its nature
// 	timer := time.NewTimer(2 * time.Second)
// 	fmt.Println("Working and timer not blocking executing...")

// 	// lets force stopping timer
// 	isStopped := timer.Stop()
// 	if isStopped {
// 		fmt.Println("Time is stopped")
// 	} else {
// 		// after reciving in chan will exceute
// 		<-timer.C
// 	}
// 	// how to reset stopped or current timer
// 	timer.Reset(3 * time.Second)
// 	fmt.Println("Reseting time...")
// 	// after reciving in chan will exceute
// 	<-timer.C
// 	fmt.Println("Finished...")
// }

// ============== timing out =====================
func longRunningOperations() {
	for i := range 2 {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	timeOut := time.After(4 * time.Second)

	isDone := make(chan bool)

	go func() {
		longRunningOperations()
		isDone <- true
	}()

	// if range 20 is less than time.After its case will happens
	// else time out or after will case will execute
	select {
	case <-isDone:
		fmt.Println("Operation completed")
	case <-timeOut:
		fmt.Println("Time has been finished")

	}

	scheduleDelay()
}

// ====================schedule delay executing===============

func scheduleDelay() {
	timer := time.NewTimer(2 * time.Second) // non blocking timer starts
	fmt.Println("Starting timer")

	go func() {
		<-timer.C
		fmt.Println("Waiting")
	}()
	// to wait to exceute what in go routine
	time.Sleep(3 * time.Second) // blocking timer
	fmt.Println("Ending of program")
}

// important note
// always stop timer to avoid memory resource leaks
// ensure using defer and timer.Stop()
