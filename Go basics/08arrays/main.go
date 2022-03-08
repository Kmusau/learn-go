package main

import "fmt"

func main() {
	fmt.Println("Learning arrays")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[3] = "Pawpaw"

	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Print(vegList)
}
