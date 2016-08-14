package main

import (
	"fmt"
)

// Author Data of author songs.
type Author struct {
	Name, Gnre string
	Songs      int
}

func newRelease(a *Author) int {
	a.Songs++
	return a.Songs
}

func main() {
	me := &Author{Name: "Luis", Gnre: "Rock", Songs: 1}
	fmt.Printf("%s releases their %dth song\n", me.Name, newRelease(me))
	fmt.Printf("%s has a total of %d songs", me.Name, me.Songs)
}
