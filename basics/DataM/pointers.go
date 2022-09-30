package main

import "fmt"

func main() {
	fmt.Println("Welcome")
	myNumber := 23
	var ptr = &myNumber // & -> reference to the direct memory location

	fmt.Println("The value of the pointer is ", ptr)
	fmt.Println("The value of the pointer is ", *ptr) //value of that pointer
	*ptr = *ptr * 2
	fmt.Println("Number", myNumber)

	
}
