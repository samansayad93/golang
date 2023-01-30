package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected firdt card of  Ace of Spades, But got %v", d[0])
	}

	if d[51] != "King of Clubs" {
		t.Errorf("Expected firdt card of  King of Clubs, But got %v", d[0])
	}
}

func TestSavetodeckandnewdeckfromfile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadeddeck := newDeckFromFile("_decktesting")

	if len(deck) != len(loadeddeck) {
		t.Errorf("Lenght of deck and loaded deck is not equal")
	}

	os.remove("_decktesting")
}
