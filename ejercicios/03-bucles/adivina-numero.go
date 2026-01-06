// Ejercicio 3.3: Adivina el número
// Genera un número aleatorio entre 1 y 100. El usuario debe adivinarlo con pistas de "más" o "menos".

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inicializamos el generador de números aleatorios con la hora actual
	rand.Seed(time.Now().UnixNano())

	// Generamos un número aleatorio entre 1 y 100
	numeroSecreto := rand.Intn(100) + 1
	intentos := 0
	adivinado := false

	fmt.Println("=== ADIVINA EL NÚMERO ===")
	fmt.Println("He pensado un número entre 1 y 100")
	fmt.Println("¡Intenta adivinarlo!")
	fmt.Println()

	// Bucle del juego
	for !adivinado {
		var intento int
		intentos++

		// Solicitamos el número al usuario
		fmt.Printf("Intento #%d - Ingresa tu número: ", intentos)
		fmt.Scanln(&intento)

		// Validamos el rango
		if intento < 1 || intento > 100 {
			fmt.Println("⚠️  El número debe estar entre 1 y 100")
			intentos-- // No contamos este intento
			continue
		}

		// Comparamos con el número secreto y damos pistas
		if intento < numeroSecreto {
			fmt.Println("📈 El número es MAYOR")
		} else if intento > numeroSecreto {
			fmt.Println("📉 El número es MENOR")
		} else {
			// ¡Adivinó!
			adivinado = true
			fmt.Println("\n🎉 ¡CORRECTO! 🎉")
			fmt.Printf("¡Has adivinado el número %d en %d intentos!\n", numeroSecreto, intentos)

			// Evaluamos el desempeño
			if intentos <= 5 {
				fmt.Println("⭐ ¡Excelente! Lo lograste muy rápido.")
			} else if intentos <= 10 {
				fmt.Println("👍 ¡Buen trabajo!")
			} else {
				fmt.Println("💪 ¡No te rindas, sigue practicando!")
			}
		}
	}
}
