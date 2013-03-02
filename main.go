package main

import (
	"fmt"
)

func printCards(d *Deck) {
	for _, card := range d.cards {
		fmt.Println(card)
	}
}

func main() {
	deck := NewDeck()
	deck.shuffle()

	fmt.Println("==== Shuffled ===")
	printCards(deck)

	fmt.Println("====  Sorted  ===")
	deck.sort()
	printCards(deck)
}
