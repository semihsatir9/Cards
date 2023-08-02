package main

import (
	"os"
	"testing"
)

//t.Errorf("Expected deck length of 16, but got %v", len(d))

// For all functions, create a test function in this
// file. What to test is up to you and the function.


func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of 52 but got %v", len(d))


	}

	if d[0] != "Ace of Spades"{
		t.Errorf("The ace of spades is not the starting value. The starting value is %v",d[0])
	}

	if d[len(d) - 1] != "Jack of Diamonds"{
		t.Errorf("The jack of diamonds is not the final value. The final value is %v",d[len(d) - 1])
	}

}

func TestSaveToFile(t *testing.T){

	// Take care of cleanup before doing any writing
	// to hardware.

	os.Remove("_decktesting")
	deck := newDeck()

	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if(len(loadedDeck) != 52){
		t.Errorf(" Expected 52. Got :  %v",len(loadedDeck))
	}


	

}