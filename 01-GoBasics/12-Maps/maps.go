package main

import (
	"fmt"
	"maps"
)

func main() {
	myMap := make(map[string]int)
	fmt.Println("myMap: ", myMap)
	myMap["ALI"] = 35
	myMap["Mahmoud"] = 35
	myMap["Karim"] = 35
	fmt.Println("myMap: ", myMap)
	// if key not found he will print false value of value type 
	// ex: string "", int 0    ..etc
	fmt.Println("myMap: ", myMap["nader"])

	delete(myMap, "ALI")
	fmt.Println("myMap: ", myMap)

	clearedMap := map[string]int{
		"Python": 50,
		"Js": 30,
		"GO": 10,
	}
	fmt.Println("clearedMap: ", clearedMap)
	clear(clearedMap)
	fmt.Println("clearedMap: ", clearedMap)


	// if key not in map you will have two values returned 0 and false
	val, ok := myMap["Karim"]
	val1, ok1 := myMap["ALI"]
	fmt.Println("ok: ", val, ok)
	fmt.Println("ok: ", val1, ok1)


	if !maps.Equal(myMap, clearedMap) {
		fmt.Println("two maps are not equal")
	}
	//  iteration 
	for k, v := range myMap {
		fmt.Println(k, v)
	}


	// nil map value 
	// map is intialized with nil 
	var nilMap map[string]map[string]int // nested 
	// var nilMap map[string]map[string]int
	if nilMap == nil{
		fmt.Println(" a nil empty map", nilMap)
	}

	mapOfSlices := map[string][]string{
		"courses": {"Js", "C++",},
	}
	fmt.Println("mapOfSlices", mapOfSlices)

	//  intitalizing with make
	makeNestedMap := make(map[string]map[string]int)

    makeNestedMap["users"] = map[string]int{
        "ali":   32,
        "ahmed": 25,
    }

    fmt.Println(makeNestedMap)
    fmt.Println(len(makeNestedMap))
}