// Ejercicio 6.1: Refactorización con funciones
// Demuestra el principio de modularidad usando funciones reutilizables.

package main

import "fmt"

func main() {
	fmt.Println("=== ESTADÍSTICAS CON FUNCIONES ===")

	// Solicitamos la cantidad de números
	cantidad := leerNumero("¿Cuántos números vas a ingresar? ")

	if cantidad < 1 {
		fmt.Println("Debes ingresar al menos 1 número")
		return
	}

	// Leemos los números
	numeros := leerLista(cantidad)

	// Mostramos la lista
	fmt.Println("\nNúmeros ingresados:", numeros)

	// Calculamos y mostramos estadísticas usando funciones modulares
	fmt.Println("\n--- ESTADÍSTICAS ---")
	fmt.Printf("Máximo: %.2f\n", max(numeros))
	fmt.Printf("Mínimo: %.2f\n", min(numeros))
	fmt.Printf("Promedio: %.2f\n", promedio(numeros))
	fmt.Printf("Suma total: %.2f\n", suma(numeros))

	// Analizamos paridad de números enteros
	fmt.Println("\n--- ANÁLISIS DE PARIDAD ---")
	for _, num := range numeros {
		tipoNum := "IMPAR"
		if esPar(int(num)) {
			tipoNum = "PAR"
		}
		fmt.Printf("%.0f es %s\n", num, tipoNum)
	}
}

// leerNumero solicita un número entero con un mensaje personalizado
// Retorna el número ingresado por el usuario
func leerNumero(mensaje string) int {
	var numero int
	fmt.Print(mensaje)
	fmt.Scanln(&numero)
	return numero
}

// leerLista solicita N números y los retorna en un slice
func leerLista(cantidad int) []float64 {
	numeros := make([]float64, cantidad)
	fmt.Println("\nIngresa los números:")
	for i := 0; i < cantidad; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scanln(&numeros[i])
	}
	return numeros
}

// esPar determina si un número es par
// Retorna true si es par, false si es impar
func esPar(n int) bool {
	return n%2 == 0
}

// max encuentra el valor máximo en un slice de números
func max(lista []float64) float64 {
	if len(lista) == 0 {
		return 0
	}
	maximo := lista[0]
	for _, num := range lista {
		if num > maximo {
			maximo = num
		}
	}
	return maximo
}

// min encuentra el valor mínimo en un slice de números
func min(lista []float64) float64 {
	if len(lista) == 0 {
		return 0
	}
	minimo := lista[0]
	for _, num := range lista {
		if num < minimo {
			minimo = num
		}
	}
	return minimo
}

// suma calcula la suma de todos los elementos del slice
func suma(lista []float64) float64 {
	total := 0.0
	for _, num := range lista {
		total += num
	}
	return total
}

// promedio calcula el promedio de los números en el slice
func promedio(lista []float64) float64 {
	if len(lista) == 0 {
		return 0
	}
	return suma(lista) / float64(len(lista))
}
