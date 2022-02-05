package main

// func main() {
// 	// only use colon on initialization
// 	card := newCard("Five", "Diamonds")
// 	fmt.Println(card)
// }

func main() {
	cards := newDeck()
	cards.print()
}
