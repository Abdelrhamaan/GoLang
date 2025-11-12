package main

import (
	"fmt"
	"math"
)


func main (){

	var a, b int = 10, 3

	// addition 
	result := a + b
	fmt.Println("Addition: ", result)
	

	// subtraction
	result = a - b
	fmt.Println("subtraction: ", result)


	// multiplication
	result = a * b
	fmt.Println("multiplication: ", result)


	// division
	result = a / b
	fmt.Println("division: ", result)


	// modulus - reminder
	result = a % b
	fmt.Println("modulus: ", result)



	// float take care of that 

	const pi float32 = 22 / 7
	fmt.Println("not floated: ", pi)	


	const pi1 float32 = 22 / 7.0
	fmt.Println("floated: ", pi1)	



	// Overflow with signed integer
	var maxInt int8 = 127 // max value of int8
	fmt.Println("Before overflow (int8):", maxInt)

	maxInt += 1 // overflow happens here
	fmt.Println("After overflow (int8):", maxInt) // becomes -128

	// Overflow with unsigned integer
	var umaxInt uint8 = 255 // max value of uint8
	fmt.Println("Before overflow (uint8):", umaxInt)

	umaxInt += 1 // overflow happens here
	fmt.Println("After overflow (uint8):", umaxInt) // becomes 0



	


	// Start with a very small float number
	var f float64 = 1e-300
	fmt.Println("Start:", f)

	// Repeatedly divide it to force underflow
	for i := 0; i < 10; i++ {
		f /= 10
		// f /= math.MaxFloat64
		fmt.Printf("Step %d: %.100e\n", i+1, f)
	}

	// Check if it became zero
	if f == 0 {
		fmt.Println("Underflow occurred — number became 0")
	} else if math.IsInf(f, 0) {
		fmt.Println("Overflow occurred — number became Inf")
	} else {
		fmt.Println("Still representable (subnormal number):", f)
	}




}
