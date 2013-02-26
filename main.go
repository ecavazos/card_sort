package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

const (
	DeckSize        = 52
	MaxCardsPerSuit = 13
)

type Suit struct {
	name string
	rank int
}

// K = 13
// Q = 12
// J = 11
// 2 .. 9
// A = 01
type Card struct {
	suit Suit
	rank int
}

func getSuits() [4]string {
	return [4]string{"Heart", "Club", "Diamond", "Spade"}
}

func (c *Card) totalRank() int {
	return c.rank + (c.suit.rank * MaxCardsPerSuit)
}

func createDeck() []Card {
	suits := getSuits()

	deck := make([]Card, DeckSize)

	for i := 0; i < DeckSize; i++ {
		suitRank := i / MaxCardsPerSuit

		suit := Suit{
			name: suits[suitRank],
			rank: suitRank,
		}

		cardRank := (i % MaxCardsPerSuit) + 1

		deck[i] = Card{suit: suit, rank: cardRank}
	}

	return deck
}

func shuffle(deck []Card) []Card {
	src := rand.NewSource(0)
	rnd := rand.New(src)
	myInt := big.NewInt(0)
	upper := big.NewInt(52)

	for i := 0; i < DeckSize; i++ {
		j := myInt.Rand(rnd, upper).Int64()
		tmp := deck[i]
		deck[i] = deck[j]
		deck[j] = tmp
	}

	return deck
}

func sort(deck []Card) []Card {
	for i := 0; i < DeckSize; i++ {
		tmp := deck[i]
		j := i

		for j > 0 && deck[j-1].totalRank() > tmp.totalRank() {
			deck[j] = deck[j-1]
			j--
		}
		deck[j] = tmp
	}

	return deck
}

func main() {

	deck := createDeck()
	deck = shuffle(deck)

	fmt.Println("=== Shuffled ===")

	for i := 0; i < len(deck); i++ {
		fmt.Println(deck[i])
	}

	fmt.Println("=== Sorted ===")

	deck = sort(deck)

	for i := 0; i < len(deck); i++ {
		fmt.Println(deck[i])
	}
}
