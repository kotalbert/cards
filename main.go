package main

import "fmt"

func main() {

	cards := newDeck()

	_ = cards.saveToFile("cards.txt")

	newCards, _ := newDeckFromFile("cards.txt")

	fmt.Println(cards.toString() == newCards.toString())

}
