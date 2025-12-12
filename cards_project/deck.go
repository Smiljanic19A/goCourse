package main

import "fmt"
import "math/rand/v2"

type deck []string

// this is a function with a receiver
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	combinations := []string{
		// Hearts
		"AH", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "JH", "QH", "KH",
		// Diamonds
		"AD", "2D", "3D", "4D", "5D", "6D", "7D", "8D", "9D", "10D", "JD", "QD", "KD",
		// Clubs
		"AC", "2C", "3C", "4C", "5C", "6C", "7C", "8C", "9C", "10C", "JC", "QC", "KC",
		// Spades
		"AS", "2S", "3S", "4S", "5S", "6S", "7S", "8S", "9S", "10S", "JS", "QS", "KS",
	}

	d := deck{}

	i := 0

	for i < len(combinations) {
		randomInt := rand.IntN(len(combinations))
		d = append(d, combinations[randomInt])
		combinations = append(combinations[:randomInt], combinations[randomInt+1:]...)
	}

	return d

}
