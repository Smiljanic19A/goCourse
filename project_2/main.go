package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var george person

	fmt.Println(george)
	fmt.Printf("%+v", george)
}
