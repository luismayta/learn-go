package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMapBasic(t *testing.T) {
	resultado := getMapBasic()

	assert.Len(t, resultado, 3, "La longitud del mapa no es la esperada")

	assert.Equal(t, 5, resultado["k1"], "Valor de k1 no es el esperado")
	assert.Equal(t, 6, resultado["k2"], "Valor de k2 no es el esperado")
	assert.Equal(t, 6, resultado["k3"], "Valor de k3 no es el esperado")
}

func TestGetPlanets(t *testing.T) {
	resultado := getPlanets()

	assert.Len(t, resultado, 2, "La longitud del mapa no es la esperada")

	assert.Equal(t, "Tatoine", resultado["Luke"], "Valor de Luke no es el esperado")
	assert.Equal(t, "Tatoine", resultado["Leia Oegana"], "Valor de Leia Oegana no es el esperado")

	_, existe := resultado["Han Solo"]
	assert.False(t, existe, "Han Solo debería haber sido eliminado")
}
