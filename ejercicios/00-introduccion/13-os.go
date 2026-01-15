// Ejemplo básico del uso del paquete os en Go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Obtener el valor de una variable de entorno
	usuario := os.Getenv("USER")
	fmt.Printf("Usuario actual: %s\n", usuario)

	// Listar los archivos en el directorio actual
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error obteniendo el directorio actual: %v\n", err)
		return
	}

	archivos, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error leyendo el directorio: %v\n", err)
		return
	}

	fmt.Println("Archivos en el directorio actual:")
	for _, archivo := range archivos {
		fmt.Println("-", archivo.Name())
	}

	// Crear un archivo temporal
	archivoTemp, err := os.CreateTemp("", "ejemplo-*.txt")
	if err != nil {
		fmt.Printf("Error creando un archivo temporal: %v\n", err)
		return
	}
	defer os.Remove(archivoTemp.Name()) // Eliminar el archivo temporal al salir

	fmt.Printf("Archivo temporal creado: %s\n", archivoTemp.Name())
}