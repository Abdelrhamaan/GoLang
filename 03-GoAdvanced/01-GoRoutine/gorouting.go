package main

import (
	"fmt"
	"time"
)

func main() {
	// first execute it in main thread with blocking
	// sayHello()

	// second extract it from main thread by using go keyward and non-blocking happens
	// nothing will printed
	// because main function finished before anything returns to main thread
	// go routines not stopping execution flow
	fmt.Println("Before go routine")
	go sayHello()
	fmt.Println("After go func")

	// go err := dodoError() not accepted
	var err error
	go func() {
		err = doError()
	}()
	// if err != nil {
	// 	fmt.Println("erro happens", err)
	// } else {
	// 	fmt.Println("Working Completed")
	// }

	// you will see that cahr a printed before number because go scheduler assign it to thread before num
	go printNumbers()
	go printLetters()

	//  simple way to wait go routine
	time.Sleep(3 * time.Second)

	// this way dowork func error will occur
	// but above wil not happens
	if err != nil {
		fmt.Println("erro happens", err)
	} else {
		fmt.Println("Working Completed")
	}

}

func sayHello() {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello from go routines")
}

func printNumbers() {
	iterations := 8
	for v := range iterations {
		fmt.Println("num is: ", v)
		fmt.Println(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, v := range "abcdef" {
		fmt.Println("char is: ", string(v))
		fmt.Println(200 * time.Millisecond)

	}
}

func doError() error {
	time.Sleep(time.Second)
	return fmt.Errorf("Error occure while running goroutine")
}
