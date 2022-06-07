package main

// func main() {
// 	// only use colon on initialization
// 	card := newCard("Five", "Diamonds")
// 	fmt.Println(card)
// }

func main() {
	cards := newDeck()
	//cards.saveToFile("rhino_cards")
	// cards := newDeckFromFile("rhino_cards")
	cards.shuffle()
	cards.print()
}
