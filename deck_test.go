package main

import "testing"

var d = newDeck()

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
