// Examples with iota
package main

import (
	"fmt"
)

const (
	pi = 3.1416
	A
	piTambien
	igualApi
)

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(pi)
	fmt.Println(piTambien)
	fmt.Println(igualApi)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
