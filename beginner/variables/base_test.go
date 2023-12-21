package variables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsString(i interface{}) bool {
	_, ok := i.(string)
	return ok
}

func IsInt(i interface{}) bool {
	_, ok := i.(int)
	return ok
}

func TestBase(t *testing.T) {
	assert.True(t, IsString(name), "Name should be a string")
	assert.True(t, IsString(location), "Location should be a string")
	assert.True(t, IsInt(age), "Location should be a int")
}
