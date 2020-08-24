package main

import "fmt"

// Create a new type of 'deck'
// which is slice of strings
type deck []string // this new type called deck kind of extends or borrows all the beheavor of slices

func newDeck() deck {
	cards := deck{}

	//slice of strings
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

//   receiver
//Any variable of type deck now gets access to the print method
//d is the reference to the actual copy of the deck-type variable. In this case cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
