package beginner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrings(t *testing.T) {
	assert.Equal(
		t,
		"Hola, me llamo LM",
		getGreeting(),
		"It should be Hola, me llamo LM",
	)
	assert.Equal(t, "Luis Mayta", getGreetingName(), "It should be Luis Mayta")
}
