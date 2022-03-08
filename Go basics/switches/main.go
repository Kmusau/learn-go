package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Exploring the switch statement")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Dice number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1 and you can open")
	case 2:
		fmt.Println("You got 2 and move 2 steps")
	case 3:
		fmt.Println("You got 3 and move 3 steps")
	case 4:
		fmt.Println("You got 4 and move 4 steps")
	case 5:
		fmt.Println("You got 5 and move 5 steps")
	case 6:
		fmt.Println("You got 6, move 2 steps and roll the dice again")
	}
}
