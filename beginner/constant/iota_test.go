package constant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIota(t *testing.T) {
	const epsilon = 1e-6
	assert.InEpsilon(t, 3.1416, pi, epsilon)
	assert.InEpsilon(t, 3.1416, igualAPI, epsilon)
}
