package main

import "fmt"

func main() {
	var ptr *int

	var myVariable = 22

	var myptr = &myVariable

	fmt.Println("value of actual pointer is ", myptr)
	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", myVariable)
	fmt.Println("value of actual pointer is ", *myptr)

	*myptr = *myptr + 2

	fmt.Print(*myptr)
}
