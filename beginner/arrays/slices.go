package arrays

import "fmt"

func PanicSlice() {
	var eb = make([]string, 2, 5)
	eb[0] = "Epi"
	eb[1] = "Blas"
	eb[2] = "Coco"
	fmt.Println("Saludemos a:", eb)
}

func Example1() {
	var eb = make([]string, 2, 5)
	fmt.Println("Tamaño slice:   ", len(eb))
	fmt.Println("Capacidad slice:", cap(eb))

	var array [5]string
	fmt.Println("Tamaño array:   ", len(array))
	fmt.Println("Capacidad array:", cap(array))
}

func Example2() {
	var personajes = [...]string{"Chema", "Don-Pimpón",
		"Espinete", "Caponata"}
	personas := personajes[:2]
	preferidos := personajes[1:3]
	animales := personajes[2:]
	simpaticos := personajes[:]
	fmt.Println("Personajes:", personajes)
	fmt.Println("Personas  :", personas)
	fmt.Println("Preferidos:", preferidos)
	fmt.Println("Animales  :", animales)
	fmt.Println("Simpaticos:", simpaticos)
}

func Example3() {
	var personajes = [...]string{"Chema", "Don-Pimpón", "Espinete", "Caponata"}
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
	var villanos = [...]string{"Darth-Vader", "Vizzini", "Célula"}

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

	// personajes = append(personajes, "blancanieves", "Sabio",
	// "Gruñon", "Mocoso", "Tímido", "Mudito", "Dormilón",
	// "Feliz")
	// fmt.Println(personajes, len(personajes), cap(personajes))
}

func Example5() {
	var extraterrestres = [...]string{"Yuppi", "Karel", "Alf", "E.T."}
	trozoExtraterrestres := make([]string, 3, 10)
	copiados := copy(trozoExtraterrestres, extraterrestres[:])

	fmt.Println("Extraterrestres        :", extraterrestres)
	fmt.Println("Algunos extraterrestres:", trozoExtraterrestres)
	fmt.Println("Elementos copiados     :", copiados)
}

func Example6() {
	var soloHola = make([]byte, 4)
	copy(soloHola, "Hola a todos")
	fmt.Println(soloHola)
}
