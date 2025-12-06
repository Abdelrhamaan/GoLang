package main

import (
	"fmt"
)

type Employee struct {
	Id     int64
	Name   string
	Age    int16
	Salary int64
}

func main() {
	// how to create channel
	myChan := make(chan string)
	//  wrong usage of channels to learn
	// you cannot make channel inside any function you must make goroutine to send and rec msgs
	// myChan <- "Hello from channels!"

	// receiver := <-myChan
	// fmt.Println("Recieved Msg: ", receiver)

	// right way to wrap putting mssages in channels with goroutines
	go func() {
		myChan <- "Hello from channels!"
	}()

	receiver := <-myChan
	fmt.Println("Recieved Msg: ", receiver)
	fmt.Println("End!!!")

	// make channels of type struct
	employees := map[int][]Employee{
		1: {
			{
				Id:     1,
				Name:   "Ali",
				Age:    25,
				Salary: 60000,
			},
		},
		2: {
			{
				Id:     2,
				Name:   "Noor",
				Age:    29,
				Salary: 70000,
			},
		},
	}

	empChan := make(chan []Employee)

	go func() {
		for _, v := range employees {
			empChan <- v
		}
		close(empChan)
	}()
	for emp := range empChan {
		fmt.Println("Received:", emp)
	}

}

// func main() {
// 	// make channels of type struct
// 	employees := map[int][]Employee{
// 		1: {
// 			{
// 				Id:     1,
// 				Name:   "Ali",
// 				Age:    25,
// 				Salary: 60000,
// 			},
// 		},
// 		2: {
// 			{
// 				Id:     2,
// 				Name:   "Noor",
// 				Age:    29,
// 				Salary: 70000,
// 			},
// 		},
// 	}

// 	empChan := make(chan []Employee)

// 	go func() {
// 		for _, v := range employees {
// 			empChan <- v
// 		}
// 		close(empChan)
// 	}()
// 	go func() {

// 		for emp := range empChan {
// 			fmt.Println("Received:", emp)
// 		}
// 	}()
// 	// or use wait group to perevent main thread from finishing exceution and waiting ro goroutines
// 	// if you will not use wait group memory leaks will happens
// 	time.Sleep(time.Second)
// 	fmt.Println("End of Program!!")

// }
