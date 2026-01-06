// Ejercicio 1.2: Conversor de unidades
// Convierte °C→°F, km→millas, CLP→USD usando valores ingresados por el usuario.

package main

import "fmt"

func main() {
	// Constantes para conversiones
	const (
		factorFahrenheit = 9.0 / 5.0
		offsetFahrenheit = 32.0
		factorMillas     = 0.621371
		tasaCambioUSD    = 0.0011 // CLP a USD (tasa aproximada, ajustar según necesidad)
	)

	var opcion int

	// Menú de opciones
	fmt.Println("=== CONVERSOR DE UNIDADES ===")
	fmt.Println("1. Celsius a Fahrenheit")
	fmt.Println("2. Kilómetros a Millas")
	fmt.Println("3. CLP a USD")
	fmt.Print("Selecciona una opción (1-3): ")
	fmt.Scanln(&opcion)

	// Switch para manejar la conversión seleccionada
	switch opcion {
	case 1:
		// Conversión °C → °F
		var celsius float64
		fmt.Print("Ingresa temperatura en Celsius: ")
		fmt.Scanln(&celsius)
		fahrenheit := celsius*factorFahrenheit + offsetFahrenheit
		fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)

	case 2:
		// Conversión km → millas
		var kilometros float64
		fmt.Print("Ingresa distancia en kilómetros: ")
		fmt.Scanln(&kilometros)
		millas := kilometros * factorMillas
		fmt.Printf("%.2f km = %.2f millas\n", kilometros, millas)

	case 3:
		// Conversión CLP → USD
		var clp float64
		fmt.Print("Ingresa cantidad en CLP: ")
		fmt.Scanln(&clp)
		usd := clp * tasaCambioUSD
		fmt.Printf("$%.2f CLP = $%.2f USD\n", clp, usd)

	default:
		fmt.Println("Opción inválida. Por favor selecciona 1, 2 o 3.")
	}
}
