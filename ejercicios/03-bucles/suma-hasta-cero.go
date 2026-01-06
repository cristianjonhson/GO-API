// Ejercicio 3.2: Suma hasta cero
// Lee números hasta que el usuario ingrese 0. Calcula suma y promedio de los números ingresados.

package main

import "fmt"

func main() {
	var numero float64
	var suma float64 = 0
	var contador int = 0

	fmt.Println("=== SUMA HASTA CERO ===")
	fmt.Println("Ingresa números (0 para terminar)")
	fmt.Println()

	// Bucle que continúa hasta que se ingrese 0
	for {
		fmt.Print("Número: ")
		fmt.Scanln(&numero)

		// Si ingresa 0, salimos del bucle
		if numero == 0 {
			break
		}

		// Sumamos el número y aumentamos el contador
		suma += numero
		contador++
	}

	// Mostramos los resultados
	fmt.Println("\n--- RESULTADOS ---")
	fmt.Printf("Cantidad de números ingresados: %d\n", contador)
	fmt.Printf("Suma total: %.2f\n", suma)

	// Calculamos y mostramos el promedio (si se ingresó al menos un número)
	if contador > 0 {
		promedio := suma / float64(contador)
		fmt.Printf("Promedio: %.2f\n", promedio)
	} else {
		fmt.Println("No se ingresaron números para calcular el promedio.")
	}
}
