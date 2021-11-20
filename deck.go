package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// Create a new type which is slice of strings
// 'Deck' extends functionality of []string, analogous to typedef in c++
type Deck []string

// Func to return a list of cards: a deck
func newDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

	// Create combination of cards
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// A receiver (D Deck) on func for 'Deck'
func (D Deck) printDeck() {
	for i, card := range D {
		fmt.Println(i, card)
	}
}

// Returns a hand. Returns two deck
// Not a receiver
func deal(D Deck, handSize int) (Deck, Deck) {
	return D[:handSize], D[handSize:]
}

// Convert a Deck to string. A receiver
func (D Deck) toString() string {
	// Type conversion: Deck to string
	return strings.Join([]string(D), ",")
}

// Func to write Deck to file
func (D Deck) saveToFile(filename string) error {
	// 0666 : Anyone can read/write
	return ioutil.WriteFile(filename, []byte(D.toString()), 0666)
}

// Func to read Deck from file
// Takes file in which Deck is saved
func newDeckFromFile(filename string) Deck {
	// Here, err is value of type error.
	// If nothing went wrongL it will be nil
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error in file reading: ", err)
		os.Exit(1)
	}

	// Convert byte slice to string
	// Split strings on ','
	deckStr := strings.Split(string(byteSlice), ",")

	// We are only able to directly convert to Deck
	// because Deck is really a slice of string
	return Deck(deckStr)
}

// Shuffles a deck
func (D Deck) shuffle() {
	for i := range D {
		// Generates a random number between 0 and len-1
		randomPos := rand.Intn(len(D) - 1)

		// Swap cards
		D[i], D[randomPos] = D[randomPos], D[i]
	}
}
