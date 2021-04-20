package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Person Struct for example json.
type Person struct {
	Name, City string
}

func main() {
	var (
		person Person
		people []Person
	)

	jsonStr := "{ \"Name\": \"Marcus\", \"City\": \"San Jose\"}"

	if err := json.Unmarshal([]byte(jsonStr), &person); err != nil {
		fmt.Println("Error Parsing Json", err)
	}

	fmt.Printf("Name: %v, City: %v\n", person.Name, person.City)
	file, err := ioutil.ReadFile("./beginner/json/extras/names.json")
	if err != nil {
		fmt.Println("Error parsing json", err)
	}

	if err := json.Unmarshal(file, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)

	response, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error encoding json")
	}

	fmt.Println(string(response))
}
