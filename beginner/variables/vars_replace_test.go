package variables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarsReplace(t *testing.T) {
	assert.Equal(t, "Hello itachi uchiha", getMessage(), "It should be Hello itachi uchiha")
}
