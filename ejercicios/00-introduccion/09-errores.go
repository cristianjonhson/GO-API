// 09. MANEJO DE ERRORES
// Go usa errores explícitos en lugar de excepciones

package main

import (
    "errors"
    "fmt"
    "os"
    "strconv"
)

// ==========================================
// 1. ERROR BÁSICO
// ==========================================
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}

// ==========================================
// 2. ERROR FORMATEADO
// ==========================================
func validarEdad(edad int) error {
    if edad < 0 {
        return fmt.Errorf("edad inválida: %d (debe ser positiva)", edad)
    }
    if edad > 150 {
        return fmt.Errorf("edad inválida: %d (demasiado alta)", edad)
    }
    return nil
}

// ==========================================
// 3. ERROR PERSONALIZADO
// ==========================================
type ErrorValidacion struct {
    Campo   string
    Valor   interface{}
    Mensaje string
}

func (e ErrorValidacion) Error() string {
    return fmt.Sprintf("error en campo '%s': %s (valor: %v)", 
        e.Campo, e.Mensaje, e.Valor)
}

func validarUsuario(nombre string, edad int) error {
    if nombre == "" {
        return ErrorValidacion{
            Campo:   "nombre",
            Valor:   nombre,
            Mensaje: "no puede estar vacío",
        }
    }
    if edad < 18 {
        return ErrorValidacion{
            Campo:   "edad",
            Valor:   edad,
            Mensaje: "debe ser mayor de 18",
        }
    }
    return nil
}

// ==========================================
// 4. MÚLTIPLES ERRORES
// ==========================================
func procesarDatos(datos []string) ([]int, error) {
    resultado := make([]int, 0, len(datos))
    
    for i, dato := range datos {
        num, err := strconv.Atoi(dato)
        if err != nil {
            return nil, fmt.Errorf("error en posición %d: %w", i, err)
        }
        resultado = append(resultado, num)
    }
    
    return resultado, nil
}

// ==========================================
// 5. PANIC Y RECOVER
// ==========================================
func operacionArriesgada() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado del panic:", r)
        }
    }()
    
    fmt.Println("Antes del panic")
    panic("¡Algo salió muy mal!")
    fmt.Println("Después del panic (no se ejecuta)")
}

// ==========================================
// 6. ERROR WRAPPING (Go 1.13+)
// ==========================================
func leerArchivo(ruta string) error {
    _, err := os.ReadFile(ruta)
    if err != nil {
        return fmt.Errorf("no se pudo leer el archivo: %w", err)
    }
    return nil
}

func main() {
    fmt.Println("=== MANEJO DE ERRORES EN GO ===\n")

    // ==========================================
    // PATRÓN BÁSICO: if err != nil
    // ==========================================
    fmt.Println("1. Patrón básico:")
    resultado, err := dividir(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("10 / 2 = %.2f\n", resultado)

    _, err = dividir(10, 0)
    if err != nil {
        fmt.Println("Error esperado:", err)
    }
    fmt.Println()

    // ==========================================
    // ERROR FORMATEADO
    // ==========================================
    fmt.Println("2. Error formateado:")
    if err := validarEdad(-5); err != nil {
        fmt.Println(err)
    }
    if err := validarEdad(200); err != nil {
        fmt.Println(err)
    }
    if err := validarEdad(25); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("Edad válida: 25")
    }
    fmt.Println()

    // ==========================================
    // ERROR PERSONALIZADO
    // ==========================================
    fmt.Println("3. Error personalizado:")
    if err := validarUsuario("", 25); err != nil {
        fmt.Println(err)
        // Type assertion para acceder a campos personalizados
        if errVal, ok := err.(ErrorValidacion); ok {
            fmt.Printf("  Campo: %s, Valor: %v\n", errVal.Campo, errVal.Valor)
        }
    }
    
    if err := validarUsuario("Ana", 16); err != nil {
        fmt.Println(err)
    }
    fmt.Println()

    // ==========================================
    // MÚLTIPLES ERRORES
    // ==========================================
    fmt.Println("4. Procesamiento con errores:")
    datos := []string{"10", "20", "30"}
    numeros, err := procesarDatos(datos)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Datos procesados:", numeros)
    }

    datosInvalidos := []string{"10", "abc", "30"}
    _, err = procesarDatos(datosInvalidos)
    if err != nil {
        fmt.Println("Error esperado:", err)
    }
    fmt.Println()

    // ==========================================
    // PANIC Y RECOVER
    // ==========================================
    fmt.Println("5. Panic y Recover:")
    operacionArriesgada()
    fmt.Println("El programa continúa después del recover\n")

    // ==========================================
    // ERROR WRAPPING
    // ==========================================
    fmt.Println("6. Error wrapping:")
    err = leerArchivo("archivo_inexistente.txt")
    if err != nil {
        fmt.Println("Error:", err)
        
        // Verificar si contiene un error específico
        if errors.Is(err, os.ErrNotExist) {
            fmt.Println("  El archivo no existe")
        }
    }
    fmt.Println()

    // ==========================================
    // MEJORES PRÁCTICAS
    // ==========================================
    fmt.Println("=== MEJORES PRÁCTICAS ===\n")
    fmt.Println(`
✅ DO:
  • Retornar errores explícitamente
  • Verificar errores inmediatamente (if err != nil)
  • Usar errores nombrados descriptivos
  • Documentar qué errores puede retornar una función
  • Usar fmt.Errorf con %w para wrapping
  • Usar defer + recover solo para casos excepcionales

❌ DON'T:
  • Ignorar errores (_)
  • Usar panic para control de flujo normal
  • Crear errores genéricos poco descriptivos
  • Devolver error y valor válido simultáneamente
    `)

    // ==========================================
    // PATRONES COMUNES
    // ==========================================
    fmt.Println("\n=== PATRONES COMUNES ===\n")
    
    // Patrón 1: Early return
    fmt.Println("Patrón 1: Early return")
    if err := validarDatos(); err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Datos válidos, continuando...\n")

    // Patrón 2: Error handling con defer
    fmt.Println("Patrón 2: Defer para cleanup:")
    if err := operacionConRecursos(); err != nil {
        fmt.Println("Error:", err)
    }
}

func validarDatos() error {
    // Simulación de validación
    return nil
}

func operacionConRecursos() error {
    // Simulación de recurso
    fmt.Println("  Abriendo recurso...")
    defer fmt.Println("  Cerrando recurso (defer)")
    
    // Operaciones...
    fmt.Println("  Procesando...")
    
    return nil
}