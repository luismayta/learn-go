package main

import "fmt"

var (
	name string
	age int
	location string
)

func main(){
	name, location := "name", "location"
	fmt.Printf(name)
	fmt.Printf(location)
}
