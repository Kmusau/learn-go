package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	message := "user input"
	fmt.Println(message)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter your name: ")

	//comma ok / comma err syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Your name is, ", input)
}
