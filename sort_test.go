package main

import (
	"testing"
)

func sortedCards() (c []Card) {
	c = make([]Card, 52)
	for i := range c[0:13] {
		c[i] = Card{Suit{"Diamond", 0}, i + 1}
	}
	for i := range c[13:26] {
		c[i+13] = Card{Suit{"Club", 1}, i + 1}
	}
	for i := range c[26:39] {
		c[i+26] = Card{Suit{"Heart", 2}, i + 1}
	}
	for i := range c[39:52] {
		c[i+39] = Card{Suit{"Spade", 3}, i + 1}
	}
	return c
}

func TestSort(t *testing.T) {
	sortedCards := sortedCards()

	deck := NewDeck()
	deck.shuffle()
	deck.sort()

	for i, card := range deck.cards {
		if sortedCards[i] != card {
			t.Fatal("The deck was not properly sorted!")
		}
	}
}

func BenchmarkSort(b *testing.B) {
	deck := NewDeck()
	deck.shuffle()

	for i := 0; i < b.N; i++ {
		deck.sort()	
	}
}
