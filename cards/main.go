package main

func main() {
	cards := newDeck()

	hand, RemainingCards := deal(cards, 5)

	hand.print()
	RemainingCards.print()
}
