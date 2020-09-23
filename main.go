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

	// Pass props to deck type (Custom Type)
	cards := deck{"Ace of Spades", "Five of Diamonds", "Three of Hearts"}
	// Invoke print function
	cards.print()

}

func newCard() string {
	return "Ace of spades"
}
