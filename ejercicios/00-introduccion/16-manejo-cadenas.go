// Ejemplo básico de manejo de cadenas en Go
package main

import (
	"fmt"
	"strings"
)

func main() {
	cadena := "Hola, mundo"

	// Concatenación
	concatenada := cadena + "!"
	fmt.Println("Concatenada:", concatenada)

	// Búsqueda
	contiene := strings.Contains(cadena, "mundo")
	fmt.Println("Contiene 'mundo':", contiene)

	// Reemplazo
	reemplazada := strings.ReplaceAll(cadena, "mundo", "Go")
	fmt.Println("Reemplazada:", reemplazada)
}