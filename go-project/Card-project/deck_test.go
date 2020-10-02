package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck of length 16 but got %d", len(d))
	}

	if d[0] != "ace of spades" {
		t.Errorf("expected first card of ace of spades but got %v", d[0])

	}

	if d[len(d)-1] != "four of clubs" {
		t.Errorf("expected last card of four of clubs but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.savetofile("_decktesting")

	loadedDeck := ReadFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("expected 16 cards in deck,got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
