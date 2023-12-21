package goroutine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoroutine(t *testing.T) {
	assert.Equal(t, "Hello Lucho", execute(), "It should be Hello Lucho")
}
