package main

import (
	"encoding/json"
	"fmt"
	// "io/ioutil"
)

// Person Struct for example json
type Person struct {
	Name, City string
}

var person Person

var people []Person

func main() {
	jsonStr := "{ \"Name\": \"Marcus\", \"City\": \"San Jose\"}"

	if err := json.Unmarshal([]byte(jsonStr), &person); err != nil {
		fmt.Println("Error Parsing Json", err)
	}
	fmt.Printf("Name: %v, City: %v\n", person.Name, person.City)
}
