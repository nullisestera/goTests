package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
// Custom Type deck
type deck []string

// newDeck
// Function without arguments
func newDeck() deck {
	cards := deck{}

	// Assign new values
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	// Loop over Suits
	// We use underscore to avoid unused index variables warning
	for _, suit := range cardSuits {
		// Loop over Values
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print
// Receiver function
// "cards" may be whatever variable that represent the param sent, i.e.: "c"
// By convention we usually called the first letter of the custom type
// i.e.: d is for deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal
// Function with arguments
// Always you must pass argument (space) type.
// In this case, d is for deck custom type, and handSize is integer
// The second () is that I expected return
func deal(d deck, handSize int) (deck, deck) {
	// [:indexNotInclude] 0 to index not included
	// [indexInclude:] index to final
	return d[:handSize], d[handSize:]
}
