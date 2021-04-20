package main

import "fmt"

func main() {
	planets := make(map[string]string)
	planets["Han Solo"] = "Corellia"
	planets["Luke"] = "Tatoine"
	fmt.Printf("Luke this in: %s", planets["Luke"])         // nolint
	fmt.Printf("Han Solo this in: %s", planets["Han Solo"]) // nolint
	planets["Luke"] = "Alderaan"
	delete(planets, "Han Solo")
	fmt.Printf("Luke now this in: %s", planets["Luke"])           // nolint
	fmt.Printf("Han Solo now this in: %s", planets["Han Solo"])   // nolint
	fmt.Printf("Leia Organa this in: %s", planets["Leia Organa"]) // nolint
}
