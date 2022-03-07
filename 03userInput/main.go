package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	message := "User sign up"
	fmt.Println(message)

	usernameRegistrationReader := bufio.NewReader(os.Stdin)
	passwordRegistrationReader := bufio.NewReader(os.Stdin)
	usernameLoginReader := bufio.NewReader(os.Stdin)
	passwordLoginReader := bufio.NewReader(os.Stdin)

	fmt.Println("enter your name: ")
	//comma ok / comma err syntax
	usernameRegistration, _ := usernameRegistrationReader.ReadString('\n')

	fmt.Println("enter your password")
	passwordRegistration, _ := passwordRegistrationReader.ReadString('\n')

	fmt.Println("Thank you for signing up. You can now login")

	fmt.Println("enter your username to login ")
	usernameLogin, _ := usernameLoginReader.ReadString('\n')
	fmt.Println("enter your password ")
	passwordLogin, _ := passwordLoginReader.ReadString('\n')

	if usernameRegistration == usernameLogin && passwordRegistration == passwordLogin {
		fmt.Println("Succefully loged in. Welcome ", usernameRegistration)
	} else {
		fmt.Println("Incorrect password or username")
	}
}
