package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in GO")

	highScores := make([]int, 4)

	highScores[0] = 435
	highScores[1] = 800
	highScores[2] = 500
	highScores[3] = 560
	//highScores[4] = 560  index out of range error

	highScores = append(highScores, 555, 66, 7, 8, 666)
	//fmt.Println("Scores are ", highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	//! Removing Element

	courses := []string{"Swift", "Python", "Java", "JavaScript", "Flutter"}
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
