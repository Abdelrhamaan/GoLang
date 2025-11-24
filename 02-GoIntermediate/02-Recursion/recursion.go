package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factorial(3))
}

func factorial(n int) int {
	// 01-Base case is when n == 0 return 1
	if n == 0 {
		return 1
	}
	// 02-Recursive case return (factorial -1) * n
	return factorial(n-1) * n

}
