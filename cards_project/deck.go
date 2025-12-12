package main

import "fmt"

type deck []string

// this is a function with a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
