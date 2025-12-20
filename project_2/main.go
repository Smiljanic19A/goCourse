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
	fmt.Printf("%+v\n", george)

	george.firstName = "George"
	fmt.Printf("%+v\n", george)
	george.lastName = "Gorge"
	fmt.Printf("%+v\n", george)
}
