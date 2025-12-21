package main

import "fmt"

func main() {
	someSlice := []string{"a", "b", "c"}
	someString := "hello"
	updateString(someString)
	updateSlice(someSlice)

	fmt.Println(someSlice)
	fmt.Println(someString)
}

func updateSlice(slice []string) {
	slice[0] = "hello"
}

func updateString(s string) {
	s = "new"
}

// Why does this work - why did it not rqueire a pointer?/
