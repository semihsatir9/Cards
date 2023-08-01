package main

import "fmt"

// Variables can be initialized outside of a 
// function but they CANNOT be given value outside

func main() {
	//var card string = "Ace of Spades"
	card := newCard()

	// Using := makes Go automatically assign the 
	// type of the variable for the given data value

	// Only use it when you are defining the data
	// for the first time!

	card = "Five of Diamonds"
	fmt.Println(card)
}

func  newCard() string{
	return "Five of Diamonds"
}
