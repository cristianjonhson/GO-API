// Ejemplo básico de punteros en Go
package main

import "fmt"

func main() {
	// Declaración de una variable
	num := 42

	// Crear un puntero a la variable
	ptr := &num

	// Imprimir el valor y la dirección de memoria
	fmt.Println("Valor de num:", num)
	fmt.Println("Dirección de memoria de num:", ptr)

	// Modificar el valor a través del puntero
	*ptr = 100
	fmt.Println("Nuevo valor de num:", num)
}