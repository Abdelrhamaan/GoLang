package main

import "fmt"

func main() {
	carType := "Baic u5"
	switch carType {
	case "mercedes":
		fmt.Println("Car Type is mercedes")
	case "toyota":
		fmt.Println("Car Type is toyota")
	case "hyndai":
		fmt.Println("Car Type is hyndai")
	default:
		fmt.Println("Unkonw type")
	}

	fruit := "banana"
	switch fruit {
	case "apple", "orange":
		fmt.Println("good")
	case "watermilon", "banana":
		fmt.Println("prefect")
	}

	number := 15
	// when using expression not use var with swtich
	switch  {
	case number < 5 :
		fmt.Println("Number is too little")
	case number > 5 && number < 10:
		fmt.Println("I think it good")
	default:
		fmt.Println("i think it prefect")
	}


	// using fall through any case is satsfied will execute
	// not just the case 
	// i mean cases can be executed

	num := 2

	switch {
	case num > 1:
		fmt.Println("Greater than 1...")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	}

	checkType(5)
	checkType(5.0)
	checkType("ssss")
	checkType(true)

}



func checkType(x interface{}){
	// we cannot use fallthrough with type 
	switch x.(type){
	case int:
		fmt.Println("Its an Integer")
	case string:
		fmt.Println("Its a string")
	case float64:
		fmt.Println("Its a float")
	default:
		fmt.Println("Unkown")
	}	
}