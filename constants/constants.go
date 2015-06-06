package main

import "fmt"

const Pi = 3.1416

const (
    StatusOk = 200
    StatusNotFound = 404
)

func main() {
    fmt.Println("Value Pi %f", Pi)
    fmt.Println("Value StatusOk %d", StatusOk)
    fmt.Println("Value StatusNotFound %d", StatusNotFound)
}
