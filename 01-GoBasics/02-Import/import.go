package main

import (
	"fmt"
	// "net/http"
	// we can name package like
	foo "net/http"
)

func main (){
	fmt.Println("hello go standard library:)")
	// res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	res, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error: ", err)
		return 
	}
	defer res.Body.Close()
	fmt.Println("response body is ", res)
}
