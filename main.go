package main

import (
	"fmt"
)

func main() {

	deck := NewDeck()
	deck.shuffle()

	fmt.Println("=== Shuffled ===")

	for i := 0; i < len(deck.cards); i++ {
		fmt.Println(deck.cards[i])
	}

	fmt.Println("=== Sorted ===")

	deck.sort()

	for i := 0; i < len(deck.cards); i++ {
		fmt.Println(deck.cards[i])
	}
}
