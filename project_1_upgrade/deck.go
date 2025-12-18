package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	d := deck{}
	suites := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	values := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Ace", "Jack", "Queen", "King"}

	for _, suit := range suites {
		for _, value := range values {
			d = append(d, value+" Of "+suit)
		}
	}

	return d
}

func deal(d deck, handSize int) (deck, deck, error) {
	n := len(d)
	cardsToLeave := n - handSize
	if cardsToLeave < 0 {
		return deck{}, deck{}, errors.New("deck is too small to deal")
	}

	return d[:cardsToLeave], d[cardsToLeave:], nil
}

func fromFile(filename string) (deck, error) {
	var d deck
	deckBytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return deck{}, err
	}

	cardSlice := strings.Split(string(deckBytes), "\n")
	for _, card := range cardSlice {
		d = append(d, card)
	}
	return d, nil
}

func (d deck) shuffle(timesToShuffle int) deck {

	if timesToShuffle < 1 {
		fmt.Println("shuffle requires at least 1 times to shuffle")
		return d
	}
	if len(d) < 2 {
		fmt.Println("shuffle requires at least two times to shuffle")
		return d
	}

	fmt.Println("Shuffling deck of length", len(d))

	for i := 0; i < timesToShuffle; i++ {
		for i, _ := range d {
			n := rand.Intn(len(d) - 1)
			d[i], d[n] = d[n], d[i]
		}
	}
	return d
}

func (d deck) print() {
	fmt.Println(d.toString())
}

func (d deck) toString() string {
	return strings.Join(d, "\n")
}

func (d deck) toBytes() []byte {
	return []byte(d.toString())
}

func (d deck) toFile(filename string) error {
	return os.WriteFile(filename, d.toBytes(), 0644)
}
