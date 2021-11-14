package main

import (
	"fmt"
	"strings"
)

// Create a new type which is slice of strings
// 'Deck' extends functionality of []string, analogous to typedef in c++
type Deck []string

// Func to return a list of cards: a deck
func newDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Joker", "Queen", "King"}

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
