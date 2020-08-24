package main

import "fmt"

// Create a new type of 'deck'
// which is slice of strings
type deck []string // this new type called deck kind of extends or borrows all the beheavor of slices

//   receiver
//Any variable of type deck now gets access to the print method
//d is the reference to the actual copy of the deck-type variable. In this case cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
