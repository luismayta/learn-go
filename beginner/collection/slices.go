package collection

import (
	"fmt"
)

func PanicSlice() []string {
	eb := make([]string, 3, 5)
	eb[0] = "One"
	eb[1] = "Two"
	eb[2] = "Three"

	return eb
}

func Example1() (int, int) {
	eb := make([]string, 2, 5)

	return len(eb), cap(eb)
}

func Example2() []string {
	personajes := [...]string{
		"Chema",
		"Don-Pimpón",
		"Espinete",
		"Caponata",
	}
	personas := personajes[:2]
	preferidos := personajes[1:3]
	animales := personajes[2:]
	simpaticos := personajes[:]
	fmt.Println("Personajes:", personajes)
	fmt.Println("Personas  :", personas)
	fmt.Println("Preferidos:", preferidos)
	fmt.Println("Animales  :", animales)
	fmt.Println("Simpaticos:", simpaticos)
	return personajes[:]
}

func Example3() {
	personajes := [...]string{"Chema", "Don-Pimpón", "Espinete", "Caponata"}
	personas := personajes[:2]
	preferidos := personajes[1:3]
	animales := personajes[2:]
	simpaticos := personajes[:]
	personajes[0] = "Princesa-Leia"
	personajes[1] = "Han-Solo"
	personajes[2] = "Chewaca"
	personajes[3] = "Jabba"
	fmt.Println("Personajes:", personajes)
	fmt.Println("Personas  :", personas)
	fmt.Println("Preferidos:", preferidos)
	fmt.Println("Animales  :", animales)
	fmt.Println("Simpaticos:", simpaticos)
}

func Example4() {
	villanos := [...]string{"Darth-Vader", "Vizzini", "Célula"}

	personajes := make([]string, 0, 5)
	fmt.Println(personajes, len(personajes), cap(personajes))

	personajes = append(personajes, "Kakaroto")
	fmt.Println(personajes, len(personajes), cap(personajes))

	personajes = append(personajes, "Karel", "Karate-Kid")
	fmt.Println(personajes, len(personajes), cap(personajes))

	personajes = append(personajes, "Karel", "Karate-Kid")
	fmt.Println(personajes, len(personajes), cap(personajes))

	personajes = append(personajes, villanos[:]...)
	fmt.Println(personajes, len(personajes), cap(personajes))

	muchosPersonajes := append(personajes, personajes[1:4]...)
	personajes[1] = "Rumpelstiltskin"
	fmt.Println(personajes, len(personajes), cap(personajes))
	fmt.Println(">", muchosPersonajes, len(muchosPersonajes),
		cap(muchosPersonajes))
}

func Example5() {
	extraterrestres := [...]string{"Yuppi", "Karel", "Alf", "E.T."}
	trozoExtraterrestres := make([]string, 3, 10)
	copiados := copy(trozoExtraterrestres, extraterrestres[:])

	fmt.Println("Extraterrestres        :", extraterrestres)
	fmt.Println("Algunos extraterrestres:", trozoExtraterrestres)
	fmt.Println("Elementos copiados     :", copiados)
}

func Example6() string {
	soloHola := make([]byte, 4)
	copy(soloHola, "Hola a todos")
	return string(soloHola)
}
