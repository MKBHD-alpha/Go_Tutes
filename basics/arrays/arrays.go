package main

import "fmt"

func main() {
	fmt.Println("Welcome")

	var myList [4]string
	myList[0] = "Apple"
	myList[1] = "Banana"
	myList[2] = "Mango"
	myList[3] = "Guava"
	fmt.Println("Fruit List is ", myList)
	fmt.Println("Fruit List is ", len(myList))

	secondArray := [4]int{0, 1, 2, 3}
	fmt.Println("Second Array is ", secondArray)
	fmt.Println("Hekki ")
	
}
