package main

import "fmt"

type deck []string

// this is a function with a receiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {

	d := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}

	return d
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // return from start up to handsize and from the end all the way to handsize
}
