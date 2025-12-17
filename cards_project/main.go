package main

import "fmt"

func main() {
	cards := newDeck()
	//cards.print()

	hand, _ := deal(cards, 5)

	//hand.print()
	//remainingDeck.print()
	//test := "hello there"

	//cardString := toString(hand)
	fmt.Println(hand.toBytes())
}
