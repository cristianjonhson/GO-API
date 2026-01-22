// Ejemplo básico de uso de paquetes externos en Go
package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	// Generar un UUID
	id := uuid.New()
	fmt.Println("UUID generado:", id)
}