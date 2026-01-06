// Ejercicio 6.2: Validador de rangos
// Implementa una función de validación reutilizable para diferentes casos.

package main

import "fmt"

func main() {
	fmt.Println("=== SISTEMA DE VALIDACIÓN ===")

	// Caso 1: Validar edad (0-120)
	fmt.Println("\n--- VALIDAR EDAD ---")
	edad := solicitarValorEnRango("Ingresa tu edad: ", 0, 120)
	fmt.Printf("✓ Edad válida: %d años\n", edad)

	// Caso 2: Validar nota (1.0-7.0)
	fmt.Println("\n--- VALIDAR NOTA ---")
	nota := solicitarValorEnRangoFloat("Ingresa tu nota (1.0-7.0): ", 1.0, 7.0)
	fmt.Printf("✓ Nota válida: %.1f\n", nota)

	// Clasificamos la nota
	if nota >= 4.0 {
		fmt.Println("Estado: APROBADO ✓")
	} else {
		fmt.Println("Estado: REPROBADO ✗")
	}

	// Caso 3: Validar día de la semana (1-7)
	fmt.Println("\n--- VALIDAR DÍA DE LA SEMANA ---")
	dia := solicitarValorEnRango("Ingresa número de día (1-7): ", 1, 7)
	nombreDia := obtenerNombreDia(dia)
	fmt.Printf("✓ Día válido: %s\n", nombreDia)

	// Caso 4: Validar temperatura (-50 a 50)
	fmt.Println("\n--- VALIDAR TEMPERATURA ---")
	temp := solicitarValorEnRangoFloat("Ingresa temperatura en °C (-50 a 50): ", -50, 50)
	fmt.Printf("✓ Temperatura válida: %.1f°C\n", temp)
	clasificarTemperatura(temp)
}

// validarRango verifica si un valor entero está dentro de un rango [min, max]
// Retorna true si está en el rango, false en caso contrario
func validarRango(valor, min, max int) bool {
	return valor >= min && valor <= max
}

// validarRangoFloat verifica si un valor float está dentro de un rango [min, max]
// Retorna true si está en el rango, false en caso contrario
func validarRangoFloat(valor, min, max float64) bool {
	return valor >= min && valor <= max
}

// solicitarValorEnRango pide un número entero hasta que esté en el rango válido
// Retorna el valor validado
func solicitarValorEnRango(mensaje string, min, max int) int {
	var valor int
	for {
		fmt.Print(mensaje)
		fmt.Scanln(&valor)

		if validarRango(valor, min, max) {
			return valor
		}
		fmt.Printf("⚠️  Error: El valor debe estar entre %d y %d. Intenta de nuevo.\n", min, max)
	}
}

// solicitarValorEnRangoFloat pide un número float hasta que esté en el rango válido
// Retorna el valor validado
func solicitarValorEnRangoFloat(mensaje string, min, max float64) float64 {
	var valor float64
	for {
		fmt.Print(mensaje)
		fmt.Scanln(&valor)

		if validarRangoFloat(valor, min, max) {
			return valor
		}
		fmt.Printf("⚠️  Error: El valor debe estar entre %.1f y %.1f. Intenta de nuevo.\n", min, max)
	}
}

// obtenerNombreDia convierte un número (1-7) en el nombre del día
func obtenerNombreDia(dia int) string {
	dias := []string{"Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado", "Domingo"}
	if dia >= 1 && dia <= 7 {
		return dias[dia-1]
	}
	return "Desconocido"
}

// clasificarTemperatura da una descripción según la temperatura
func clasificarTemperatura(temp float64) {
	switch {
	case temp < 0:
		fmt.Println("🥶 Hace mucho frío")
	case temp < 15:
		fmt.Println("❄️  Hace frío")
	case temp < 25:
		fmt.Println("😊 Temperatura agradable")
	case temp < 35:
		fmt.Println("☀️  Hace calor")
	default:
		fmt.Println("🔥 Hace mucho calor")
	}
}
