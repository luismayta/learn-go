package variables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingle(t *testing.T) {
	assert.Equal(t, "itachi of konoha", getSingleVar(), "It should be itachi of konoha")
}
