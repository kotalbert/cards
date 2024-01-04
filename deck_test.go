package main

import (
	"os"
	"testing"
)

var d = newDeck()

const TestFilename = "deck-testing.txt"

// Make sure deck is created with correct number of cards (52)
func TestNewDeckSize(t *testing.T) {
	e := 52
	a := len(d)

	if e != a {
		t.Errorf("Expected deck length of %d, but got %d", e, a)
	}
}

// Make sure first card in new deck is "Ace of Spades"
func TestNewDeckFirstCard(t *testing.T) {
	e := "Ace of Spades"
	a := d[0]

	if e != a {
		t.Errorf("Expected %s, but got %s", e, a)
	}
}

// Make sure first card in new deck is "King of Clubs"
func TestNewDeckLastCard(t *testing.T) {
	e := "King of Clubs"
	a := d[len(d)-1]

	if e != a {
		t.Errorf("Expected %s, but got %s", e, a)
	}
}

func clearWorkingDirectory() {
	_ = os.Remove(TestFilename)
}

func TestDeckToFile(t *testing.T) {
	// fixme: create setup() and teardown()
	clearWorkingDirectory()

	err := d.saveToFile(TestFilename)
	if err != nil {
		t.Errorf("Save to file Error: %v", err)
	}

	clearWorkingDirectory()

}

func TestDeckFromFile(t *testing.T) {
	clearWorkingDirectory()

	err := d.saveToFile(TestFilename)
	if err != nil {
		t.Errorf("Save to file Error: %v", err)
	}

	a, err2 := newDeckFromFile(TestFilename)

	if err2 != nil {
		t.Errorf("Error: %v", err2)
	}

	if a.toString() != d.toString() {
		t.Errorf("Deck loaded from file does not match expected deck")
	}

	clearWorkingDirectory()
}
