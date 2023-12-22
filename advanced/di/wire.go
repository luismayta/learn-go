//go:build wireinject
// +build wireinject

package di

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
