package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	
	fmt.Println("Switch Case in GO !!!")
	fmt.Println("Dice Roller")

	//@ Random Number Generator

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Value is", diceNumber)

	//! Switch Case Statement
	switch diceNumber {
	case 1:
	
		fmt.Println("Value is 1")
	case 2:
		fmt.Println("Value is 2")
	case 3:
		fmt.Println("Value is 3")
		fallthrough //? Used to redirect to the next case
	case 4:
		fmt.Println("Value is 4")
	case 5:
		fmt.Println("Value is 5")
	case 6:
		fmt.Println("Value is 6")

	}
}
