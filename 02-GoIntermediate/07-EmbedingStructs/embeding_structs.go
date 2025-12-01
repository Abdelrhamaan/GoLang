package main

import "fmt"

// make inheritance using embedding

type animal struct {
	name string
	age  int
}

type dog struct {
	// this is unnamed field if you named fields of animal will not be promoted to dog struct
	animal
	kind  string
	isRun bool
}

func (a animal) makeSound() {
	fmt.Println("animal sound")
}

// override method in parent struct by making same in the child class
func (d dog) makeSound() {
	fmt.Println("HoHO")
}

func main() {
	myDog := dog{
		animal: animal{name: "jack", age: 2},
		kind:   "booldog",
		isRun:  true,
	}

	fmt.Println("myDog: ", myDog.name)
	fmt.Print("myDog: ")
	myDog.makeSound()

}
