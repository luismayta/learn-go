package main

import "fmt"

const Pi = 3.1416

const (
    StatusOk = 200
    StatusNotFound = 404
)

func main() {
    fmt.Printf("Value Pi %f", Pi)
    fmt.Printf("Value StatusOk %d", StatusOk)
    fmt.Printf("Value StatusNotFound %d", StatusNotFound)
}
