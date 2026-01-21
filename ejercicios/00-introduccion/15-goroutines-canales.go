// Ejemplo básico de goroutines y canales en Go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Crear un canal
	ch := make(chan string)

	// Iniciar una goroutine
	go func() {
		ch <- "Hola desde la goroutine"
	}()

	// Leer del canal
	mensaje := <-ch
	fmt.Println(mensaje)

	// Ejemplo con múltiples goroutines
	for i := 1; i <= 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d ejecutándose\n", id)
			time.Sleep(1 * time.Second)
		}(i)
	}

	// Esperar para que las goroutines terminen
	time.Sleep(2 * time.Second)
	fmt.Println("Fin del programa")
}