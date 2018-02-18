package main

import "github.com/dnvriend/deck"

func main() {
	cards := deck.NewDeck()
	cards.Shuffle()
	cards.Print()
}