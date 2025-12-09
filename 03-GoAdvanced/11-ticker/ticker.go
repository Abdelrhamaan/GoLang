package main

import (
	"fmt"
	"time"
)

// ==========Schedul logging and predoic task and polling for updates every preodic task==============

// func main() {
// 	// it will run every 5 seconds and will stop just in case main function finished
// 	// newTicker := time.NewTicker(5 * time.Second)

// 	// defer newTicker.Stop()

// 	// for tick := range newTicker.C {
// 	// 	fmt.Println("Tick is: ", tick)
// 	// }

// 	newTicker := time.NewTicker(2 * time.Second)
// 	defer newTicker.Stop()
// 	for {
// 		select {
// 		case <-newTicker.C:
// 			pridoicTask()
// 		}
// 	}
// }

// func pridoicTask() {
// 	fmt.Println("Task is run every specific time", time.Now())
// }

// ===========Stopping ticker after sometimes=======================

func main() {
	// dont forget to stop the ticker
	newTicker := time.NewTicker(time.Second)
	stop := time.After(6 * time.Second)
	defer newTicker.Stop()

	fmt.Println("Starting....")

	for {
		select {
		case tick := <-newTicker.C:
			fmt.Println("Tick at:", tick)
		case <-stop:
			fmt.Println("Stopped.......")
			return
		}
	}
}
