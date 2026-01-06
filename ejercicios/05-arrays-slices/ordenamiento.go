// Ejercicio 5.2: Ordenamiento (Bubble Sort)
// Implementa el algoritmo Bubble Sort para ordenar un slice de números.

package main

import "fmt"

func main() {
	var cantidad int

	// Solicitamos la cantidad de números
	fmt.Print("¿Cuántos números deseas ordenar? ")
	fmt.Scanln(&cantidad)

	if cantidad < 1 {
		fmt.Println("Debes ingresar al menos 1 número")
		return
	}

	// Creamos un slice para almacenar los números
	numeros := make([]int, cantidad)

	// Pedimos cada número
	fmt.Println("\nIngresa los números:")
	for i := 0; i < cantidad; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scanln(&numeros[i])
	}

	// Mostramos el arreglo original
	fmt.Println("\nArreglo original:", numeros)

	// Implementación de Bubble Sort
	// Comparamos cada par de elementos adyacentes y los intercambiamos si están en orden incorrecto
	for i := 0; i < len(numeros)-1; i++ {
		// En cada pasada, el elemento más grande "burbujea" hacia el final
		for j := 0; j < len(numeros)-1-i; j++ {
			// Si el elemento actual es mayor que el siguiente, los intercambiamos
			if numeros[j] > numeros[j+1] {
				// Intercambio (swap) de elementos
				numeros[j], numeros[j+1] = numeros[j+1], numeros[j]
			}
		}
	}

	// Mostramos el arreglo ordenado
	fmt.Println("Arreglo ordenado:", numeros)

	// Información adicional
	fmt.Println("\n--- INFO ---")
	fmt.Printf("Se realizaron %d pasadas para ordenar los elementos\n", cantidad-1)
	fmt.Println("Algoritmo utilizado: Bubble Sort")
}
