package main

// import "fmt"

func main() {
	// cards := newDeck()
	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()
	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards.csv")
	// cards := newDeckFromFile("my_card.csv")
	// cards.print()
	// fmt.Println(cards.toString())

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Spades"
}
