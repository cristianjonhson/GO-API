// Ejercicio 9.1: Parser seguro con manejo de errores
// Lee un string y lo convierte a int o float, con validación y manejo de errores.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== PARSER SEGURO ===")

	// Caso 1: Leer un número entero
	fmt.Println("\n--- CONVERSIÓN A ENTERO ---")
	entero := leerEnteroSeguro("Ingresa un número entero: ")
	fmt.Printf("✓ Número ingresado: %d\n", entero)

	// Caso 2: Leer un número decimal
	fmt.Println("\n--- CONVERSIÓN A DECIMAL ---")
	decimal := leerFloatSeguro("Ingresa un número decimal: ")
	fmt.Printf("✓ Número ingresado: %.2f\n", decimal)

	// Caso 3: Operaciones con los números
	fmt.Println("\n--- OPERACIONES ---")
	fmt.Printf("Suma: %d + %.2f = %.2f\n", entero, decimal, float64(entero)+decimal)
	fmt.Printf("Multiplicación: %d × %.2f = %.2f\n", entero, decimal, float64(entero)*decimal)
}

// leerEnteroSeguro solicita un string y lo convierte a entero
// Si falla la conversión, vuelve a pedir el valor hasta que sea válido
func leerEnteroSeguro(mensaje string) int {
	for {
		var input string
		fmt.Print(mensaje)
		fmt.Scanln(&input)

		// Intentamos convertir el string a int
		numero, err := stringAInt(input)

		if err != nil {
			// Si hay error, mostramos el mensaje y volvemos a pedir
			fmt.Printf("❌ Error: %v. Intenta de nuevo.\n", err)
			continue
		}

		// Si no hay error, retornamos el número
		return numero
	}
}

// leerFloatSeguro solicita un string y lo convierte a float64
// Si falla la conversión, vuelve a pedir el valor hasta que sea válido
func leerFloatSeguro(mensaje string) float64 {
	for {
		var input string
		fmt.Print(mensaje)
		fmt.Scanln(&input)

		// Intentamos convertir el string a float64
		numero, err := stringAFloat(input)

		if err != nil {
			// Si hay error, mostramos el mensaje y volvemos a pedir
			fmt.Printf("❌ Error: %v. Intenta de nuevo.\n", err)
			continue
		}

		// Si no hay error, retornamos el número
		return numero
	}
}

// stringAInt convierte un string a int
// Retorna el número y un error si la conversión falla
func stringAInt(s string) (int, error) {
	numero, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("'%s' no es un número entero válido", s)
	}
	return numero, nil
}

// stringAFloat convierte un string a float64
// Retorna el número y un error si la conversión falla
func stringAFloat(s string) (float64, error) {
	numero, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("'%s' no es un número decimal válido", s)
	}
	return numero, nil
}

// validarPositivo verifica que un número sea positivo
// Retorna error si el número es negativo o cero
func validarPositivo(numero float64) error {
	if numero <= 0 {
		return fmt.Errorf("el número debe ser positivo (recibido: %.2f)", numero)
	}
	return nil
}

// validarRangoNumerico verifica que un número esté en un rango
// Retorna error si está fuera del rango
func validarRangoNumerico(numero, min, max float64) error {
	if numero < min || numero > max {
		return fmt.Errorf("el número %.2f está fuera del rango [%.2f, %.2f]", numero, min, max)
	}
	return nil
}
