package main

import (
	"fmt"
	"slices"
)

func main() {
	// slice is a reference to an underlying array
	// think in slice like a struct
	var numbersSlice []int
	var numbersSlice3 = []int{}
	numbersSlice1 := []int{1, 2, 3}
	numbersSlice2 := make([]int, 5)
	_ , _, _, _= numbersSlice, numbersSlice3, numbersSlice1, numbersSlice2

	// how to make slice from array
	myarr := [6]int{1, 2, 3 ,4, 5, 6}
	mySlice := myarr[1:4]
	fmt.Println("My slice: ", mySlice)
	mySlice = append(mySlice, 7, 8)
	fmt.Println("My slice after appending: ", mySlice)
	myCopiedSlice := make([]int, len(mySlice))
	copy(myCopiedSlice, mySlice) // deep copy
	fmt.Println("My Copied slice after appending: ", myCopiedSlice)
	// Print header pointer of copied slice
    fmt.Printf("myCopiedSlice header: %p\n", &myCopiedSlice)

    // Print underlying copied slice array pointer
    fmt.Printf("my Slice data:   %p\n", &mySlice)

	// nil slice 
	myNilSlice := []int{}
	fmt.Println("Nil Slice", myNilSlice)

	// equaltiy 
	fmt.Println(slices.Equal(mySlice, myCopiedSlice))
	fmt.Println(slices.Equal(mySlice, myNilSlice))

}

