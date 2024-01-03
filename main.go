package main

func main() {

	cards := newDeck()

	_ = cards.saveToFile("cards.txt")

}
