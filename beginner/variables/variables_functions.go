package main

import "fmt"

anonima := func() {
    fmt.Printf("Itachi Uchiha")
}

func main() {
    anonima()
    //fmt.Printf("I Am %s", anonima())
}
