package main

import "fmt"

// for global package level we cannot use := for delcaring var
// we must use var or const

// additioninfo := "new" // cannot use that
var additioninfo = "new"

func main(){
	var firstname string = "ali"
	var secondname = "ahmed"
	var lastname string
	var age int // declared with zero 
	email := "ali@gmail.com"

	// change it in run time 
	additioninfo = "old"
	
	fmt.Println("info: ", firstname, secondname, lastname, age, email, additioninfo)

	// Default Values 
	// Numeric values : 0
	// String values : ""
	// Boalen values : false
	// Pointers, Slices, Functions, Structs: nil


	// Naming Conventions:

	// PascalCase
	// Eg. CalculateArea
	// Structs, Interfaces, Enums


	// snake_case
	// Eg. db_models.go
	// files names 

	// UPPER CASE 
	// Eg. MAX_NUMBER
	// used for constants

	// chamelCase
	// Eg. addNums
	// used for variabled, functions, ..etc

	// packages names should be lowercase without underscore same name of file name

}
