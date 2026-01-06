// Ejercicio 4.2: Palíndromo
// Verifica si una palabra o frase es un palíndromo (se lee igual al derecho y al revés).
// Ignora espacios, mayúsculas y signos de puntuación.

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var texto string

	// Solicitamos el texto al usuario
	fmt.Print("Ingresa una palabra o frase: ")
	// Usamos Scan para leer toda la línea incluyendo espacios
	fmt.Scanln(&texto)

	// Limpiamos el texto: removemos espacios y convertimos a minúsculas
	textoLimpio := ""
	for _, caracter := range strings.ToLower(texto) {
		// Solo consideramos letras y números
		if unicode.IsLetter(caracter) || unicode.IsDigit(caracter) {
			textoLimpio += string(caracter)
		}
	}

	// Verificamos si es palíndromo comparando con su reverso
	longitud := len(textoLimpio)
	esPalindromo := true

	// Comparamos caracteres desde ambos extremos hacia el centro
	for i := 0; i < longitud/2; i++ {
		if textoLimpio[i] != textoLimpio[longitud-1-i] {
			esPalindromo = false
			break
		}
	}

	// Mostramos el resultado
	fmt.Println("\n--- RESULTADO ---")
	fmt.Printf("Texto original: \"%s\"\n", texto)
	fmt.Printf("Texto limpio: \"%s\"\n", textoLimpio)

	if esPalindromo {
		fmt.Println("✓ ¡Es un PALÍNDROMO! 🎉")
	} else {
		fmt.Println("✗ NO es un palíndromo")
	}
}
