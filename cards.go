package main

import (
	"math/big"
	"math/rand"
	"time"
)

// K = 13
// Q = 12
// J = 11
// 2 .. 9
// A = 01

const (
	DeckSize        = 52
	MaxCardsPerSuit = 13
)

type Suit struct {
	name string
	rank int
}

type Card struct {
	suit Suit
	rank int
}

type Deck struct {
	cards []Card
}

func getSuits() [4]string {
	return [4]string{"Diamond", "Club", "Heart", "Spade"}
}

func NewDeck() *Deck {
	deck := new(Deck)

	deck.cards = make([]Card, DeckSize)

	suits := getSuits()

	for i := range deck.cards {
		suitRank := i / MaxCardsPerSuit

		suit := Suit{
			name: suits[suitRank],
			rank: suitRank,
		}

		cardRank := (i % MaxCardsPerSuit) + 1

		deck.cards[i] = Card{suit: suit, rank: cardRank}
	}

	return deck
}

func (c *Card) totalRank() int {
	return c.rank + (c.suit.rank * MaxCardsPerSuit)
}

func (deck *Deck) shuffle() {
	src := rand.NewSource(time.Now().Unix())
	rnd := rand.New(src)
	myInt := big.NewInt(0)
	upper := big.NewInt(52)

	cards := deck.cards

	for i := range cards {
		j := myInt.Rand(rnd, upper).Int64()
		cards[i], cards[j] = cards[j], cards[i] // swap
	}
}
