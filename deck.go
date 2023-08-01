package main

import "fmt"

// Custom Types
// In Go, there is no object oriented classes.
// In this case, we extend a custom data type from
// an already existing type.

// A lot similar to classes on OOP languages.

// Functions on custom types takes a receiver.
// It's like the insides of the paranthesis of 
// Java functions. The syntax is found below. *1
// Only use this abbreviation when you are initializing
// a new method to a type.

type deck []string

func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","King","Queen","Jack"}


	// Underscore lets Go skip a variable that we know
	// we wont use

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, (value + " of " + suit))


		}
	}

	return cards

}

//*1
func (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}

}

// Indexing and Dividing Slices
// Zero indexed like any other language.
// fruits[0:2] means give all objects from 0 included 
// to 2 (2 not included)

//fruits[2:] means give me all from index 2 to the end

// You can assign 2 different return values in Go
func deal(d deck, handsize int) (deck,deck){
	return d[:handsize],d[handsize:]
	

}
