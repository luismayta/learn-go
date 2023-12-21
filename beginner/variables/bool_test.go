package variables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	assert.True(t, casiQueSi, "It should be true")
	assert.False(t, casiQueNo, "It should be false")
}
