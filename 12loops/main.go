package main

import "fmt"

func main() {
	fmt.Println("Loops in Go")

	var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days)

	fmt.Println("normal for loop")

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println("for range loop")

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println("for range loop with key-value pair")

	for index, day := range days {
		fmt.Printf("Index is %v and day is %v \n", index, day)
	}

	// while loop type in golang
	value := 1
	for value <= 5 {
		fmt.Println("Value is ", value)
		value++
	}
}
