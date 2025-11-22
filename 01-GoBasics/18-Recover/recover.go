package main

import (
	"fmt"
	"os"
)

func main() {
	mySlice := []int{5, 3, 2, -4, 1}
	process(mySlice...)
	//  if i removed recover handling this line will not executed 
	fmt.Println("After calling process")


	defer fmt.Println("Defer will not executed")
	fmt.Println("Before Exit")
	os.Exit(1)
	fmt.Println("This will never executed :)")
	
}


func process (nums ...int) {

	// handle panic happening by defer and recover
	defer func ()  {
		if r := recover(); r != nil {
			fmt.Println("Sorry Num is less than zero: ", r)
		}
	}()

	for _, v := range nums {
		if v > 0 {
			fmt.Println("Num is greater than zero: ", v)
		}
		if v < 0 {
			panic(v)
		}
	}


}