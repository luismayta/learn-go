package di

import (
	"log"
)

type IEvent interface {
	Start()
}

func NewEvent(g IGreeter) IEvent {
	return &Event{Greeter: g}
}

type Event struct {
	Greeter IGreeter
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	log.Println(msg.Value())
}
