package main

import "fmt"

func main() {
	fmt.Println("Learning struts")

	kevo := User{"Kelvin", "kelvin@dawascope.com", true, 23}

	fmt.Println(kevo)
	fmt.Printf("All details are %+v \n", kevo)
	fmt.Printf("Kevo's name is %v and email is %v \n", kevo.Name, kevo.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
