package main

import (
	"os"
	"testing"
)

// all func in test are gonna be calle dby a go-test-runner
// t of type *testing .T

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected length of 16, but got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spaces, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Error("Expected length of 16, but got", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
