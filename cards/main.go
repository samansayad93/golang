package main

import "fmt"

func main() {
	cards := []string{"Five of Diamonds", newcard()}
	cards = append(cards, "Ace of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newcard() string {
	return "Ace of Diamonds"
}
