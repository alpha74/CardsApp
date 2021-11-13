package main

func main() {
	cards := newDeck()

	hand, remainDeck := deal(cards, 3)

	hand.printDeck()
	remainDeck.printDeck()
}
