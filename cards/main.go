package main

func main() {

	cards, remainingDeck := deal(newDeck(), 5)

	cards.print()
	remainingDeck.print()

}
