package main

import "fmt"

func main() {
	someSlice := []string{"a", "b", "c"}
	updateSlice(someSlice)

	fmt.Println(someSlice)
}

func updateSlice(slice []string) {
	slice[0] = "hello"
}

// Why does this work - why did it not rqueire a pointer?/
