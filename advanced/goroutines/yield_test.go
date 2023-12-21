package goroutines

import (
	"testing"
)

func TestFib(t *testing.T) {
	resultadoEsperado := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	c := fib(10)

	for _, esperado := range resultadoEsperado {
		obtenido, ok := <-c
		if !ok {
			t.Error("El canal se cerró inesperadamente.")
		}

		if obtenido != esperado {
			t.Errorf("Error: Esperado %d, obtenido %d", esperado, obtenido)
		}
	}

	_, ok := <-c
	if ok {
		t.Error("Se esperaba que el canal estuviera cerrado.")
	}
}

func TestFibBest(t *testing.T) {
	resultadoEsperado := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	c := make(chan int)
	go fibBest(c, 10)

	for _, esperado := range resultadoEsperado {
		obtenido, ok := <-c
		if !ok {
			t.Error("El canal se cerró inesperadamente.")
		}

		if obtenido != esperado {
			t.Errorf("Error: Esperado %d, obtenido %d", esperado, obtenido)
		}
	}

	_, ok := <-c
	if ok {
		t.Error("Se esperaba que el canal estuviera cerrado.")
	}
}

func TestGetFib(t *testing.T) {
	// Se realiza la prueba simplemente para verificar que no haya panics.
	getFib()
}

func TestGetFibBest(t *testing.T) {
	// Se realiza la prueba simplemente para verificar que no haya panics.
	getFibBest()
}
