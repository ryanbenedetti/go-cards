package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	//fmt.Println(len(d))
	if len(d) != 52 {
		t.Errorf("expected deck length of 52 but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("expected first card to be Ace of Clubs but got %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("expected last card to be King of Spades but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Clean up old tes files
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("expected 52, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
