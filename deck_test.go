package main

import (
	"os"
	"testing"
)

// Testing new Deck
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Write some assertions
	// Assert that deck has 52 cards
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	// Assert that the first card is Ace of Spades
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// Assert that the last card is King of Clubs
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
