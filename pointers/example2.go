package main

import (
	"fmt"
)

type Artist struct {
	Name, Gnre string
	Songs      int
}

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := &Artist{Name: "Luis", Gnre: "Balada", Songs: 42}
	fmt.Printf("%s releases their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
