package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // Static way
	// card := "Ace of Spades"   // Dynamic way
	// card = "Five of Diamonds" // Redeclare variables, doesn't need colon (:)

	card := newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Ace of spades"
}
