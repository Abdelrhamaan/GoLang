package main

import "fmt"

func init (){
	fmt.Println("first init function ")
}
func init (){
	fmt.Println("second init function ")
}
func init (){
	fmt.Println("third init function ")
}


func main() {
	fmt.Println("Print main function ")
}