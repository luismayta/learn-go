package main

import "fmt"

var (
	name, location string
	age int
)

func main(){
	name, location := "itachi", "konoha"
	fmt.Printf("%s of %s", name, location)
}
