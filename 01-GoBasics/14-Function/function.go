package main

import (
	"errors"
	"fmt"
)

func main() {
	
	// paramter passed in function is copy not real one
	fmt.Println(add(1, 3))

	c := 5 

	fmt.Println(editNum(c))
	fmt.Println(c)
	fmt.Println(editNum2(&c)) // EDIT function to take pointer 
	fmt.Println(c)


	//  anyomouns function 
	func (){
		fmt.Println("Hello from anoyomonus")
	}()

	hi := func(){
		fmt.Println("Assign anyomouns func to hi")
	}
	hi()
	mapFunc := map[string]func(){
		"start": hi,
		"end": func() {
			fmt.Println("finished")
		},
	}
	
	for k, v := range mapFunc {
		fmt.Printf("key is %s → executing function:\n", k)
		v() // call function separately
	}

	result := createOperation(5, 6, add)
	fmt.Println("result is ", result)
	

	multiple := takesTwoLevelParams(2)
	fmt.Println("passing two multiple levels", multiple(6))

	comapreWithError, err := compare(3, 3)

	if err != nil {
		fmt.Println("Error", err)
	}else{
		fmt.Println(comapreWithError)
	}
}


func add (a, b int) int  {
	return  a + b 
}


func editNum(c int) int{
	c -= 1
	return  c
}

func editNum2(c *int) int{
	*c -= 1
	return  *c
}
func createOperation (a int, b int, fn func(int, int) int) int {
	return  fn(a, b)
}


func takesTwoLevelParams(num int) func(x int) int{
	return func(x int) int {
		return  num * x
	}
}

func compare(x, y int) (string, error){
	if x > y {
		return "x greater than y", nil
	}

	if y > x {
		return "y greater than x", nil
	}
	return  "", errors.New("Unable to compare two numbers")
}





// What does “first-class citizen” mean? important

// 	A first-class value can:

// 	Be assigned to a variable

// 	Be passed as an argument

// 	Be returned from a function

// 	Be stored in a data structure

// 	Be created at runtime



