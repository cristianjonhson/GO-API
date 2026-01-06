// Ejercicio 2.2: Clasificador de notas
// Clasifica una nota (1.0-7.0) como "Reprobado", "Aprobado" o "Excelente".

package main

import "fmt"

func main() {
	var nota float64

	// Solicitamos la nota al usuario
	fmt.Print("Ingresa la nota (1.0 - 7.0): ")
	fmt.Scanln(&nota)

	// Validamos que la nota esté en el rango correcto
	if nota < 1.0 || nota > 7.0 {
		fmt.Println("Error: La nota debe estar entre 1.0 y 7.0")
		return
	}

	// Clasificamos la nota según rangos
	var clasificacion string
	var estado string

	if nota < 4.0 {
		// Reprobado: notas menores a 4.0
		clasificacion = "REPROBADO"
		estado = "❌"
	} else if nota < 6.0 {
		// Aprobado: notas entre 4.0 y 5.9
		clasificacion = "APROBADO"
		estado = "✓"
	} else {
		// Excelente: notas entre 6.0 y 7.0
		clasificacion = "EXCELENTE"
		estado = "⭐"
	}

	// Imprimimos el resultado
	fmt.Printf("\n%s Nota: %.1f - %s\n", estado, nota, clasificacion)
}
