package beginner

func newRelease(a *Artist) int {
	a.Songs++
	return a.Songs
}
