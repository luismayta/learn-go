//go:build wireinject
// +build wireinject

package sample01

import (
	"github.com/google/wire"
)

// wire_gen.go
func InitializeEvent(phrase string) IEvent {
	wire.Build(NewEvent,
		NewGreeter,
		NewMessage)
	return &Event{}
}