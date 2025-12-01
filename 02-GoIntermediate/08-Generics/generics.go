package main

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

// different between any and comaprable is that compareable to check == but any cannot do that
// once we creat a stack from any type it will not accept any another types
type Stack[T comparable] struct {
	elements []T
}

func (s *Stack[T]) push(elem T) {
	s.elements = append(s.elements, elem)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) isEmpty() bool {
	if len(s.elements) == 0 {
		return true
	}
	return false
}

func (s *Stack[T]) getElemByIndex(i int) (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	return s.elements[i], true
}

func (s Stack[T]) isElemIn(elem T) (T, int, bool) {
	for i, v := range s.elements {
		if v == elem {
			return elem, i, true
		}
	}
	var zero T
	return zero, 0, false
}

func (s Stack[T]) printAll() {
	for _, v := range s.elements {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func main() {
	// 1. Swap Example
	first, second := "a", "b"
	first, second = swap("a", "b")
	fmt.Println("Swapped:", first, second)

	// 2. Stack Example
	fmt.Println("\n--- Stack Example ---")
	myStack := Stack[int]{}

	// Push
	fmt.Println("Pushing 10, 20, 30")
	myStack.push(10)
	myStack.push(20)
	myStack.push(30)

	// Print All
	fmt.Print("Current Stack: ")
	myStack.printAll()

	// Is Empty
	fmt.Println("Is Stack Empty?", myStack.isEmpty())

	// Get Element By Index
	val, ok := myStack.getElemByIndex(1)
	if ok {
		fmt.Println("Element at index 1:", val)
	}

	// Is Element In
	elem, idx, found := myStack.isElemIn(20)
	elem1, idx1, found1 := myStack.isElemIn(100)
	if found {
		fmt.Printf("Element %d found at index %d\n", elem, idx)
	} else {
		fmt.Println("Element not found")
	}
	if found1 {
		fmt.Printf("Element %d found at index %d\n", elem1, idx1)
	} else {
		fmt.Println("Element not found")
	}

	// Pop
	popped, ok := myStack.pop()
	if ok {
		fmt.Println("Popped:", popped)
	}

	fmt.Print("Stack after pop: ")
	myStack.printAll()
}
