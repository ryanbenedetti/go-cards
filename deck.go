package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	// set up our seed with an int64 using time.Now
	seed := time.Now().UnixNano()
	// set up a rondom source
	source := rand.NewSource(seed)
	//create a new random
	r := rand.New(source)

	//shuffle algorithm
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
