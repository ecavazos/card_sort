package main

// Insertion sort
func (deck *Deck) sort() []Card {
	cards := deck.cards

	for i := 0; i < DeckSize; i++ {
		tmp := cards[i]
		j := i

		for j > 0 && cards[j-1].totalRank() > tmp.totalRank() {
			cards[j] = cards[j-1]
			j--
		}

		cards[j] = tmp
	}

	return cards
}
