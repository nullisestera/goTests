package main

func main() {
	// var card string = "Ace of Spades" // Static way
	// card := "Ace of Spades"   // Dynamic way
	// card = "Five of Diamonds" // Redeclare variables, doesn't need colon (:)
	//card := newCard()
	//fmt.Println(card)

	// For Loops & Slices (Arrays)
	/* colors := []string{"Blue", "Red", "Green", "Yellow"}
	for index, color := range colors {
		fmt.Println(index, color)
	} */
	/*// Invoke new Deck function
	cards := newDeck()
	// Invoke print receiver function
	// cards.print()

	// Invoke deal function
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()*/

	// Write and read in File System
	//cards := newDeck()
	//cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	// cards.print()

	// Shuffle
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
