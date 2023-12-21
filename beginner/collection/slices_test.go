package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanicSlice(t *testing.T) {
	response := PanicSlice()
	assert.NotEmpty(t, response, response)
}

func TestExample01(t *testing.T) {
	val1, val2 := Example1()
	assert.Equal(t, 2, val1)
	assert.Equal(t, 5, val2)
}

func TestExample02(t *testing.T) {
	personajes := Example2()
	assert.NotEmpty(t, personajes)
}
