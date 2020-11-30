package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 cards in new deck, but got %d", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades but got %s", d[0])
	}
	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card is King of Clubs but got %s", d[len(d)-1])

	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %d", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
