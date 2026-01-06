// Ejercicio 1.1: Hola, perfil
// Pide al usuario su nombre, edad y ciudad, luego imprime un resumen formateado.

package main

import "fmt"

func main() {
	// Variables para almacenar los datos del usuario
	var nombre, ciudad string
	var edad int

	// Pedimos el nombre
	fmt.Print("Ingresa tu nombre: ")
	fmt.Scanln(&nombre)

	// Pedimos la edad
	fmt.Print("Ingresa tu edad: ")
	fmt.Scanln(&edad)

	// Pedimos la ciudad
	fmt.Print("Ingresa tu ciudad: ")
	fmt.Scanln(&ciudad)

	// Imprimimos el resumen formateado
	fmt.Println("\n--- RESUMEN DE PERFIL ---")
	fmt.Printf("Nombre: %s\n", nombre)
	fmt.Printf("Edad: %d años\n", edad)
	fmt.Printf("Ciudad: %s\n", ciudad)
	fmt.Println("-------------------------")
}
