// Ejemplo básico de lectura de entrada del usuario en Go
package main

import "fmt"

func main() {
	var nombre string
	fmt.Print("Introduce tu nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println("Hola,", nombre)
}