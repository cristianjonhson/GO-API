// Ejercicio 1.3: Calculadora simple
// Lee 2 números y una operación (+, -, *, /), realiza el cálculo y valida división por cero.

package main

import "fmt"

func main() {
	// Variables para los números y la operación
	var num1, num2 float64
	var operacion string

	// Solicitamos el primer número
	fmt.Print("Ingresa el primer número: ")
	fmt.Scanln(&num1)

	// Solicitamos el segundo número
	fmt.Print("Ingresa el segundo número: ")
	fmt.Scanln(&num2)

	// Solicitamos la operación
	fmt.Print("Ingresa la operación (+, -, *, /): ")
	fmt.Scanln(&operacion)

	// Variable para almacenar el resultado
	var resultado float64
	var esValido bool = true

	// Realizamos la operación según el operador ingresado
	switch operacion {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		// Validación de división por cero
		if num2 == 0 {
			fmt.Println("Error: No se puede dividir por cero")
			esValido = false
		} else {
			resultado = num1 / num2
		}
	default:
		fmt.Println("Error: Operación no válida. Usa +, -, * o /")
		esValido = false
	}

	// Si la operación es válida, mostramos el resultado
	if esValido {
		fmt.Printf("\nResultado: %.2f %s %.2f = %.2f\n", num1, operacion, num2, resultado)
	}
}
