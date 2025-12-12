package main

import "math/rand/v2"

func main() {
	cards := deck{determineCardValue()}
	cards = append(cards, makeRandomCard())
	cards.print()
}

func determineCardValue() string {
	return "Ace Of Spades"
}

func makeRandomCard() string {
	randomInt := rand.IntN(3) + 1
	if randomInt == 1 {
		return "ACE OF SPADES"
	} else if randomInt == 2 {
		return "ACE OF DIAMONDS"
	}
	return "King Of Hearts"
}
