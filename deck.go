package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// define new type `deck`
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearths", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack",
		"Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return cards

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Take handSize from deck and return two new decks.
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Convert a deck to comma separated list of cards
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	bs, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	s := strings.Split(string(bs), ",")
	d := deck{}

	for _, card := range s {
		d = append(d, card)
	}

	return d, err

}

// Shuffle deck in place by randomly swapping cards.
func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		np := r.Intn(len(d) - 1)
		d[i], d[np] = d[np], d[i]
	}
	return d
}
