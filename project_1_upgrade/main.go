package main

import "fmt"

//requirements
// create deck done
// deal hand
// write to file
// read from file

func main() {
	d := newDeck()

	remainingDeck, _, err := deal(d, 5)

	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 10; i++ {
		var hand deck

		remainingDeck, hand, err = deal(remainingDeck, 5)

		if err != nil {
			fmt.Println(err)
			break
		}
		err = hand.toFile("hand.txt")
		if err != nil {
			fmt.Println("Error writing to file")
			fmt.Println(err)
		}
		fmt.Println("New deck size", len(remainingDeck), "hand size", len(hand))
	}

	randomDeck, err := fromFile("hand.txt")
	if err != nil {
		fmt.Println(err)
	}

	randomDeck.print()
}
