package main

import (
	"fmt"
	"sync"
	"time"
)

// --------- Basic wg without channels ------------
// func worker(id int, wg *sync.WaitGroup) {
// 	// decrement wg by one
// 	defer wg.Done()
// 	fmt.Printf("Workers with id %d started\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Workers with id %d finished\n", id)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	numOfWorkers := 3
// 	/*
// 		if we make wg.Add() inside worker then logic this will not wait untill goroutines execute there logic
// 		because wg must be added before any goroutine started
// 		goroutines are very fast so if it started before add mechanism main thread will execution first before find,
// 		anything added in wg..
// 	*/
// 	wg.Add(numOfWorkers) // add three wait groups

// 	for i := range numOfWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait() // this block execution until goroutines finishing

// 	// this part will not executed until workers finished
// 	fmt.Println("main func finished")
// }

// ------------ wait group with just sending channel ----------------

// func worker(id int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker with id %d started...\n", id)
// 	time.Sleep(2 * time.Second)
// 	results <- id * 2
// 	fmt.Printf("Worker with id %d finished...\n", id)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	numOfWOrkers := 3
// 	numOfJobs := 3
// 	results := make(chan int, numOfJobs)

// 	wg.Add(numOfWOrkers)

// 	for i := range numOfWOrkers {
// 		go worker(i, results, &wg)
// 	}

// 	// we used wg.wait in go routine to non blocing getting data or results from channel
// 	go func() {
// 		wg.Wait()      // returns as soon as every worker called wg.Done()
// 		close(results) // closes the channel
// 	}()

// 	for res := range results {
// 		fmt.Println("result from channel", res)
// 	}

// }

// ------------ wait group with receiving and sending channels ----------------

// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker with id %d started...\n", id)
// 	time.Sleep(2 * time.Second)
// 	for task := range tasks {
// 		results <- task * 2
// 	}
// 	fmt.Printf("Worker with id %d finished...\n", id)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	numOfWOrkers := 5
// 	numOfJobs := 50
// 	results := make(chan int, numOfJobs)
// 	tasks := make(chan int, numOfJobs)

// 	wg.Add(numOfWOrkers)

// 	for i := range numOfJobs {
// 		tasks <- i
// 	}

// 	close(tasks)

// 	for i := range numOfWOrkers {
// 		go worker(i, tasks, results, &wg)
// 	}

// 	// we used wg.wait in go routine to non blocing getting data or results from channel
// 	go func() {
// 		wg.Wait()      // returns as soon as every worker called wg.Done()
// 		close(results) // closes the channel
// 	}()

// 	for res := range results {
// 		fmt.Println("result from channel", res)
// 	}

// }

//--------------- wait group with struct ---------------------

type WorkerTask struct {
	Id   int
	Name string
}

func (t *WorkerTask) performTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task with id %d and name %s started..\n", t.Id, t.Name)
	time.Sleep(time.Second)
	fmt.Printf("Task with id %d and name %s finished..\n", t.Id, t.Name)

}

func main() {
	var wg sync.WaitGroup
	var tasks = []string{"cleaning", "insights", "prediction"}
	for i, task := range tasks {
		wg.Add(1)
		worker := WorkerTask{Id: i, Name: task}
		go worker.performTask(&wg)
	}

	wg.Wait()
	fmt.Println("Program finished...")

}
