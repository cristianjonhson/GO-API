// Ejercicio 12.2: Procesador de números con pipeline concurrente
// Implementa el patrón producer-workers-consumer usando channels.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Resultado representa el resultado procesado por un worker
type Resultado struct {
	NumeroOriginal int
	Cuadrado       int
	Cubo           int
	WorkerID       int
}

// producer genera números del 1 al N y los envía por el channel
func producer(n int, numeros chan<- int) {
	fmt.Printf("🏭 Producer: Generando números del 1 al %d...\n", n)
	for i := 1; i <= n; i++ {
		numeros <- i
		time.Sleep(time.Millisecond * 100) // Simulamos trabajo
	}
	close(numeros)
	fmt.Println("🏭 Producer: Todos los números generados")
}

// worker procesa números calculando su cuadrado y cubo
func worker(id int, numeros <-chan int, resultados chan<- Resultado, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("👷 Worker %d: Iniciado\n", id)

	for numero := range numeros {
		// Simulamos procesamiento
		time.Sleep(time.Millisecond * 200)

		cuadrado := numero * numero
		cubo := numero * numero * numero

		resultado := Resultado{
			NumeroOriginal: numero,
			Cuadrado:       cuadrado,
			Cubo:           cubo,
			WorkerID:       id,
		}

		resultados <- resultado
		fmt.Printf("👷 Worker %d: Procesó %d -> cuadrado=%d, cubo=%d\n",
			id, numero, cuadrado, cubo)
	}

	fmt.Printf("👷 Worker %d: Finalizado\n", id)
}

// consumer recibe y muestra los resultados procesados
func consumer(resultados <-chan Resultado, done chan<- bool) {
	fmt.Println("📊 Consumer: Esperando resultados...\n")

	contador := 0
	for resultado := range resultados {
		contador++
		fmt.Printf("📊 Resultado #%d: %d² = %d, %d³ = %d (worker %d)\n",
			contador,
			resultado.NumeroOriginal,
			resultado.Cuadrado,
			resultado.NumeroOriginal,
			resultado.Cubo,
			resultado.WorkerID)
	}

	fmt.Printf("\n📊 Consumer: Procesados %d resultados en total\n", contador)
	done <- true
}

func main() {
	fmt.Println("=== PROCESADOR CONCURRENTE DE NÚMEROS ===\n")

	// Configuración
	const (
		totalNumeros = 10  // Números a procesar
		numWorkers   = 3   // Cantidad de workers concurrentes
	)

	// Creamos los channels
	numeros := make(chan int, 5)      // Buffer para números
	resultados := make(chan Resultado, 5) // Buffer para resultados
	done := make(chan bool)

	// WaitGroup para sincronizar workers
	var wg sync.WaitGroup

	fmt.Printf("⚙️  Configuración:\n")
	fmt.Printf("   - Números a procesar: %d\n", totalNumeros)
	fmt.Printf("   - Workers concurrentes: %d\n\n", numWorkers)

	// Iniciamos el consumer
	go consumer(resultados, done)

	// Iniciamos los workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, numeros, resultados, &wg)
	}

	// Iniciamos el producer
	tiempoInicio := time.Now()
	go producer(totalNumeros, numeros)

	// Esperamos a que todos los workers terminen
	wg.Wait()
	close(resultados)

	// Esperamos a que el consumer termine
	<-done

	tiempoTotal := time.Since(tiempoInicio)

	// Resumen final
	fmt.Println("\n=== RESUMEN ===")
	fmt.Printf("⏱️  Tiempo total de procesamiento: %.2fs\n", tiempoTotal.Seconds())
	fmt.Printf("🚀 Con %d workers concurrentes, el procesamiento fue más rápido\n", numWorkers)
	fmt.Println("   que si se hubiera hecho secuencialmente.")

	// Cálculo teórico
	tiempoSecuencial := float64(totalNumeros) * 0.2
	fmt.Printf("\n💡 Tiempo secuencial estimado: %.1fs\n", tiempoSecuencial)
	fmt.Printf("   Mejora con concurrencia: ~%.0f%%\n",
		(1-tiempoTotal.Seconds()/tiempoSecuencial)*100)
}
