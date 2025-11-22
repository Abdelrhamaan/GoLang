package main

import "fmt"

func main() {
	process(10)
	fmt.Println("==============================")
	process(-10)
	
}

func process(num int) {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	if num < 0{
		defer fmt.Println("defer 3")
		panic("num cannot be less than zero !")
		defer fmt.Println("defer 4")

	}
	
	defer fmt.Println("defer 5")
	if num > 0 {
		fmt.Println("Num is greater than zero ")
	}
}