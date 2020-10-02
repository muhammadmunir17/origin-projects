package main

func main() {
	cards := newDeck()

	// // hand, remainingCards := deal(cards, 10)

	// // hand.print()
	// // remainingCards.print()

	// // fmt.Println([]byte(cards[2]))
	// cards.savetofile("test file")

	// cards := ReadFromFile("test")

	cards.shuffle()
	cards.print() 

}
