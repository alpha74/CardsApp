package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainDeck := deal(cards, 5)

	remainDeck.printDeck()

	// Save hand to file
	filename := "currentHand"
	fmt.Println("Saving current Hand to file: ", filename)
	hand.saveToFile(filename)

	// Read hand from file
	handFromFile := newDeckFromFile(filename)
	fmt.Println("\n Hand read from file: ")

	// Shuffle the deck
	handFromFile.shuffle()
	handFromFile.printDeck()
}
