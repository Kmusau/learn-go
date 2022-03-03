package main

import "fmt"

const LoginToken = "guvjhbn" //public because of first letter caps

func main() {
	var username string = "Kelvin"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.25678
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// implicit type
	var users = 20000
	fmt.Println(users)

	// no var keyword
	numberOfUsers := 30000
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
}
