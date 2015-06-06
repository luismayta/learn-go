package main

import "fmt"

var (
    name string = "Luis Mayta"
    age int = 30
    location string = "Lima - Peru"
)

func main() {
    fmt.Printf("%s have %d from %s", name, age, location)
}
