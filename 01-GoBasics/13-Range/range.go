package main

import "fmt"

func main() {
	myStr := "Hello World!"

	for i, v := range myStr{
		// fmt.Println("index", i, "Unicode", v)
		fmt.Printf("Index is %d, unicode is %d and rune is %c\n", i , v, v)
	}


	nums := []int{1, 2, 3, 4}

	for _, v := range nums {
		v = v * 10 // WRONG: v is a copy, does not affect nums
	}
    fmt.Println("Not Updated slice:", nums)

    // You must use the index to edit slice elements
    for i := range nums {
        nums[i] = nums[i] * 10 // Modify the slice directly
    }

    fmt.Println("Updated slice:", nums)



	ages := map[string]int{
        "Ali":   20,
        "Sara":  25,
        "Omar":  30,
    }

    // You can modify map values during iteration
    for key, value := range ages {
        ages[key] = value + 10 // updating value
    }

    fmt.Println("Updated map:", ages)
}