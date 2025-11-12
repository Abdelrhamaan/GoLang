package main

import "fmt"

func main() {
	multiple := 5

	for {
		multiple *= 2
		fmt.Println("Multiple: ", multiple)
		if multiple >= 50 {
			break
		}
	}


	num := 1

	for num < 11 {
		if num % 2 == 0 {
			fmt.Println("Even Num", num)
		}
		num++
	}
}