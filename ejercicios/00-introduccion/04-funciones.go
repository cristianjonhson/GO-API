// 04. FUNCIONES EN GO
// Funciones son bloques de código reutilizables

package main

import "fmt"

// ==========================================
// 1. FUNCIÓN BÁSICA
// ==========================================
func saludar() {
    fmt.Println("¡Hola!")
}

// ==========================================
// 2. FUNCIÓN CON PARÁMETROS
// ==========================================
func saludarPersona(nombre string) {
    fmt.Printf("¡Hola, %s!\n", nombre)
}

// Múltiples parámetros del mismo tipo
func sumar(a, b int) {
    fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

// ==========================================
// 3. FUNCIÓN CON RETORNO
// ==========================================
func multiplicar(a, b int) int {
    return a * b
}

// ==========================================
// 4. MÚLTIPLES VALORES DE RETORNO
// ==========================================
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("división por cero")
    }
    return a / b, nil
}

// ==========================================
// 5. RETORNOS NOMBRADOS
// ==========================================
func rectangulo(base, altura float64) (area, perimetro float64) {
    area = base * altura
    perimetro = 2 * (base + altura)
    return // retorno implícito de las variables nombradas
}

// ==========================================
// 6. FUNCIONES VARIÁDICAS (número variable de argumentos)
// ==========================================
func sumarTodos(numeros ...int) int {
    total := 0
    for _, num := range numeros {
        total += num
    }
    return total
}

// ==========================================
// 7. FUNCIONES COMO VALORES (first-class functions)
// ==========================================
func aplicarOperacion(a, b int, operacion func(int, int) int) int {
    return operacion(a, b)
}

// ==========================================
// 8. FUNCIONES ANÓNIMAS (closures)
// ==========================================
func crearContador() func() int {
    contador := 0
    return func() int {
        contador++
        return contador
    }
}

// ==========================================
// 9. DEFER (ejecución diferida)
// ==========================================
func ejemploDefer() {
    defer fmt.Println("3. Esto se ejecuta al final")
    fmt.Println("1. Primera línea")
    fmt.Println("2. Segunda línea")
}

// ==========================================
// 10. MÉTODOS RECEIVERS (funciones asociadas a tipos)
// ==========================================
type Persona struct {
    Nombre string
    Edad   int
}

// Receiver por valor (no modifica el original)
func (p Persona) Saludar() {
    fmt.Printf("Hola, soy %s y tengo %d años\n", p.Nombre, p.Edad)
}

// Receiver por puntero (puede modificar el original)
func (p *Persona) CumplirAnios() {
    p.Edad++
    fmt.Printf("%s ahora tiene %d años\n", p.Nombre, p.Edad)
}

func main() {
    fmt.Println("=== FUNCIONES EN GO ===\n")

    // 1. Función básica
    saludar()

    // 2. Función con parámetros
    saludarPersona("Ana")
    sumar(5, 3)

    // 3. Función con retorno
    resultado := multiplicar(4, 7)
    fmt.Printf("4 × 7 = %d\n", resultado)

    // 4. Múltiples retornos
    cociente, err := dividir(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("10 ÷ 2 = %.2f\n", cociente)
    }

    _, err2 := dividir(10, 0) // Ignorar primer retorno con _
    if err2 != nil {
        fmt.Println("Error esperado:", err2)
    }

    // 5. Retornos nombrados
    a, p := rectangulo(5, 3)
    fmt.Printf("Rectángulo: área=%.0f, perímetro=%.0f\n", a, p)

    // 6. Funciones variádicas
    suma1 := sumarTodos(1, 2, 3)
    suma2 := sumarTodos(10, 20, 30, 40, 50)
    fmt.Printf("Suma 1: %d, Suma 2: %d\n", suma1, suma2)

    // Expandir slice como argumentos variádicos
    numeros := []int{1, 2, 3, 4, 5}
    suma3 := sumarTodos(numeros...) // ... expande el slice
    fmt.Printf("Suma 3: %d\n", suma3)

    // 7. Funciones como valores
    suma := func(a, b int) int { return a + b }
    resta := func(a, b int) int { return a - b }
    
    fmt.Printf("5 + 3 = %d\n", aplicarOperacion(5, 3, suma))
    fmt.Printf("5 - 3 = %d\n", aplicarOperacion(5, 3, resta))

    // 8. Closures
    contador := crearContador()
    fmt.Println("Contador:", contador()) // 1
    fmt.Println("Contador:", contador()) // 2
    fmt.Println("Contador:", contador()) // 3

    // 9. Defer
    fmt.Println("\nEjemplo defer:")
    ejemploDefer()

    // Múltiples defers (se ejecutan en orden LIFO)
    fmt.Println("\nMúltiples defers:")
    defer fmt.Println("Tercero")
    defer fmt.Println("Segundo")
    defer fmt.Println("Primero")
    fmt.Println("Inicio")

    // 10. Métodos
    fmt.Println("\nMétodos:")
    persona := Persona{Nombre: "Carlos", Edad: 30}
    persona.Saludar()
    persona.CumplirAnios()
    persona.Saludar()
}