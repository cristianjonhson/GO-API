// Ejercicio 3.1: Tabla de multiplicar
// Genera la tabla de multiplicar de un número N del 1 al 10.

package main

import "fmt"

func main() {
	var numero int

	// Solicitamos el número al usuario
	fmt.Print("Ingresa un número para ver su tabla de multiplicar: ")
	fmt.Scanln(&numero)

	// Imprimimos encabezado
	fmt.Printf("\n=== TABLA DEL %d ===\n\n", numero)

	// Generamos la tabla del 1 al 10 usando un bucle for
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}

	fmt.Println("\n==================")
}
