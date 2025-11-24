package main

import "fmt"

func main() {
	// clousres use more memmory because it can capture refrence of
	// large object more than usual

	add := adder()
	fmt.Println(add(1))
	fmt.Println(add(2))

	// add2 will make reintializing so var will change and also value will changed
	add2 := adder()
	fmt.Println(add2(1))
	fmt.Println(add2(2))

	sub := subber(20)
	fmt.Println(sub(6))
	fmt.Println(sub(2))

	// anoymnous multiplier
	multiplier := func() func(int) int {
		intializer := 2
		return func(num int) int {
			intializer *= num
			return intializer
		}
	}()
	fmt.Println("multipler: ", multiplier(2))
	fmt.Println("multipler: ", multiplier(2))
	fmt.Println("multipler: ", multiplier(2))
}

func adder() func(int) int {
	i := 0
	return func(x int) int {
		i += x
		return i
	}
}

func subber(intial int) func(int) int {
	i := intial
	return func(y int) int {
		i -= y
		return i
	}
}
