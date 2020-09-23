package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// Receiver function
// "cards" may be whatever variable that represent the param sent, i.e.: "c"
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
