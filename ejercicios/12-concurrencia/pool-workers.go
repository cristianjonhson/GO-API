// Ejercicio 12.3: Ejemplo adicional - Pool de workers con tareas reales
// Demuestra un pool de workers procesando múltiples tareas de forma concurrente.

package main

import (
	"fmt"
	"sync"
	"time"
)

// Job representa una tarea a ejecutar
type Job struct {
	ID          int
	Descripcion string
	Duracion    time.Duration
}

// JobResult representa el resultado de una tarea
type JobResult struct {
	JobID       int
	Descripcion string
	WorkerID    int
	Duracion    time.Duration
	Completado  time.Time
}

// Worker procesa jobs de la cola
func Worker(id int, jobs <-chan Job, results chan<- JobResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("⚙️  Worker %d: Iniciando job %d - %s\n", id, job.ID, job.Descripcion)

		// Simulamos el trabajo
		inicio := time.Now()
		time.Sleep(job.Duracion)
		duracionReal := time.Since(inicio)

		// Enviamos el resultado
		result := JobResult{
			JobID:       job.ID,
			Descripcion: job.Descripcion,
			WorkerID:    id,
			Duracion:    duracionReal,
			Completado:  time.Now(),
		}

		results <- result
		fmt.Printf("✅ Worker %d: Completó job %d en %.2fs\n",
			id, job.ID, duracionReal.Seconds())
	}

	fmt.Printf("🛑 Worker %d: Sin más trabajos, finalizando\n", id)
}

func main() {
	fmt.Println("=== POOL DE WORKERS CONCURRENTES ===\n")

	// Configuración
	const numWorkers = 3
	const numJobs = 10

	// Creamos los channels
	jobs := make(chan Job, numJobs)
	results := make(chan JobResult, numJobs)

	// WaitGroup para los workers
	var wg sync.WaitGroup

	// Iniciamos el pool de workers
	fmt.Printf("🚀 Iniciando %d workers...\n\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go Worker(w, jobs, results, &wg)
	}

	// Enviamos los jobs
	fmt.Printf("📋 Enviando %d jobs a la cola...\n\n", numJobs)
	tiempoInicio := time.Now()

	for j := 1; j <= numJobs; j++ {
		// Variamos la duración de cada job
		duracion := time.Millisecond * time.Duration(500+j*100)

		job := Job{
			ID:          j,
			Descripcion: fmt.Sprintf("Tarea #%d", j),
			Duracion:    duracion,
		}
		jobs <- job
	}
	close(jobs)

	// Goroutine para cerrar el channel de resultados cuando todos los workers terminen
	go func() {
		wg.Wait()
		close(results)
	}()

	// Recolectamos los resultados
	var resultadosCompletos []JobResult
	for result := range results {
		resultadosCompletos = append(resultadosCompletos, result)
	}

	tiempoTotal := time.Since(tiempoInicio)

	// Mostramos el resumen
	fmt.Println("\n" + "=".repeat(50))
	fmt.Println("=== RESUMEN DE PROCESAMIENTO ===")
	fmt.Println("=".repeat(50))

	// Agrupamos por worker
	porWorker := make(map[int][]JobResult)
	for _, r := range resultadosCompletos {
		porWorker[r.WorkerID] = append(porWorker[r.WorkerID], r)
	}

	// Estadísticas por worker
	for workerID := 1; workerID <= numWorkers; workerID++ {
		jobs := porWorker[workerID]
		fmt.Printf("\n👷 Worker %d: Completó %d jobs\n", workerID, len(jobs))
		for _, job := range jobs {
			fmt.Printf("   - Job %d: %s (%.2fs)\n",
				job.JobID, job.Descripcion, job.Duracion.Seconds())
		}
	}

	// Estadísticas generales
	fmt.Println("\n📊 ESTADÍSTICAS GENERALES")
	fmt.Printf("   Total de jobs: %d\n", len(resultadosCompletos))
	fmt.Printf("   Workers utilizados: %d\n", numWorkers)
	fmt.Printf("   Tiempo total: %.2fs\n", tiempoTotal.Seconds())

	// Calculamos tiempo secuencial teórico
	var tiempoSecuencial float64
	for j := 1; j <= numJobs; j++ {
		tiempoSecuencial += float64(500+j*100) / 1000.0
	}

	mejora := (1 - tiempoTotal.Seconds()/tiempoSecuencial) * 100

	fmt.Printf("\n💡 COMPARACIÓN")
	fmt.Printf("\n   Tiempo secuencial estimado: %.2fs\n", tiempoSecuencial)
	fmt.Printf("   Mejora con concurrencia: %.0f%%\n", mejora)
	fmt.Printf("   Aceleración (speedup): %.2fx\n", tiempoSecuencial/tiempoTotal.Seconds())

	fmt.Println("\n" + "=".repeat(50))
}

// Función auxiliar para repetir strings (no existe nativamente en Go)
func (s string) repeat(count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
