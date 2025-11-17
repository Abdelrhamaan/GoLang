package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nums [6] int
	fmt.Println(nums)
	nums[5] = 3
	fmt.Println(nums)
	nums[len(nums) - 1] = 1
	fmt.Println(nums)
	nums[2] = 3
	fmt.Println(nums)


	employees := [3] string {
		"ali", "ahmed", "noor",
	}
	fmt.Println(employees)


	originalArray := [3]int{3, 4, 5}
	copiedArray := originalArray
	originalArray[0] = 252

	fmt.Println("originalArray: ", originalArray)
	fmt.Println("copiedArray: ", copiedArray)


	for i:=0; i < len(originalArray); i++{
		fmt.Println("Elemnt Index", i, "Value: ", originalArray[i])
	}

	for index, value := range copiedArray{
		fmt.Println("Element is: ", value, "Its Index is: ", index)
	}

	a, _ := someFucntion()
	fmt.Println("a without b", a)


	// to prevent it from decalring an error 
	b := 5
	_ = b

	// thats wrong because you can use const with array 
	// it must be numbers, string, runes, bool
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	matrix1 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	_ = matrix
	_ = matrix1


	// lets try copy array with pointers

	arr1 := [5]int{
		1, 2, 3, 4, 5,
	}
	// point for an empty array with 5 int numbers (pointer variable)
	var arr2 *[5]int // when you decalre that you decalre pointers so you can not compare two arrays

	// make refrence equal refrence
	arr2 = &arr1
	arr2[4] = 100
	fmt.Println("Array1 Type: ", reflect.TypeOf(arr1))
	fmt.Println("Array2 Type: ", reflect.TypeOf(arr2))
	fmt.Println(arr1 == *arr2)
	fmt.Println("Array1 is:", arr1)
	fmt.Println("Array2 is:", arr2) // also here give you refrence to the array
	fmt.Println("Array2 is:", &arr2) // get pointer or refrence in memmory
	fmt.Println("Array2 is:", *arr2) // get original array values



}


func someFucntion()(int, int){
	return  3, 4
}