package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string{
	
	return strings.Join([]string(d),",")

}

func (d deck) saveToFile(filename string) error{
	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

func readFromFile(filename string) deck{
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	//string(bs) // The string value of all cards
	s := strings.Split(string(bs),",") //Getting a string slice to make a deck out of

	return deck(s)
	


}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)


	for i := range d{
		newPosition := r.Intn(len(d)-1) // Generate a random number.
		// Getting a perfect shuffle requires the seed to randomly update.


		// To get a new seed value we need some syntax

		d[i], d[newPosition] = d[newPosition], d[i]

	}

}
