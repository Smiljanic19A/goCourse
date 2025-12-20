package main

import (
	"testing"
)

const ExpectedNumberOfCards = 56

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
