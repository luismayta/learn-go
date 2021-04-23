package collection

func getMapBasic() map[string]int {
	m := make(map[string]int)
	m["k1"] = 5
	m["k2"] = 6
	m["k3"] = 6

	return m
}

func getPlanets() map[string]string {
	planets := make(map[string]string)
	planets["Han Solo"] = "Corellia"
	planets["Luke"] = "Tatoine"
	planets["Leia Oegana"] = "Tatoine"
	delete(planets, "Han Solo")

	return planets
}
