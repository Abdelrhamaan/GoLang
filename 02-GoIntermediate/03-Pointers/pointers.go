package main

import "fmt"

func main() {
	// 1. A normal variable
	x := 10
	fmt.Println("Value of x:", x)

	// 2. Getting the address (Pointer)
	// The & operator returns the memory address of a variable
	ptr := &x
	fmt.Println("Address of x (ptr):", ptr)

	// 3. Dereferencing
	// The * operator lets us access the value at that address
	fmt.Println("Value at address ptr:", *ptr)

	// 4. Modifying via pointer
	// We can change the value of x by assigning to *ptr
	*ptr = 20
	fmt.Println("New value of x:", x) // x is now 20!

	// 5. Passing by reference (using pointers in functions)
	changeValue(ptr)
	fmt.Println("Value after function call:", x)

	tryPointer()

	// zero value of pointer is nil
	var ptr2 *bool
	var ptr3 *map[int]string
	var ptrArray *[]int
	fmt.Println("ptr2", ptr2)
	fmt.Println("ptr3", ptr3)
	fmt.Println("ptrArray", ptrArray)

	if ptrArray == nil {
		myArr := []int{1, 2, 3, 4}
		ptrArray = &myArr

		fmt.Println("myArr", myArr)
		fmt.Println("ptrArray", ptrArray)
		fmt.Println("depointer", *ptrArray)

		// changing two new arr
		newArr := []int{5, 6, 7}
		*ptrArray = newArr

		fmt.Println("newArr", newArr)
		fmt.Println("ptrArray", ptrArray)
		fmt.Println("depointer", *ptrArray)
	}

	myStr := "Ali"
	myPtr := &myStr
	fmt.Println("myStr", myStr)

	modifyString(myPtr, "Noor")

	fmt.Println("myStr", myStr)

}

func changeValue(num *int) {
	*num = 100
}

func tryPointer() {
	var ptr *int
	num := 5
	ptr = &num
	fmt.Println("num", num)
	fmt.Println("real memory of num:", &num)
	fmt.Println("real memory of ptr:", &ptr)
	fmt.Println("ptr", ptr)
	fmt.Println("real num", *ptr)

	// change num
	*ptr = 20

	fmt.Println("num", num)
	fmt.Println("ptr", ptr)
	fmt.Println("real num", *ptr)
}

func modifyString(ptrString *string, newValue string) {
	*ptrString = newValue
}
