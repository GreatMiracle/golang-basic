package funcDeck

func (d Deck) DealDeck(handSize int) (Deck, Deck) {
	hand := d[:handSize]
	remainingDeck := d[handSize:]

	return hand, remainingDeck
}
