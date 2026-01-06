// Ejercicio 4.3: Frecuencia de palabras
// Cuenta cuántas veces aparece cada palabra en un texto.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Usamos bufio.Scanner para leer líneas completas
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("=== CONTADOR DE FRECUENCIA DE PALABRAS ===")
	fmt.Println("Ingresa un texto (presiona Enter al finalizar):")

	var texto string
	if scanner.Scan() {
		texto = scanner.Text()
	}

	// Convertimos a minúsculas y dividimos el texto en palabras
	palabras := strings.Fields(strings.ToLower(texto))

	// Mapa para almacenar la frecuencia de cada palabra
	frecuencia := make(map[string]int)

	// Contamos cada palabra
	for _, palabra := range palabras {
		// Limpiamos signos de puntuación básicos
		palabra = strings.Trim(palabra, ".,;:!?¿¡\"'")
		if palabra != "" {
			frecuencia[palabra]++
		}
	}

	// Mostramos los resultados
	fmt.Println("\n--- FRECUENCIA DE PALABRAS ---")
	fmt.Printf("Total de palabras: %d\n", len(palabras))
	fmt.Printf("Palabras únicas: %d\n\n", len(frecuencia))

	// Imprimimos cada palabra con su frecuencia
	for palabra, cantidad := range frecuencia {
		fmt.Printf("'%s': %d vez/veces\n", palabra, cantidad)
	}
}
