package main

import "fmt"

func main() {
	fmt.Println("Methods")

	kevo := User{"Kelvin", "kelvin@dawascope.com", true, 23}

	// fmt.Println(kevo)
	// fmt.Printf("All details are %+v \n", kevo)
	// fmt.Printf("Kevo's name is %v and email is %v \n", kevo.Name, kevo.Email)

	kevo.GetStatus()
	kevo.Newmail()

	fmt.Printf("All details are %+v \n", kevo)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("The status is ", u.Status)
}

func (u User) Newmail() {
	u.Email = "test@go.dev"
	fmt.Println("The new email is ", u.Email)
}
