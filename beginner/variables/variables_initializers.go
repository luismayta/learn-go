package main

import (
	"fmt"
)

var (
	name     = "Luis Mayta"
	age      = 30
	location = "Lima - Peru"
	alias    = "Itachi uchiha"
)

func main() {
	fmt.Printf("%s have %d from %s alias %s",
		name,
		age,
		location,
		alias,
	)
}
