// Ejercicio 11: Tests de funciones puras
// Implementa tests unitarios para funciones comunes con casos borde.

package main

import "fmt"

// esPalindromo verifica si un string es palíndromo (se lee igual al revés)
func esPalindromo(s string) bool {
	// Convertimos a minúsculas y removemos espacios
	var limpio string
	for _, char := range s {
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char >= '0' && char <= '9' {
			if char >= 'A' && char <= 'Z' {
				char = char + 32 // Convertir a minúscula
			}
			limpio += string(char)
		}
	}

	// Comparamos desde ambos extremos
	longitud := len(limpio)
	for i := 0; i < longitud/2; i++ {
		if limpio[i] != limpio[longitud-1-i] {
			return false
		}
	}
	return true
}

// max encuentra el valor máximo en un slice
func max(numeros []float64) (float64, error) {
	if len(numeros) == 0 {
		return 0, fmt.Errorf("slice vacío")
	}

	maximo := numeros[0]
	for _, num := range numeros {
		if num > maximo {
			maximo = num
		}
	}
	return maximo, nil
}

// promedio calcula el promedio de un slice de números
func promedio(numeros []float64) (float64, error) {
	if len(numeros) == 0 {
		return 0, fmt.Errorf("slice vacío")
	}

	suma := 0.0
	for _, num := range numeros {
		suma += num
	}
	return suma / float64(len(numeros)), nil
}

// dividir realiza división con manejo de errores
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("división por cero")
	}
	return a / b, nil
}

// factorial calcula el factorial de un número
func factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial de número negativo no está definido")
	}
	if n == 0 || n == 1 {
		return 1, nil
	}

	resultado := 1
	for i := 2; i <= n; i++ {
		resultado *= i
	}
	return resultado, nil
}

// esPrimo verifica si un número es primo
func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Esta es solo la implementación. Los tests están en funciones_test.go
func main() {
	fmt.Println("=== FUNCIONES PARA TESTING ===")
	fmt.Println("\nEste archivo contiene funciones puras para ser testeadas.")
	fmt.Println("Ejecuta: go test -v")
	fmt.Println("\nFunciones disponibles:")
	fmt.Println("- esPalindromo(s string) bool")
	fmt.Println("- max(numeros []float64) (float64, error)")
	fmt.Println("- promedio(numeros []float64) (float64, error)")
	fmt.Println("- dividir(a, b float64) (float64, error)")
	fmt.Println("- factorial(n int) (int, error)")
	fmt.Println("- esPrimo(n int) bool")
}
