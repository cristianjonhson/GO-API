// Ejercicio 12.1: Descarga simulada con goroutines
// Simula 5 tareas concurrentes que "duermen" entre 1-3 segundos y reportan cuando terminan.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tarea representa una descarga simulada
type Tarea struct {
	ID       int
	Nombre   string
	Duracion time.Duration
}

// Resultado contiene el resultado de una tarea completada
type Resultado struct {
	TareaID  int
	Nombre   string
	Duracion time.Duration
	Exito    bool
	Mensaje  string
}

// descargarArchivo simula la descarga de un archivo
func descargarArchivo(tarea Tarea, resultados chan<- Resultado) {
	fmt.Printf("📥 [Tarea %d] Iniciando descarga: %s\n", tarea.ID, tarea.Nombre)

	// Simulamos el tiempo de descarga
	time.Sleep(tarea.Duracion)

	// Simulamos un 90% de éxito
	rand.Seed(time.Now().UnixNano())
	exito := rand.Float32() < 0.9

	resultado := Resultado{
		TareaID:  tarea.ID,
		Nombre:   tarea.Nombre,
		Duracion: tarea.Duracion,
		Exito:    exito,
	}

	if exito {
		resultado.Mensaje = "Descarga completada"
		fmt.Printf("✅ [Tarea %d] %s - Completada en %.1fs\n",
			tarea.ID, tarea.Nombre, tarea.Duracion.Seconds())
	} else {
		resultado.Mensaje = "Error de red"
		fmt.Printf("❌ [Tarea %d] %s - Falló\n", tarea.ID, tarea.Nombre)
	}

	// Enviamos el resultado por el channel
	resultados <- resultado
}

func main() {
	fmt.Println("=== SIMULADOR DE DESCARGAS CONCURRENTES ===\n")

	// Creamos las tareas a ejecutar
	tareas := []Tarea{
		{ID: 1, Nombre: "video.mp4", Duracion: time.Second * 2},
		{ID: 2, Nombre: "documento.pdf", Duracion: time.Second * 1},
		{ID: 3, Nombre: "imagen.jpg", Duracion: time.Second * 3},
		{ID: 4, Nombre: "audio.mp3", Duracion: time.Second * 2},
		{ID: 5, Nombre: "archivo.zip", Duracion: time.Second * 3},
	}

	// Creamos un channel para recibir resultados
	resultados := make(chan Resultado, len(tareas))

	// Iniciamos todas las descargas como goroutines
	tiempoInicio := time.Now()
	for _, tarea := range tareas {
		go descargarArchivo(tarea, resultados)
	}

	// Recolectamos todos los resultados
	fmt.Println("\n⏳ Esperando que todas las descargas terminen...\n")

	var completadas []Resultado
	for i := 0; i < len(tareas); i++ {
		resultado := <-resultados
		completadas = append(completadas, resultado)
	}

	tiempoTotal := time.Since(tiempoInicio)

	// Mostramos el resumen
	fmt.Println("\n=== RESUMEN DE DESCARGAS ===")
	exitosas := 0
	fallidas := 0

	for _, r := range completadas {
		if r.Exito {
			exitosas++
			fmt.Printf("✅ %s - %s (%.1fs)\n", r.Nombre, r.Mensaje, r.Duracion.Seconds())
		} else {
			fallidas++
			fmt.Printf("❌ %s - %s\n", r.Nombre, r.Mensaje)
		}
	}

	fmt.Printf("\nTotal: %d tareas\n", len(tareas))
	fmt.Printf("Exitosas: %d\n", exitosas)
	fmt.Printf("Fallidas: %d\n", fallidas)
	fmt.Printf("Tiempo total: %.1fs\n", tiempoTotal.Seconds())

	fmt.Println("\n💡 Nota: Con concurrencia, todas las tareas se ejecutaron en paralelo,")
	fmt.Println("   tomando aproximadamente el tiempo de la tarea más lenta (~3s)")
	fmt.Println("   en lugar de la suma de todas (~11s en secuencial).")

	close(resultados)
}
