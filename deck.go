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

type deck []string

//*1
func (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}

}
