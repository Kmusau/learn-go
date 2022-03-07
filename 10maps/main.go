package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RR"] = "Ruby on Rails"
	languages["SB"] = "Spring Boot"

	fmt.Println("List of all the languages ", languages)
	fmt.Println("SB shorts for ", languages["SB"])

	delete(languages, "RR")

	fmt.Println(languages)

	//loops in maps

	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n ", key, value)
	}
}
