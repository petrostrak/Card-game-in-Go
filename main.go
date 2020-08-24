package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "Ace of Hearts") // append doesn't modify the current slice, instead it returns a new one

	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
