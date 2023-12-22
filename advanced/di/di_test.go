package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiInitializeEvent(t *testing.T) {
	assert.NotEmpty(t, InitializeEvent("hi"), "It should not be empty")
}

func TestDiInitialize(t *testing.T) {
	Initialize()
}
