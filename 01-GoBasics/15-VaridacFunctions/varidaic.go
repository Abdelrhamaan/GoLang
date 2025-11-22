package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Ellipses (...)  ---> pass zer  or nums of argument to func
	// varidic params must be last thing in params in func


	fmt.Println(total(3, 4, 5, 6))

	str, sum := totaWithString( "the sum of nums is: ", 3, 4, 5)
	fmt.Println(str, sum)


	// passing slices and making unpacking 
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	str2, sum2 := totaWithString("passing slices and sum nums in it: ", mySlice...)
	fmt.Println(str2, sum2)

}

func total(nums ...int) int  {
	t := reflect.TypeOf(nums)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	total := 0
	for _, v := range nums {
		total += v
	}
	return  total


	
}


func totaWithString(str string, nums ...int) (string, int)  {
	t := reflect.TypeOf(nums)
	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	total := 0
	for _, v := range nums {
		total += v
	}
	return str, total
}