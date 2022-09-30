package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	/* 1st type of Taking user input */
	var myInp string
	fmt.Scanln(&myInp)
	fmt.Println(myInp)

	/* Second Method */
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter; ")

	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	theinput, _ := reader.ReadString('\n')
	fmt.Println(theinput)
}
