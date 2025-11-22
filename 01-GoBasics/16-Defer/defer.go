package main

import "fmt"

// | Feature                     | Description                         |
// | --------------------------- | ----------------------------------- |
// | Runs after function returns | ✔ always                            |
// | Execution order             | LIFO                                |
// | Arguments evaluated early   | ✔                                   |
// | Common use                  | Close file/db, unlock mutex, timing |


func main() {
	process(10)
}



func process(x int){
	fmt.Println("Process Started")
	defer fmt.Println("Deffered x is: ", x)
	defer fmt.Println("defer 1")  // defer function is evaluated when seen in function 
	defer fmt.Println("defer 2") // but executed after returning of func
	x++
	fmt.Println("x is: ", x)

	fmt.Println("Process Finished")
}