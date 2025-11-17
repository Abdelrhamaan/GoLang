package main

import "fmt"

func main() {
	score := 90
	
	if score >= 90{
		fmt.Println("Grade A..")
	}else if score >= 80 {
		fmt.Println("Grade B..")
	}else if score >= 70 {
		fmt.Println("Grade C..")
	}else{
		fmt.Println("Grade D..")
	}

	if 10 % 2 == 0 || 5 % 2 ==0 {
		fmt.Println("Either 10 or 5 is even")
	} 

	if 10 % 2 == 0 && 6 % 2 ==0 {
		fmt.Println("Both 10 or 5 is even")
	} 
}