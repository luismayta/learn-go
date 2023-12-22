package sample01

import (
	"time"
)

type IGreeter interface {
	Greet() IMessage
}

func NewGreeter(m IMessage) IGreeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return &Greeter{Message: m, Grumpy: grumpy}
}

type Greeter struct {
	Message IMessage
	Grumpy  bool
}

func (g Greeter) Greet() IMessage {
	if g.Grumpy {
		return NewMessage("Go away!")
	}
	return g.Message
}
