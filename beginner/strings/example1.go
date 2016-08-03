package main

import (
	"fmt"
)

func example1() {
	var s1 string = "Hola, me llamo Luis Mayta"
	fmt.Println(s1)
	fmt.Println("¿Cómo has dicho que te llamas?")
	fmt.Println(s1[15:])
}

func example2() {
	var s1 string = "Hola, me llamo Íñigo Montoya"
	s1[3] = 'X'
	fmt.Println(s1)
}

func main() {
	example1()
	// example2()
}
