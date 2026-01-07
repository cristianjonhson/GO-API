// 12. CONTEXT EN GO
// Context permite propagar señales de cancelación, timeouts y valores entre goroutines

package main

import (
    "context"
    "fmt"
    "time"
)

// ==========================================
// 1. CONTEXT BÁSICO - Background y TODO
// ==========================================
func ejemploContextBasico() {
    fmt.Println("1. Context básico:")
    
    // Background: context raíz, nunca se cancela
    ctx := context.Background()
    fmt.Printf("Context background: %v\n", ctx)
    
    // TODO: placeholder cuando no estás seguro qué context usar
    ctxTodo := context.TODO()
    fmt.Printf("Context TODO: %v\n\n", ctxTodo)
}

// ==========================================
// 2. CONTEXT CON CANCELACIÓN
// ==========================================
func ejemploWithCancel() {
    fmt.Println("2. Context con cancelación:")
    
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // Siempre cancelar para liberar recursos
    
    go func() {
        for {
            select {
            case <-ctx.Done():
                fmt.Println("  Goroutine cancelada:", ctx.Err())
                return
            default:
                fmt.Println("  Trabajando...")
                time.Sleep(time.Millisecond * 300)
            }
        }
    }()
    
    time.Sleep(time.Second)
    fmt.Println("  Cancelando context...")
    cancel()
    time.Sleep(time.Millisecond * 100)
    fmt.Println()
}

// ==========================================
// 3. CONTEXT CON TIMEOUT
// ==========================================
func operacionLenta(ctx context.Context, id int) error {
    select {
    case <-time.After(time.Second * 2): // Simula operación lenta
        fmt.Printf("  Operación %d completada\n", id)
        return nil
    case <-ctx.Done():
        fmt.Printf("  Operación %d cancelada: %v\n", id, ctx.Err())
        return ctx.Err()
    }
}

func ejemploWithTimeout() {
    fmt.Println("3. Context con timeout:")
    
    // Context con timeout de 1 segundo
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    
    fmt.Println("  Iniciando operación con timeout de 1s...")
    err := operacionLenta(ctx, 1)
    if err != nil {
        fmt.Printf("  Error: %v\n", err)
    }
    fmt.Println()
}

// ==========================================
// 4. CONTEXT CON DEADLINE
// ==========================================
func ejemploWithDeadline() {
    fmt.Println("4. Context con deadline:")
    
    // Deadline absoluto (tiempo específico)
    deadline := time.Now().Add(time.Millisecond * 800)
    ctx, cancel := context.WithDeadline(context.Background(), deadline)
    defer cancel()
    
    fmt.Printf("  Deadline: %v\n", deadline.Format("15:04:05.000"))
    
    err := operacionLenta(ctx, 2)
    if err != nil {
        fmt.Printf("  Error: %v\n", err)
    }
    fmt.Println()
}

// ==========================================
// 5. CONTEXT CON VALORES
// ==========================================
type key string

const (
    userIDKey key = "userID"
    requestIDKey key = "requestID"
)

func ejemploWithValue() {
    fmt.Println("5. Context con valores:")
    
    // Crear context con valores
    ctx := context.Background()
    ctx = context.WithValue(ctx, userIDKey, "user123")
    ctx = context.WithValue(ctx, requestIDKey, "req456")
    
    // Recuperar valores
    procesarRequest(ctx)
    fmt.Println()
}

func procesarRequest(ctx context.Context) {
    userID := ctx.Value(userIDKey)
    requestID := ctx.Value(requestIDKey)
    
    fmt.Printf("  UserID: %v\n", userID)
    fmt.Printf("  RequestID: %v\n", requestID)
    
    // Verificar si el valor existe
    if uid, ok := ctx.Value(userIDKey).(string); ok {
        fmt.Printf("  Procesando request para usuario: %s\n", uid)
    }
}

// ==========================================
// 6. PROPAGACIÓN DE CONTEXT
// ==========================================
func ejemploPropagacion() {
    fmt.Println("6. Propagación de context:")
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
    defer cancel()
    
    // Agregar valores al context
    ctx = context.WithValue(ctx, userIDKey, "user789")
    
    // Propagar a funciones hijas
    nivel1(ctx)
    fmt.Println()
}

func nivel1(ctx context.Context) {
    fmt.Println("  Nivel 1:")
    fmt.Printf("    UserID: %v\n", ctx.Value(userIDKey))
    nivel2(ctx)
}

func nivel2(ctx context.Context) {
    fmt.Println("  Nivel 2:")
    fmt.Printf("    UserID: %v\n", ctx.Value(userIDKey))
    
    // Verificar si el context fue cancelado
    select {
    case <-ctx.Done():
        fmt.Printf("    Context cancelado: %v\n", ctx.Err())
    default:
        fmt.Println("    Context activo")
    }
}

