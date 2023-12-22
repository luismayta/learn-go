package di

func Initialize() {
	e := InitializeEvent("Hi there!")
	e.Start()
}
