// 10. CONCURRENCIA EN GO
// Goroutines, Channels, Select, Sync

package main

import (
    "fmt"
    "sync"
    "time"
)

// ==========================================
// 1. GOROUTINES BÁSICAS
// ==========================================
func tarea(id int) {
    fmt.Printf("Tarea %d iniciada\n", id)
    time.Sleep(time.Millisecond * 500)
    fmt.Printf("Tarea %d completada\n", id)
}

// ==========================================
// 2. CHANNELS BÁSICOS
// ==========================================
func productor(ch chan<- int) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Produciendo: %d\n", i)
        ch <- i // Enviar al channel
        time.Sleep(time.Millisecond * 200)
    }
    close(ch) // Cerrar el channel cuando termine
}

func consumidor(ch <-chan int, done chan<- bool) {
    for valor := range ch { // Recibir hasta que se cierre
        fmt.Printf("  Consumiendo: %d\n", valor)
        time.Sleep(time.Millisecond * 300)
    }
    done <- true
}

// ==========================================
// 3. BUFFERED CHANNELS
// ==========================================
func ejemploBufferedChannel() {
    ch := make(chan int, 3) // Buffer de tamaño 3
    
    // Enviar sin bloquear hasta llenar el buffer
    ch <- 1
    ch <- 2
    ch <- 3
    
    fmt.Println("Buffer lleno, valores enviados")
    
    // Recibir valores
    fmt.Println("Recibido:", <-ch)
    fmt.Println("Recibido:", <-ch)
    fmt.Println("Recibido:", <-ch)
}

// ==========================================
// 4. SELECT (multiplexing de channels)
// ==========================================
func ejemploSelect() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(time.Millisecond * 500)
        ch1 <- "Mensaje desde ch1"
    }()
    
    go func() {
        time.Sleep(time.Millisecond * 300)
        ch2 <- "Mensaje desde ch2"
    }()
    
    // Select espera el primer channel que esté listo
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        }
    }
}

// ==========================================
// 5. WAITGROUP
// ==========================================
func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrementar contador al terminar
    
    fmt.Printf("Worker %d iniciando\n", id)
    time.Sleep(time.Millisecond * 500)
    fmt.Printf("Worker %d terminado\n", id)
}

// ==========================================
// 6. MUTEX (sincronización)
// ==========================================
type Contador struct {
    mu    sync.Mutex
    valor int
}

func (c *Contador) Incrementar() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.valor++
}

func (c *Contador) Valor() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.valor
}

func main() {
    fmt.Println("=== CONCURRENCIA EN GO ===\n")

    // ==========================================
    // 1. GOROUTINES
    // ==========================================
    fmt.Println("1. Goroutines básicas:")
    for i := 1; i <= 3; i++ {
        go tarea(i) // Lanzar goroutine
    }
    time.Sleep(time.Second) // Esperar a que terminen
    fmt.Println()

    // ==========================================
    // 2. CHANNELS
    // ==========================================
    fmt.Println("2. Channels (productor-consumidor):")
    ch := make(chan int)
    done := make(chan bool)
    
    go productor(ch)
    go consumidor(ch, done)
    
    <-done // Esperar a que el consumidor termine
    fmt.Println()

    // ==========================================
    // 3. BUFFERED CHANNELS
    // ==========================================
    fmt.Println("3. Buffered channels:")
    ejemploBufferedChannel()
    fmt.Println()

    // ==========================================
    // 4. SELECT
    // ==========================================
    fmt.Println("4. Select (multiplexing):")
    ejemploSelect()
    fmt.Println()

    // ==========================================
    // 5. WAITGROUP
    // ==========================================
    fmt.Println("5. WaitGroup:")
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1) // Incrementar contador
        go worker(i, &wg)
    }
    
    wg.Wait() // Esperar a que todos terminen
    fmt.Println("Todos los workers terminaron\n")

    // ==========================================
    // 6. MUTEX
    // ==========================================
    fmt.Println("6. Mutex (protección de datos):")
    contador := &Contador{}
    var wg2 sync.WaitGroup
    
    // 100 goroutines incrementando concurrentemente
    for i := 0; i < 100; i++ {
        wg2.Add(1)
        go func() {
            defer wg2.Done()
            contador.Incrementar()
        }()
    }
    
    wg2.Wait()
    fmt.Printf("Valor final del contador: %d\n\n", contador.Valor())

    // ==========================================
    // 7. SELECT CON TIMEOUT
    // ==========================================
    fmt.Println("7. Select con timeout:")
    ch3 := make(chan string)
    
    go func() {
        time.Sleep(time.Second * 2)
        ch3 <- "Mensaje tardío"
    }()
    
    select {
    case msg := <-ch3:
        fmt.Println("Recibido:", msg)
    case <-time.After(time.Second):
        fmt.Println("Timeout: no se recibió mensaje a tiempo")
    }
    fmt.Println()

    // ==========================================
    // 8. CHANNEL DIRECCIONALES
    // ==========================================
    fmt.Println("8. Channels direccionales:")
    ch4 := make(chan int)
    
    // Solo envío
    go func(ch chan<- int) {
        ch <- 42
    }(ch4)
    
    // Solo recepción
    go func(ch <-chan int) {
        valor := <-ch
        fmt.Printf("Valor recibido: %d\n", valor)
    }(ch4)
    
    time.Sleep(time.Millisecond * 100)
    fmt.Println()

    // ==========================================
    // PATRONES COMUNES
    // ==========================================
    fmt.Println("=== PATRONES COMUNES ===\n")
    fmt.Println(`
✅ Worker Pool:
  - Crear N workers que consumen de un channel
  - Útil para limitar concurrencia

✅ Fan-Out/Fan-In:
  - Distribuir trabajo a múltiples goroutines
  - Recolectar resultados en un channel

✅ Pipeline:
  - Cadena de stages conectados por channels
  - Cada stage procesa y pasa al siguiente

✅ Cancelación con Context:
  - Propagar señales de cancelación
  - Timeout y deadline management
    `)

    fmt.Println("\n💡 Nunca compartas memoria comunicando")
    fmt.Println("💡 En su lugar, comunica compartiendo memoria (channels)")
}