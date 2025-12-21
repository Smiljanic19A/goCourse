package main

import "fmt"

type person struct {
	name     string
	lastName string
	age      int
	contact
}

func (p person) toString() string {
	return "Name: " + p.name + "\n" + "Surname: " + p.lastName + "\n" + "Age: " + string(p.age) + "\n" + "Contact: " + p.contact.toString()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	(*p).name = name
}
