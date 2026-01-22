// Este programa muestra cómo usar un paquete externo en Go.
// En este caso, usamos el paquete "github.com/google/uuid" para generar un UUID.
// Los paquetes externos se instalan con "go get" y se importan como cualquier otro paquete.

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