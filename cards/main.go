package main

func main() {
	cards := deck{"Five of Diamonds", newcard()}
	cards = append(cards, "Ace of Spades")

	cards.print()
}

func newcard() string {
	return "Ace of Diamonds"
}
