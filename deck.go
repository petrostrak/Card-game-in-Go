package main

import "fmt"

// Create a new type of 'deck'
// which is slice of strings
type deck []string // this new type called deck kind of extends or borrows all the beheavor of slices

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
