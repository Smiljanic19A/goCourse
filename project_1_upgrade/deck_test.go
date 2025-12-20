package main

import (
	"math/rand"
	"os"
	"testing"
)

const ExpectedNumberOfCards = 56
const IoFileName = "DeckOfCards.txt"

func cleanupAfterTests(t *testing.T) {

	_, err := os.Stat(IoFileName)

	if err == nil {
		err := os.Remove(IoFileName)
		if err != nil {
			t.Errorf("Error removing %s file", IoFileName)
		}
	}
}
func TestNumberOfCards(t *testing.T) {
	d := newDeck()
	if len(d) != ExpectedNumberOfCards {
		t.Errorf("Expected %d cards but got %d", ExpectedNumberOfCards, len(d))
	}
}

func TestFirstAndLastCard(t *testing.T) {
	d := newDeck()

	if d[0] != "One Of Spades" {
		t.Errorf("Expected first card but got %s", d[0])
	}

	if d[len(d)-1] != "King Of Diamonds" {
		t.Errorf("Expected last card but got %s", d[len(d)-1])
	}
}

func TestDeckIo(t *testing.T) {
	cleanupAfterTests(t)
	d := newDeck()
	d.shuffle(33)

	err := d.toFile(IoFileName)
	if err != nil {
		t.Errorf("%v", err)
	}

	ioDeck, err := fromFile(IoFileName)
	if err != nil {
		t.Errorf("%v", err)
	}

	for i := 0; i < 56; i++ {
		randomInt := rand.Intn(len(d) - 1)

		if d[randomInt] != ioDeck[randomInt] {
			t.Errorf("Decks dont match")
		}

	}
	cleanupAfterTests(t)

}
