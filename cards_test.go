package main

import (
	"testing"
)

func Test_getSuits(t *testing.T) {
	actual := getSuits()
	expected := [4]string{"Diamond", "Club", "Heart", "Spade"}

	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func cardsShouldBeSorted(cards []Card, suitName string, t *testing.T) {
	for i, card := range cards {
		if card.suit.name != suitName {
			t.Errorf("Expected suit to be %s, got %s", suitName, card.suit.name)
		}
		if card.rank != i+1 {
			t.Errorf("Expected rank to be %d, got %d", i+1, card.rank)
		}
	}

}

func TestNewDeck(t *testing.T) {
	deck := NewDeck()

	if len(deck.cards) != 52 {
		t.Errorf("Expected 52 cards, got %d cards", len(deck.cards))
	}

	// First 13 cards should be all the Diamond
	cardsShouldBeSorted(deck.cards[0:13], "Diamond", t)

	// Next 13 cards should be all the Clubs
	cardsShouldBeSorted(deck.cards[13:26], "Club", t)

	// Next 13 cards should be all the Heart
	cardsShouldBeSorted(deck.cards[26:39], "Heart", t)

	// last 13 cards should be all the Spades
	cardsShouldBeSorted(deck.cards[39:52], "Spade", t)
}

func Test_totalRank(t *testing.T) {
	card := Card{suit: Suit{rank: 0}, rank: 2}
	if card.totalRank() != 2 {
		t.Errorf("Expected 2, got %d", card.totalRank())
	}
	card = Card{suit: Suit{rank: 1}, rank: 2}
	if card.totalRank() != 15 {
		t.Errorf("Expected 15, got %d", card.totalRank())
	}
	card = Card{suit: Suit{rank: 2}, rank: 2}
	if card.totalRank() != 28 {
		t.Errorf("Expected 28, got %d", card.totalRank())
	}
	card = Card{suit: Suit{rank: 3}, rank: 2}
	if card.totalRank() != 41 {
		t.Errorf("Expected 41, got %d", card.totalRank())
	}
}

func Test_shuffle(t *testing.T) {
	deck := NewDeck()
	pre_shuffle := make([]Card, 52)
	copy(pre_shuffle, deck.cards)

	for i := 0; i < 10; i++ {
		deck.shuffle()
		diffCount := 0

		for j, card := range deck.cards {
			if pre_shuffle[j] != card {
				diffCount++
			}
		}

		if diffCount == 0 {
			t.Errorf("The last shuffle matched the current shuffle at iteration %d", i)
		}

		t.Logf("shuffle: iteration %d was %0.2f%% shuffled.\n", i, float64(diffCount)/52*100)
	}
}
