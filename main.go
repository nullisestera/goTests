package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // Static way
	// card := "Ace of Spades"   // Dynamic way
	// card = "Five of Diamonds" // Redeclare variables, doesn't need colon (:)

	card := newCard()

	fmt.Println(card)

	// For Loops & Arrays (Slice)
	colors := []string{"Blue", "Red", "Green", "Yellow"}
	for index, color := range colors {
		fmt.Println(index, color)
	}

}

func newCard() string {
	return "Ace of spades"
}
