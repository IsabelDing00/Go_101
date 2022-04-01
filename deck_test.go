/**************************************************************
Write tests for every function for this project, which also can be
tests for every functions

1. Code to make sure that a deck is created with x number of cards
	create a new deck  -> write if statement to see if the deck has the right number of cards
	-> if it doesn't, tell the go test handler that something went wrong


2. clean up:
	delete any file in current working directory with the name "_decktesting" -> create a deck
	-> save to file "_decktesing" -> load from file -> Assert deck length -> delete ant files in current in
	current working directory with the name "_decktesting"
**************************************************************/

package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Error("Expected deck length of 16, but got", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Error("Expected Ace of Spades, but got", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Error("Expected Four of Clubs, but got", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Error("Expected 16 cards in deck, but got", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
