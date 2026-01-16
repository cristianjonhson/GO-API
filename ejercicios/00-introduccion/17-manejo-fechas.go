// Ejemplo básico de manejo de fechas y horas en Go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Fecha y hora actual
	hoy := time.Now()
	fmt.Println("Fecha y hora actual:", hoy)

	// Formatear fecha
	formateada := hoy.Format("02-01-2006 15:04:05")
	fmt.Println("Fecha formateada:", formateada)

	// Sumar días
	mañana := hoy.Add(24 * time.Hour)
	fmt.Println("Mañana será:", mañana)
}