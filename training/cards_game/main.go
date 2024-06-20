package main

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of hearts"
	// card = "Ace of clubs"

	// cards := newDeck()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards.print()

	// hand, remainingCards := deal(cards, 5)
	// fmt.Println()
	// fmt.Println("printing hand")
	// hand.print()
	// fmt.Println()
	// fmt.Println("printing remaining cards")
	// remainingCards.print()
	// fmt.Println()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
