// Ejercicio 5.3: Búsqueda (Lineal y Binaria)
// Implementa búsqueda lineal y búsqueda binaria en un slice de números.

package main

import (
	"fmt"
	"sort"
)

func main() {
	// Definimos un slice de ejemplo
	numeros := []int{23, 45, 12, 67, 34, 89, 5, 78, 90, 11}

	fmt.Println("=== ALGORITMOS DE BÚSQUEDA ===")
	fmt.Println("Arreglo de números:", numeros)

	var buscar int
	fmt.Print("\nIngresa el número a buscar: ")
	fmt.Scanln(&buscar)

	// 1. BÚSQUEDA LINEAL
	fmt.Println("\n--- BÚSQUEDA LINEAL ---")
	indiceLineal := busquedaLineal(numeros, buscar)

	if indiceLineal != -1 {
		fmt.Printf("✓ Número %d encontrado en la posición %d\n", buscar, indiceLineal)
	} else {
		fmt.Printf("✗ Número %d no encontrado\n", buscar)
	}

	// 2. BÚSQUEDA BINARIA
	fmt.Println("\n--- BÚSQUEDA BINARIA ---")
	fmt.Println("(Requiere que el arreglo esté ordenado)")

	// Primero ordenamos el slice para búsqueda binaria
	numerosOrdenados := make([]int, len(numeros))
	copy(numerosOrdenados, numeros)
	sort.Ints(numerosOrdenados)

	fmt.Println("Arreglo ordenado:", numerosOrdenados)

	indiceBinaria := busquedaBinaria(numerosOrdenados, buscar)

	if indiceBinaria != -1 {
		fmt.Printf("✓ Número %d encontrado en la posición %d (arreglo ordenado)\n", buscar, indiceBinaria)
	} else {
		fmt.Printf("✗ Número %d no encontrado\n", buscar)
	}

	// Comparación de eficiencia
	fmt.Println("\n--- COMPARACIÓN ---")
	fmt.Println("Búsqueda Lineal: O(n) - Revisa cada elemento secuencialmente")
	fmt.Println("Búsqueda Binaria: O(log n) - Divide el espacio de búsqueda a la mitad en cada paso")
	fmt.Println("La búsqueda binaria es más eficiente pero requiere un arreglo ordenado")
}

// busquedaLineal recorre el slice secuencialmente buscando el valor
// Retorna el índice si lo encuentra, -1 si no existe
func busquedaLineal(arr []int, valor int) int {
	for i, num := range arr {
		if num == valor {
			return i
		}
	}
	return -1
}

// busquedaBinaria busca en un slice ORDENADO dividiendo el espacio de búsqueda a la mitad
// Retorna el índice si lo encuentra, -1 si no existe
func busquedaBinaria(arr []int, valor int) int {
	izquierda := 0
	derecha := len(arr) - 1

	// Mientras haya elementos por revisar
	for izquierda <= derecha {
		// Calculamos el punto medio
		medio := (izquierda + derecha) / 2

		// Si encontramos el valor
		if arr[medio] == valor {
			return medio
		}

		// Si el valor es menor, buscamos en la mitad izquierda
		if arr[medio] > valor {
			derecha = medio - 1
		} else {
			// Si es mayor, buscamos en la mitad derecha
			izquierda = medio + 1
		}
	}

	// No se encontró el valor
	return -1
}
