// Ejercicio 5.1: Máximo, mínimo y promedio
// Calcula el número máximo, mínimo y el promedio de una lista de números.

package main

import "fmt"

func main() {
	var cantidad int

	// Solicitamos la cantidad de números
	fmt.Print("¿Cuántos números vas a ingresar? ")
	fmt.Scanln(&cantidad)

	// Validamos que sea al menos 1
	if cantidad < 1 {
		fmt.Println("Debes ingresar al menos 1 número")
		return
	}

	// Creamos un slice para almacenar los números
	numeros := make([]float64, cantidad)

	// Pedimos cada número
	fmt.Println("\nIngresa los números:")
	for i := 0; i < cantidad; i++ {
		fmt.Printf("Número %d: ", i+1)
		fmt.Scanln(&numeros[i])
	}

	// Inicializamos máximo y mínimo con el primer elemento
	maximo := numeros[0]
	minimo := numeros[0]
	suma := 0.0

	// Recorremos el slice para encontrar máximo, mínimo y calcular suma
	for _, numero := range numeros {
		// Actualizamos máximo
		if numero > maximo {
			maximo = numero
		}
		// Actualizamos mínimo
		if numero < minimo {
			minimo = numero
		}
		// Sumamos para el promedio
		suma += numero
	}

	// Calculamos el promedio
	promedio := suma / float64(cantidad)

	// Mostramos los resultados
	fmt.Println("\n--- ESTADÍSTICAS ---")
	fmt.Printf("Cantidad de números: %d\n", cantidad)
	fmt.Printf("Número máximo: %.2f\n", maximo)
	fmt.Printf("Número mínimo: %.2f\n", minimo)
	fmt.Printf("Promedio: %.2f\n", promedio)
	fmt.Printf("Suma total: %.2f\n", suma)
}
