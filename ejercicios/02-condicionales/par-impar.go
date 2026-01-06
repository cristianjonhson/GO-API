// Ejercicio 2.1: Par o impar + positivo/negativo
// Clasifica un número como par/impar y positivo/negativo/cero.

package main

import "fmt"

func main() {
	var numero int

	// Solicitamos un número al usuario
	fmt.Print("Ingresa un número entero: ")
	fmt.Scanln(&numero)

	// Determinamos si es par o impar
	var tipoPar string
	if numero%2 == 0 {
		tipoPar = "PAR"
	} else {
		tipoPar = "IMPAR"
	}

	// Determinamos si es positivo, negativo o cero
	var tipoSigno string
	if numero > 0 {
		tipoSigno = "POSITIVO"
	} else if numero < 0 {
		tipoSigno = "NEGATIVO"
	} else {
		tipoSigno = "CERO"
	}

	// Imprimimos la clasificación
	fmt.Printf("\nEl número %d es: %s y %s\n", numero, tipoPar, tipoSigno)
}
