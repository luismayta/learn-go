package main

import "fmt"

func main() {
	planetas := make(map[string]string)

	planetas["Han Solo"] = "Corellia"
	planetas["Luke"] = "Tatoine"
	fmt.Println("Luke está en:", planetas["Luke"])
	fmt.Println("Han Solo está en:", planetas["Han Solo"])
	planetas["Luke"] = "Alderaan"
	delete(planetas, "Han Solo")
	fmt.Println("Luke está en:", planetas["Luke"])
	fmt.Println("Han Solo está en:", planetas["Han Solo"])
	fmt.Println("Leia Organa está en:", planetas["Leia Organa"])
}
