package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
}
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingCards := deal(d, 5)

	if len(hand) != 5 {
		t.Errorf("Expected hand size of 5, but got %v", len(hand))
	}

	if len(remainingCards) != 47 {
		t.Errorf("Expected remaining cards size of 47, but got %v", len(remainingCards))
	}
}

func TestToString(t *testing.T) {
	d := newDeck()
	s := d.toString()

	if s == "" {
		t.Errorf("Expected string not to be empty")
	}
}
