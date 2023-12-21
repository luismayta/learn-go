package variables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnonym(t *testing.T) {
	assert.Equal(t, "Itachi Uchiha", anonym(), "It should be Itachi Uchiha")
}
