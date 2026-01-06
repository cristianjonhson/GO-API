// Ejercicio 9.2: Función dividir con manejo de errores
// Implementa una función de división que retorna error en caso de división por cero.

package main

import (
	"fmt"
)

const errorFormat = "❌ Error: %v\n"

func main() {
	fmt.Println("=== CALCULADORA CON MANEJO DE ERRORES ===")

	// Solicitamos los números
	fmt.Print("Ingresa el primer número: ")
	var a float64
	fmt.Scanln(&a)

	fmt.Print("Ingresa el segundo número: ")
	var b float64
	fmt.Scanln(&b)

	// Intentamos realizar operaciones con manejo de errores

	// División
	resultado, err := dividir(a, b)
	if err != nil {
		fmt.Printf("❌ Error en división: %v\n", err)
	} else {
		fmt.Printf("✓ %.2f ÷ %.2f = %.2f\n", a, b, resultado)
	}

	// Raíz cuadrada
	raiz, err := raizCuadrada(a)
	if err != nil {
		fmt.Printf("❌ Error en raíz cuadrada de %.2f: %v\n", a, err)
	} else {
		fmt.Printf("✓ √%.2f = %.2f\n", a, raiz)
	}

	// Porcentaje
	porcentaje, err := calcularPorcentaje(a, b)
	if err != nil {
		fmt.Printf("❌ Error al calcular porcentaje: %v\n", err)
	} else {
		fmt.Printf("✓ %.2f es el %.2f%% de %.2f\n", a, porcentaje, b)
	}

	// Ejemplo de uso múltiple con validación
	fmt.Println("\n--- CALCULADORA INTERACTIVA ---")
	ejecutarCalculadora()
}

// dividir realiza la división de dos números
// Retorna el resultado y un error si el divisor es cero
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return a / b, nil
}

// raizCuadrada calcula la raíz cuadrada de un número
// Retorna error si el número es negativo (en matemática real)
func raizCuadrada(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("no se puede calcular raíz cuadrada de número negativo (%.2f)", n)
	}

	// Implementación simple usando el método de Newton-Raphson
	if n == 0 {
		return 0, nil
	}

	estimacion := n / 2
	for i := 0; i < 10; i++ {
		estimacion = (estimacion + n/estimacion) / 2
	}

	return estimacion, nil
}

// calcularPorcentaje calcula qué porcentaje representa 'parte' del 'total'
// Retorna error si el total es cero
func calcularPorcentaje(parte, total float64) (float64, error) {
	if total == 0 {
		return 0, fmt.Errorf("el total no puede ser cero")
	}
	return (parte / total) * 100, nil
}

// potencia calcula a elevado a la n
// Retorna error si el exponente es negativo con base cero
func potencia(base, exponente float64) (float64, error) {
	if base == 0 && exponente < 0 {
		return 0, fmt.Errorf("0 elevado a exponente negativo es indefinido")
	}

	resultado := 1.0
	expAbs := exponente
	if exponente < 0 {
		expAbs = -exponente
	}

	for i := 0.0; i < expAbs; i++ {
		resultado *= base
	}

	if exponente < 0 {
		resultado = 1 / resultado
	}

	return resultado, nil
}

// ejecutarCalculadora ejecuta un menú interactivo de calculadora
func ejecutarCalculadora() {
	for {
		opcion := mostrarMenuYObtenerOpcion()

		if opcion == 5 {
			fmt.Println("¡Hasta luego! 👋")
			return
		}

		ejecutarOperacion(opcion)
	}
}

// mostrarMenuYObtenerOpcion muestra el menú y retorna la opción seleccionada
func mostrarMenuYObtenerOpcion() int {
	fmt.Println("\nOperaciones disponibles:")
	fmt.Println("1. División")
	fmt.Println("2. Raíz cuadrada")
	fmt.Println("3. Calcular porcentaje")
	fmt.Println("4. Potencia")
	fmt.Println("5. Salir")

	var opcion int
	fmt.Print("\nSelecciona una operación: ")
	fmt.Scanln(&opcion)
	return opcion
}

// ejecutarOperacion ejecuta la operación seleccionada
func ejecutarOperacion(opcion int) {
	switch opcion {
	case 1:
		operacionDivision()
	case 2:
		operacionRaizCuadrada()
	case 3:
		operacionPorcentaje()
	case 4:
		operacionPotencia()
	default:
		fmt.Println("❌ Opción inválida")
	}
}

// operacionDivision ejecuta la operación de división
func operacionDivision() {
	var a, b float64
	fmt.Print("Dividendo: ")
	fmt.Scanln(&a)
	fmt.Print("Divisor: ")
	fmt.Scanln(&b)

	if resultado, err := dividir(a, b); err != nil {
		fmt.Printf(errorFormat, err)
	} else {
		fmt.Printf("✓ Resultado: %.2f ÷ %.2f = %.2f\n", a, b, resultado)
	}
}

// operacionRaizCuadrada ejecuta la operación de raíz cuadrada
func operacionRaizCuadrada() {
	var n float64
	fmt.Print("Número: ")
	fmt.Scanln(&n)

	if resultado, err := raizCuadrada(n); err != nil {
		fmt.Printf(errorFormat, err)
	} else {
		fmt.Printf("✓ Resultado: √%.2f = %.2f\n", n, resultado)
	}
}

// operacionPorcentaje ejecuta la operación de cálculo de porcentaje
func operacionPorcentaje() {
	var parte, total float64
	fmt.Print("Parte: ")
	fmt.Scanln(&parte)
	fmt.Print("Total: ")
	fmt.Scanln(&total)

	if resultado, err := calcularPorcentaje(parte, total); err != nil {
		fmt.Printf(errorFormat, err)
	} else {
		fmt.Printf("✓ %.2f es el %.2f%% de %.2f\n", parte, resultado, total)
	}
}

// operacionPotencia ejecuta la operación de potencia
func operacionPotencia() {
	var base, exponente float64
	fmt.Print("Base: ")
	fmt.Scanln(&base)
	fmt.Print("Exponente: ")
	fmt.Scanln(&exponente)

	if resultado, err := potencia(base, exponente); err != nil {
		fmt.Printf(errorFormat, err)
	} else {
		fmt.Printf("✓ %.2f^%.2f = %.2f\n", base, exponente, resultado)
	}
}
