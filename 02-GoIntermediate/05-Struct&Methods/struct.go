package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	firstName  string
	middleName string
	lastName   string
	age        int
	salary     float64
	position   string
	userData   userData
	// embeding anyomnous field
	moreData
}

type userData struct {
	country   string
	city      string
	email     string
	telephone string
}

// embeding anyomnous field
type moreData struct {
	writtenData string
	notes       string
	tags        []string
	isActive    bool
}

func (p Person) getFullName() string {
	return p.firstName + " " + p.lastName
}
func (a Person) getFullUserAddress() string {
	return a.userData.country + "/" + a.userData.city
}

func (s Person) applyBonusOnSalaryWithoutPointer() float64 {
	s.salary *= 1.25
	return s.salary
}

// we use pointer reciever instead of value(copy) reciever for two reasons:
// 01- to edit original field of struct
// 02- to prevent copying large structs
func (s *Person) applyBonusOnSalaryWithPointer() float64 {
	s.salary *= 1.25
	return s.salary
}

func main() {

	newPerson := Person{
		firstName: "Mohamed",
		lastName:  "Ali",
		age:       25,
		salary:    50000,
	}
	newPerson.userData.country = "Egypt"
	newPerson.userData.city = "Beni-suef"

	// Using embedded anonymous field - ALL fields are promoted ro parent struct!
	// You can access them directly without going through moreData
	newPerson.writtenData = "This is some additional data about the person"
	newPerson.notes = "Employee of the month"
	newPerson.tags = []string{"developer", "go-expert", "team-lead"}
	newPerson.isActive = true

	// Or you can also access them explicitly through the type name
	newPerson.moreData.writtenData = "This is some additional data about the person"
	fmt.Println("person ", newPerson.firstName)
	fmt.Println("person ", newPerson.lastName)
	fmt.Println("person ", newPerson.salary)
	fmt.Println("person ", newPerson.position)
	fmt.Println("person ", newPerson.userData)
	fmt.Println("person ", newPerson.getFullUserAddress())
	fmt.Println("person ", reflect.TypeOf(newPerson.position))

	fmt.Println("full name", newPerson.getFullName())

	fmt.Println("====================================")
	// Demonstrating embedded anonymous field access - ALL fields are promoted!
	fmt.Println("Written data (direct access):", newPerson.writtenData)
	fmt.Println("Notes (direct access):", newPerson.notes)
	fmt.Println("Tags (direct access):", newPerson.tags)
	fmt.Println("Is Active (direct access):", newPerson.isActive)
	fmt.Println("Written data (explicit access):", newPerson.moreData.writtenData)
	fmt.Println("Full moreData struct:", newPerson.moreData)

	fmt.Println("====================================")
	// edit with and without pointers
	fmt.Println("Original salary:", newPerson.salary)

	fmt.Println("Calling applyBonusOnSalaryWithoutPointer():", newPerson.applyBonusOnSalaryWithoutPointer())
	fmt.Println("Salary after WITHOUT pointer:", newPerson.salary) // Still 50000!

	fmt.Println("Calling applyBonusOnSalaryWithPointer():", newPerson.applyBonusOnSalaryWithPointer())
	fmt.Println("Salary after WITH pointer:", newPerson.salary) // Now 62500!
	fmt.Println("====================================")

	// we have anymonus struct like anymonus func
	car1 := struct {
		carType  string
		carModel int
		carPrice float64
		carColor string
	}{
		carType:  "Mercedes",
		carModel: 2025,
		carPrice: 2_000_000,
		carColor: "Red",
	}
	fmt.Println("car", car1)
	fmt.Println("car", reflect.TypeOf(car1))
	fmt.Println("car", car1.carModel)
	fmt.Println("car", car1.carPrice)
	fmt.Println("====================================")

	// Methods can be associated with any custom type, not just structs!
	var num myInt = 42
	var negNum myInt = -10
	fmt.Println("Is 42 positive?", num.isPositive())
	fmt.Println("Is -10 positive?", negNum.isPositive())
	fmt.Println("Tring new method", negNum.welcomeMessage())

	var greeting myStr = "hello world"
	fmt.Println("Original string:", greeting)
	fmt.Println("Uppercase (manual):", greeting.toUpper())
	fmt.Println("Uppercase (built-in):", greeting.toUpperBuiltIn())
	fmt.Println("Length:", greeting.strLen())

	fmt.Println("====================================")
	// Demonstrating more built-in string methods
	var text myStr = "  Go Programming Language  "
	fmt.Println("Original text:", text)
	fmt.Println("Lowercase:", text.toLowerBuiltIn())
	fmt.Println("Trimmed:", text.trimSpaces())
	fmt.Println("Contains 'Programming'?", text.containsSubstr("Programming"))
	fmt.Println("Replace 'Go' with 'Golang':", text.replaceStr("Go", "Golang"))

	var sentence myStr = "apple,banana,orange,grape"
	fmt.Println("Original sentence:", sentence)
	fmt.Println("Split by comma:", sentence.splitBy(","))

}

// Custom type based on int
type myInt int

// 1. NAMED VALUE RECEIVER (n myInt)
// - You give the receiver a name (n)
// - You CAN access the value using the name
// - Works on a COPY of the value
func (n myInt) isPositive() bool {
	return n > 0 // Can use 'n' to access the value
}

// 2. ANONYMOUS VALUE RECEIVER (myInt)
// - No name given to the receiver
// - You CANNOT access the value (no variable to reference it)
// - Useful when you don't need the value, just want to attach a method to the type
func (myInt) welcomeMessage() string {
	return "Welcome in myInt Type" // Cannot access the actual number value!
}

// 3. POINTER RECEIVER (*myInt)
// - Receiver is a POINTER (notice the *)
// - You CAN modify the original value (not a copy)
// - Use when you need to change the value
func (s *myInt) changingNum() int {
	*s = *s * 2 // Multiply the original value by 2
	return int(*s)
}

// Custom type based on string
type myStr string

// Method to convert string to uppercase
func (s myStr) toUpper() string {
	result := ""
	for _, char := range s {
		// Convert lowercase letters to uppercase
		if char >= 'a' && char <= 'z' {
			result += string(char - 32) // 32 is the difference between 'a' and 'A' in ASCII
		} else {
			result += string(char)
		}
	}
	return result
}

// Method to get string length
func (s myStr) strLen() int {
	return len(s)
}

// ========== Using Go's built-in strings package ==========

// Method to convert string to uppercase using built-in function
func (s myStr) toUpperBuiltIn() string {
	return strings.ToUpper(string(s))
}

// Method to convert string to lowercase using built-in function
func (s myStr) toLowerBuiltIn() string {
	return strings.ToLower(string(s))
}

// Method to check if string contains a substring
func (s myStr) containsSubstr(substr string) bool {
	return strings.Contains(string(s), substr)
}

// Method to replace occurrences in string
func (s myStr) replaceStr(old, new string) string {
	return strings.ReplaceAll(string(s), old, new)
}

// Method to split string by separator
func (s myStr) splitBy(sep string) []string {
	return strings.Split(string(s), sep)
}

// Method to trim whitespace
func (s myStr) trimSpaces() string {
	return strings.TrimSpace(string(s))
}
