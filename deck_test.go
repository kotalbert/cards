package main

import "testing"

// Make sure deck is created with correct number of cards (52)
func TestNewDeckSize(t *testing.T) {
	d := newDeck()

	e := 52
	a := len(d)

	if a != e {
		t.Errorf("Expected deck length of %d, but got %d", e, a)
	}
}

// Make sure first card in new deck is "Ace of Spades"
func TestNewDeckFirstCard(t *testing.T) {
	t.Fatalf("Implement")

}

// Make sure first card in new deck is "King of Clubs"
func TestNewDeckLastCard(t *testing.T) {
	t.Fatalf("Implement")
}
