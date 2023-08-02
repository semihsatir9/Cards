package main


//import "fmt"

// Variables can be initialized outside of a
// function but they CANNOT be given value outside

func main() {
	// Declaring Variables


	// var card string = "Ace of Spades"
	
	// card := newCard()

	// Using := makes Go automatically assign the 
	// type of the variable for the given data value

	// Only use it when you are defining the data
	// for the first time!
	// card = "Five of Diamonds"
	// For changing the data after declaration


	// Slices

	// Like arraylists on Java. Many functions to be used.
	// Must declare a type for the slice. The slice
	// will only take data of that type.

	//Append is .add of arraylists


	// To understand the deck type, go to the
	// deck.go file.

	//var cards deck = newDeck()


	// Recieving 2 different returns from a single
	// function.

	// hand, remainingCards := deal(cards,5)

	// hand.print()
	// remainingCards.print()

	//cards.saveToFile("This")

	newDeck := readFromFile("This")

	newDeck.print()

	newDeck.shuffle()

	newDeck.print()



	// For loop
	
	// Iterating over a slice, use range to define
	// the size of the slice.

	// Use := because after every iteration, the old
	// data is thrown away. You have to assign new values
	// every time the loop iterates.

	// for i, card := range d{
	// 	 fmt.Println(i,card)
	// }
	
	

}
