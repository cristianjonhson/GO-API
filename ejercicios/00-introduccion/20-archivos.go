// Ejemplo básico de lectura y escritura de archivos en Go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Crear un archivo
	archivo, err := os.Create("ejemplo.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer archivo.Close()

	// Escribir en el archivo
	_, err = archivo.WriteString("Hola, archivo\n")
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}
	fmt.Println("Archivo creado y escrito con éxito")

	// Leer el archivo
	contenido, err := os.ReadFile("ejemplo.txt")
	if err != nil {
		fmt.Println("Error al leer el archivo:", err)
		return
	}
	fmt.Println("Contenido del archivo:", string(contenido))
}