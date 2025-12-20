package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	dude := person{
		firstName: "Dude",
		lastName:  "Mary",
		age:       25,
	}
	secondDude := person{"alex", "jognson", 33}

	fmt.Println(dude)
	fmt.Println(secondDude)
}
