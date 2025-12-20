package main

import "fmt"

type contactInfo struct {
	email string
	zip   string
}
type person struct {
	firstName string
	lastName  string
	age       int
	contact   contactInfo
}

func main() {
	var george person

	fmt.Println(george)
	fmt.Printf("%+v\n", george)

	george.firstName = "George"
	fmt.Printf("%+v\n", george)
	george.lastName = "Gorge"
	fmt.Printf("%+v\n", george)
	george.contact.email = "abc@def"
	fmt.Printf("%+v\n", george)
}
