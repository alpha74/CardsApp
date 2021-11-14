package main

import "fmt"

func main() {
	cards := newDeck()

	hand, remainDeck := deal(cards, 3)

	fmt.Println("Saving current Hand to file")
	hand.saveToFile("currentHand")
	remainDeck.printDeck()
}
