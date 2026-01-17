// Ejemplo básico de defer, panic y recover en Go
package main

import "fmt"

func main() {
	// Manejo de errores con recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de un pánico:", r)
		}
	}()

	fmt.Println("Inicio del programa")

	// Generar un pánico
	panic("Algo salió mal")

	fmt.Println("Fin del programa") // No se ejecutará
}