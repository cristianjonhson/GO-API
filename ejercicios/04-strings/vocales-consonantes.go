// Ejercicio 4.1: Contador de vocales y consonantes
// Cuenta cuántas vocales y consonantes hay en un texto ingresado por el usuario.

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var texto string

	// Solicitamos el texto al usuario
	fmt.Print("Ingresa un texto: ")
	fmt.Scanln(&texto)

	// Convertimos a minúsculas para facilitar la comparación
	texto = strings.ToLower(texto)

	// Contadores
	vocales := 0
	consonantes := 0

	// Definimos las vocales
	esVocal := func(c rune) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
	}

	// Recorremos cada carácter del texto
	for _, caracter := range texto {
		// Solo contamos letras
		if unicode.IsLetter(caracter) {
			if esVocal(caracter) {
				vocales++
			} else {
				consonantes++
			}
		}
	}

	// Mostramos los resultados
	fmt.Println("\n--- RESULTADOS ---")
	fmt.Printf("Vocales: %d\n", vocales)
	fmt.Printf("Consonantes: %d\n", consonantes)
	fmt.Printf("Total de letras: %d\n", vocales+consonantes)
}
