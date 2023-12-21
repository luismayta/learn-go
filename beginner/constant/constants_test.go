package constant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	assert.Equal(t, 3.1416, Pi)
	assert.Equal(t, 200, StatusOk)
	assert.Equal(t, 404, StatusNotFound)
}
