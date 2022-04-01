/**************************************************************
Build a Card project
Features:
	newDeck: create and return a list of playing cards. Essentially an array
	print: Log out the contents of deck of cards
	shuffle: Shuffles all the cards in a deck
	deal: Create a hand of cards
	saveToFile: Save a list of cards to a file on the local machine
	newDeckFromFile: Load a list of cards from the local machine

**************************************************************/

package main

func main() {
	// TODO 4: use the shuffle()
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// TODO 3: readFile()
	// cards := newDeckFromFile("my_card")
	// cards.print()

	// TODO 2: saveToFile()
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())

	// TODO 1: deal()
	// cards := newDeck()
	// handDeck, remainingCards := deal(cards, 5)
	// handDeck.print()
	// remainingCards.print()
}
