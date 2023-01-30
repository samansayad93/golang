package main

func main() {
	cards := newDeck()

	hand, RemainingCards := deal(cards, 5)

	hand.print()
	RemainingCards.print()

	cards.saveToFile("my_cards")

	cards = newDeckFromFile("my_cards")
	cards.print()

	cards.shuffle()
	cards.print()
}
