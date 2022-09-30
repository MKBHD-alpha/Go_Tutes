package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	deepraj := User{"Deepraj", "deepraj@go.dev", true, 19}
	fmt.Println(deepraj)
	fmt.Printf("Deepraj Details %+v\n", deepraj)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
