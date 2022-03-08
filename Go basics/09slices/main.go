package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices")

	var fruitlist = []string{"apple", "banana", "mangoes"}
	fmt.Println(fruitlist)
	fruitlist = append(fruitlist, "tomato", "peach")
	fmt.Println(fruitlist)
	// fruitlist = append(fruitlist[1:3])
	// fmt.Println(fruitlist)
	fruitlist = append(fruitlist[:3])
	fmt.Println(fruitlist)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 965
	highScores[2] = 354
	highScores[3] = 582

	fmt.Println(highScores)

	highScores = append(highScores, 584, 615, 461)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove a value from a slice based on index
	var courses = []string{"React", "Angular", "Vue", "Quasar", "goFibre"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
