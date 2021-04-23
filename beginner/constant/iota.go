package beginner

import (
	"fmt"
)

const (
	pi = 3.1416
	A
	piTambien
	igualAPI
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

func PrintIota() {
	fmt.Println(pi)
	fmt.Println(piTambien)
	fmt.Println(igualAPI)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
