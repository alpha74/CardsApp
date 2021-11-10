package main

import "fmt"

// Create a new type which is slice of strings
// 'Deck' extends functionality of []string, analogous to typedef in c++
type Deck []string

// A receiver (D Deck) on func for 'Deck'
func (D Deck) printDeck() {
	for i, card := range D {
		fmt.Println(i, card)
	}
}
