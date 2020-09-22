package main

import "fmt"

func main() {
	// var card string = "Ace of Spades" // Static way
	card := "Ace of Spades"   // Dynamic way
	card = "Five of Diamonds" // Redeclare variables, doesn't need colon (:)
	fmt.Println(card)
}
