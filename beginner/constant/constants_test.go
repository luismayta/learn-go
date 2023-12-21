package constant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstants(t *testing.T) {
	const epsilon = 1e-6
	assert.InEpsilon(t, 3.1416, Pi, epsilon)
	assert.Equal(t, 200, StatusOk)
	assert.Equal(t, 404, StatusNotFound)
}
