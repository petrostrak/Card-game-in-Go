package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

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

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//   receiver
//Any variable of type deck now gets access to the print method
//d is the reference to the actual copy of the deck-type variable. In this case cards
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) //0666 permission to all wrx
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // ReadFile function returns a []byte and an error

	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ", ")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		// Swaps i and newPosition with newPosition and i
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
