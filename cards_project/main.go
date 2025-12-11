package main

import "fmt"

func main() {
	//var card string = "Ace Of Spades"
	card := determineCardValue()

	fmt.Println(card)
}

func determineCardValue() string {
	return "Ace Of Spades"
}
