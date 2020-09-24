package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

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
	for _, card := range d {
		fmt.Println(card)
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

// Function receiver convert to String
func (d deck) toString() string {
	// strings.Join (slice, "separator")
	return strings.Join([]string(d), ",")
}

// Function receiver save to a File
func (d deck) saveToFile(filename string) error {
	// ioutil.WriteFile syntax
	// filename: the file name in string
	// byte, a byte slice (string converted)
	// permisions. i.e.: 0666 (everybody can write and read)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Function with params new deck from file (read file)
func newDeckFromFile(filename string) deck {
	// bs: byte slice received by ioutil.ReadFile function
	// err: error object
	bs, err := ioutil.ReadFile(filename)
	// nil === expected result
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		log.Fatal(err)
		os.Exit(1)
	}

	// Convert in a string and separate
	s := strings.Split(string(bs), ",")
	// Return deck with string separated
	return deck(s)
}

// Function receiver shuffle
func (d deck) shuffle() {
	// Loop over deck
	for i := range d {
		// create a source to randomize (Int64), in this case with time (time is the seed to NewSource)
		source := rand.NewSource(time.Now().UnixNano())
		// create a new random with created source
		r := rand.New(source)
		// new position === random integer length of deck - 1
		newPos := r.Intn(len(d) - 1)
		// Sort positions (Swap)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
