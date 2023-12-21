package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Person Struct for example json.
type Person struct {
	Name string `json:"Name"`
	City string `json:"City"`
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
	file, erra := os.Open("./beginner/json/extras/names.json")
	if erra != nil {
		fmt.Println("Error parsing json", erra)
	}

	scanner := bufio.NewScanner(file)

	if err := json.Unmarshal(scanner.Bytes(), &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	defer func() {
		if errc := file.Close(); errc != nil {
			fmt.Println("Wrong to close file:", errc)
		}
	}()

	fmt.Println(people)

	response, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error encoding json")
	}

	fmt.Println(string(response))
}
