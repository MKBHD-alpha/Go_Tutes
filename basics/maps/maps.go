package main

import "fmt"

func main() {
	fmt.Println("Maps in GoLang")

	langs := make(map[string]string) //[key]value
	langs["JS"] = "JavaScript"
	langs["RB"] = "Ruby"
	langs["PY"] = "Python"
	langs["GO"] = "GoLang"

	fmt.Println("Languages are ", langs)
	fmt.Println("Js Short for", langs["JS"])

	delete(langs, "RB")
	fmt.Println(langs)

	//loops

	for _, v := range langs {
		fmt.Printf("%v\t", v)

	}

}