// ==========================================
// 7. WORKER POOL CON CONTEXT
// ==========================================
func ejemploWorkerPool() {
    fmt.Println("7. Worker Pool con context:")
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
    defer cancel()
    
    jobs := make(chan int, 10)
    results := make(chan int, 10)
    
    // Crear 3 workers
    for w := 1; w <= 3; w++ {
        go worker(ctx, w, jobs, results)
    }
    
    // Enviar trabajos
    go func() {
        for j := 1; j <= 5; j++ {
            select {
            case jobs <- j:
            case <-ctx.Done():
                close(jobs)
                return
            }
        }
        close(jobs)
    }()
    
    // Recolectar resultados
    for i := 1; i <= 5; i++ {
        select {
        case result := <-results:
            fmt.Printf("  Resultado: %d\n", result)
        case <-ctx.Done():
            fmt.Printf("  Timeout alcanzado: %v\n", ctx.Err())
            return
        }
    }
    fmt.Println()
}

func worker(ctx context.Context, id int, jobs <-chan int, results chan<- int) {
    for {
        select {
        case job, ok := <-jobs:
            if !ok {
                return
            }
            fmt.Printf("  Worker %d procesando job %d\n", id, job)
            time.Sleep(time.Millisecond * 500)
            
            select {
            case results <- job * 2:
            case <-ctx.Done():
                fmt.Printf("  Worker %d cancelado\n", id)
                return
            }
        case <-ctx.Done():
            fmt.Printf("  Worker %d detenido: %v\n", id, ctx.Err())
            return
        }
    }
}

// ==========================================
// 8. HTTP REQUEST CON CONTEXT (SIMULADO)
// ==========================================
func ejemploHTTPRequest() {
    fmt.Println("8. Simulación de HTTP request con context:")
    
    ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
    defer cancel()
    
    err := fetchData(ctx, "https://api.example.com/data")
    if err != nil {
        fmt.Printf("  Error: %v\n", err)
    }
    fmt.Println()
}

func fetchData(ctx context.Context, url string) error {
    fmt.Printf("  Iniciando request a: %s\n", url)
    
    // Simular request HTTP
    select {
    case <-time.After(time.Second): // Simula respuesta lenta
        fmt.Println("  Datos recibidos")
        return nil
    case <-ctx.Done():
        return fmt.Errorf("request cancelado: %w", ctx.Err())
    }
}

// ==========================================
// 9. CANCELACIÓN EN CADENA
// ==========================================
func ejemploCancelacionCadena() {
    fmt.Println("9. Cancelación en cadena:")
    
    // Context padre
    ctxPadre, cancelPadre := context.WithCancel(context.Background())
    defer cancelPadre()
    
    // Context hijo (se cancela automáticamente cuando el padre se cancela)
    ctxHijo, cancelHijo := context.WithTimeout(ctxPadre, time.Second*5)
    defer cancelHijo()
    
    go func() {
        <-ctxHijo.Done()
        fmt.Printf("  Context hijo cancelado: %v\n", ctxHijo.Err())
    }()
    
    time.Sleep(time.Millisecond * 500)
    fmt.Println("  Cancelando context padre...")
    cancelPadre()
    time.Sleep(time.Millisecond * 100)
    fmt.Println()
}

// ==========================================
// 10. MEJORES PRÁCTICAS
// ==========================================
func ejemploMejoresPracticas() {
    fmt.Println("10. Mejores prácticas:\n")
    
    fmt.Println(`
✅ DO:
  • Pasar context como primer parámetro de funciones
  • Siempre llamar a cancel() con defer
  • Usar context.Background() como raíz
  • Propagar context a través de la cadena de llamadas
  • Verificar ctx.Done() en operaciones largas
  • Usar context para timeouts y cancelaciones
  • Context es inmutable (crear nuevos para agregar valores)

❌ DON'T:
  • No guardar context en structs (pasarlo como parámetro)
  • No usar context.Value() para datos críticos
  • No ignorar el error de ctx.Err()
  • No crear context sin cancelación y olvidarlo
  • No usar valores nil como keys en context.Value
  • No pasar nil como context (usar context.Background())
    `)
}

// ==========================================
// MAIN
// ==========================================
func main() {
    fmt.Println("=== CONTEXT EN GO ===\n")
    
    ejemploContextBasico()
    ejemploWithCancel()
    ejemploWithTimeout()
    ejemploWithDeadline()
    ejemploWithValue()
    ejemploPropagacion()
    ejemploWorkerPool()
    ejemploHTTPRequest()
    ejemploCancelacionCadena()
    ejemploMejoresPracticas()
    
    fmt.Println("\n💡 Context es fundamental para:")
    fmt.Println("   • Cancelación de operaciones")
    fmt.Println("   • Timeouts y deadlines")
    fmt.Println("   • Propagación de valores request-scoped")
    fmt.Println("   • Coordinación de goroutines")
}