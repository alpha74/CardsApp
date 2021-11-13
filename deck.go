package main

import "fmt"

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
