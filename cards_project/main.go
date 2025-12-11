package main

import "fmt"
import "math/rand/v2"

func main() {
	cards := []string{determineCardValue()}

	cards = append(cards, makeRandomCard())
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
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
