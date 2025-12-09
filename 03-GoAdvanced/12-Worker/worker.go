package main

import (
	"fmt"
	"strconv"
	"time"
)

// +--------------+         +-----------------------------+
// |  jobs chan   | ----->  |   worker goroutine (1..N)   |
// +--------------+         +-----------------------------+
//                                 |   |
//                                 v   v
//                            process jobs
//                                 |
//                                 v
//                          results channel (optional)

// <- chan recieve only channel
// chan <- send only channel
// chan bidirectional channel

// ------------ Basic worker Pool Pattern ------------------

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for jb := range jobs {
// 		fmt.Printf("Worker id is %d and job js %d\n", id, jb)
// 		// simulate working something
// 		time.Sleep(time.Second)
// 		results <- jb * 2

// 	}
// }

// func main() {
// 	jobsChan := make(chan int, 10)
// 	resultsChan := make(chan int, 10)

// 	numOfWorkers := 3
// 	numOfJobs := 6

// 	// create workers
// 	for jb := range numOfWorkers {
// 		go worker(jb+1, jobsChan, resultsChan)
// 	}

// 	// send values to task channel
// 	for jb := range numOfJobs {
// 		fmt.Println("Sending data in jobs")
// 		jobsChan <- jb + 1
// 	}
// 	close(jobsChan)

// 	// consume tasks from results chan
// 	for range numOfJobs {
// 		fmt.Println("recieving data in results", <-resultsChan)

// 	}

// }

// create world real scenario
// create workers
type Order struct {
	id       int
	price    int
	mealName string
}

func workerChefs(orderChan <-chan Order, deliveringChan chan<- int) {
	for ord := range orderChan {
		fmt.Printf("order id %d, price %d, meal %s\n", ord.id, ord.price, ord.mealName)
		// simualte making meals
		time.Sleep(2 * time.Second)

		// send meals to delivering channels

		deliveringChan <- ord.id

	}
}

func main() {
	numOfChefs := 10
	numOfOrders := 100
	chanCapacity := 50
	price := 20
	mealBaseName := "Meal"

	orderChan := make(chan Order, chanCapacity)
	deliveringChan := make(chan int, chanCapacity)

	for range numOfChefs {
		go workerChefs(orderChan, deliveringChan)
	}

	// creating orders and sending them to chefs in order chan
	for num := range numOfOrders {
		orderId := num + 1
		orderChan <- Order{id: orderId, price: price, mealName: mealBaseName + strconv.Itoa(orderId)}
	}
	close(orderChan)

	// consuming it in reuslt channel
	for range numOfOrders {
		<-deliveringChan
	}
}
