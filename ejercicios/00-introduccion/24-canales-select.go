// Ejemplo básico de uso de canales con select
package main

import (
	"fmt"
	"time"
)

func enviarMensajes(canal chan string, mensaje string, delay time.Duration) {
	for {
		time.Sleep(delay)
		canal <- mensaje
	}
}

func main() {
	canal1 := make(chan string)
	canal2 := make(chan string)

	// Lanzar goroutines
	go enviarMensajes(canal1, "Mensaje del canal 1", 2*time.Second)
	go enviarMensajes(canal2, "Mensaje del canal 2", 3*time.Second)

	// Usar select para manejar múltiples canales
	for i := 0; i < 5; i++ {
		select {
		case msg1 := <-canal1:
			fmt.Println(msg1)
		case msg2 := <-canal2:
			fmt.Println(msg2)
		}
	}
}