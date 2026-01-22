// Este programa muestra cómo usar goroutines con un WaitGroup para sincronizar su ejecución.
// Un WaitGroup permite esperar a que un grupo de goroutines termine antes de continuar.

package main

import (
	"fmt"
	"sync"
)

func tarea(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Tarea %d iniciada\n", id)
	// Simular trabajo
}

func main() {
	var wg sync.WaitGroup

	// Lanzar varias goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go tarea(i, &wg)
	}

	// Esperar a que todas terminen
	wg.Wait()
	fmt.Println("Todas las tareas completadas")
}