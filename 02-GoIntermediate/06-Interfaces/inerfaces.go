package main

import (
	"fmt"
	"math"
)

// important note :
//  any struct should implement all methods in interface to can use it after that
// if struct not implementing all methods the struct will still valid but you will not able to use it with interface

type geometry interface {
	area() float64
	premeiter() float64
}

// New interface
type describer interface {
	describe()
}

type using interface {
	use() []string
}

type circle struct {
	radius float64
}
type rect struct {
	length float64
	height float64
}

func (r rect) area() float64 {
	return r.height * r.length
}

func (r rect) premeiter() float64 {
	return 2 * (r.height + r.length)
}
func (r rect) use() []string {
	return []string{"shapes, lands, drawings"}
}

// Implement describer for rect
func (r rect) describe() {
	fmt.Printf("Rectangle [Length: %.2f, Height: %.2f]\n", r.length, r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) premeiter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func (c circle) use() []string {
	return []string{"balls", "coins"}
}

// Implement describer for circle
func (c circle) describe() {
	fmt.Printf("Circle [Radius: %.2f]\n", c.radius)
}

func main() {
	r := rect{length: 10, height: 5}
	c := circle{radius: 5}

	fmt.Println("--- Using Geometry Interface ---")
	measure(r)
	measure(c)

	fmt.Println("\n--- Using Describer Interface ---")
	printDescription(r)
	printDescription(c)

	fmt.Println("\n--- Using Using Interface ---")
	printUsing(r)
	printUsing(c)

	fmt.Println("\n--- Using Interface to print many types ---")
	myPrinter("firs", 2, 3, true, []int{})

}

// Helper function using geometry interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Printf("Area: %.2f\n", g.area())
	fmt.Printf("Perimeter: %.2f\n", g.premeiter())
}

// Helper function using describer interface
func printDescription(d describer) {
	d.describe()
}

// Helper function using use interface
func printUsing(u using) {
	fmt.Println("u interface", u)
	fmt.Printf("Type of interface is %T\n", u)
	fmt.Println("Using Interface", u.use())
}

// using interface to use any type
func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func acceptAnyKindOfValue(i interface{}) {
	fmt.Println(i)
}
