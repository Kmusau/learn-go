package main

import "fmt"

func main() {
	greeterTwo()
	fmt.Println("Studying functions in Golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is ", result)

	proResult, additionalMessage := proAdder(2, 5, 3, 8, 5, 9)
	fmt.Println("Pro result is: ", proResult, additionalMessage)
}

func greeter() {
	fmt.Println("Namstey from this side")
}

func greeterTwo() {
	fmt.Println("Hello world")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "This is an additional message returned"
}
