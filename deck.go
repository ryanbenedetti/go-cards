package main

import "fmt"

// Create a new type of deck
// which is a slice of strings
type deck []string

func newDeck() deck {
	// create and return list
	// of playing cards
	cards := deck{}

	// card slices
	cardSuits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	cardFaces := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// two for loops
	for _, suit := range cardSuits {
		for _, face := range cardFaces {
			//fmt.Println(_, face)
			cards = append(cards, face+" of "+suit)
		}
		//fmt.Println(_, suit)
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
