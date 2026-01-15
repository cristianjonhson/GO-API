// 08. ESTRUCTURAS DE CONTROL DE FLUJO
// if, else, switch, for, loops

package main

import "fmt"

func main() {
    fmt.Println("=== CONTROL DE FLUJO EN GO ===\n")
    
    demostrarIfElse()
    demostrarSwitch()
    demostrarForLoops()
    demostrarForRange()
    demostrarGoto()
    demostrarDefer()
    demostrarBreakContinueLabels()
    demostrarPatronesComunes()
}

    // ==========================================
    // 1. IF / ELSE
    // ==========================================
func demostrarIfElse() {
    fmt.Println("1. IF / ELSE:")
    edad := 20

    if edad >= 18 {
        fmt.Println("Eres mayor de edad")
    }

    // if con inicialización
    if edad := 15; edad >= 18 {
        fmt.Println("Mayor de edad")
    } else {
        fmt.Println("Menor de edad")
    }

    // if-else anidado
    nota := 85
    if nota >= 90 {
        fmt.Println("Excelente")
    } else if nota >= 70 {
        fmt.Println("Aprobado")
    } else {
        fmt.Println("Reprobado")
    }
    fmt.Println()
}

    // ==========================================
    // 2. SWITCH
    // ==========================================
func demostrarSwitch() {
    fmt.Println("2. SWITCH:")
    
    // Switch básico
    dia := 3
    switch dia {
    case 1:
        fmt.Println("Lunes")
    case 2:
        fmt.Println("Martes")
    case 3:
        fmt.Println("Miércoles")
    default:
        fmt.Println("Otro día")
    }

    // Switch con múltiples casos
    letra := "a"
    switch letra {
    case "a", "e", "i", "o", "u":
        fmt.Println("Es una vocal")
    default:
        fmt.Println("Es una consonante")
    }

    // Switch con condiciones
    numero := 15
    switch {
    case numero < 0:
        fmt.Println("Negativo")
    case numero == 0:
        fmt.Println("Cero")
    case numero > 0:
        fmt.Println("Positivo")
    }

    // Switch con inicialización
    switch resultado := 10 * 5; {
    case resultado > 100:
        fmt.Println("Mayor a 100")
    case resultado < 100:
        fmt.Println("Menor a 100")
    default:
        fmt.Println("Igual a 100")
    }

    // Switch de tipos (type switch)
    var valor interface{} = "Hola"
    switch v := valor.(type) {
    case int:
        fmt.Printf("Es un int: %d\n", v)
    case string:
        fmt.Printf("Es un string: %s\n", v)
    default:
        fmt.Printf("Tipo desconocido: %T\n", v)
    }
    fmt.Println()
}

    // ==========================================
    // 3. FOR (único loop en Go)
    // ==========================================
func demostrarForLoops() {
    fmt.Println("3. FOR LOOPS:")

    // For clásico (estilo C)
    fmt.Print("For clásico: ")
    for i := 0; i < 5; i++ {
        fmt.Print(i, " ")
    }
    fmt.Println()

    // For como while
    fmt.Print("For como while: ")
    contador := 0
    for contador < 5 {
        fmt.Print(contador, " ")
        contador++
    }
    fmt.Println()

    // For infinito (equivalente a while true)
    fmt.Print("For infinito (con break): ")
    n := 0
    for {
        if n >= 5 {
            break
        }
        fmt.Print(n, " ")
        n++
    }
    fmt.Println()

    // For con continue
    fmt.Print("For con continue: ")
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            continue // Salta números pares
        }
        fmt.Print(i, " ")
    }
    fmt.Println("\n")
}

    // ==========================================
    // 4. FOR RANGE (iteración sobre colecciones)
    // ==========================================
func demostrarForRange() {
    fmt.Println("4. FOR RANGE:")

    // Range sobre slice
    numeros := []int{10, 20, 30, 40, 50}
    fmt.Print("Range sobre slice: ")
    for indice, valor := range numeros {
        fmt.Printf("[%d]=%d ", indice, valor)
    }
    fmt.Println()

    // Range solo con valores (ignorar índice)
    fmt.Print("Solo valores: ")
    for _, valor := range numeros {
        fmt.Print(valor, " ")
    }
    fmt.Println()

    // Range solo con índices
    fmt.Print("Solo índices: ")
    for indice := range numeros {
        fmt.Print(indice, " ")
    }
    fmt.Println()

    // Range sobre map
    edades := map[string]int{
        "Ana":    25,
        "Carlos": 30,
        "Luis":   28,
    }
    fmt.Println("Range sobre map:")
    for nombre, edad := range edades {
        fmt.Printf("  %s: %d años\n", nombre, edad)
    }

    // Range sobre string (itera sobre runas)
    texto := "Hola"
    fmt.Print("Range sobre string: ")
    for i, runa := range texto {
        fmt.Printf("[%d]=%c ", i, runa)
    }
    fmt.Println("\n")
}

    // ==========================================
    // 5. GOTO (no recomendado, evitar)
    // ==========================================
func demostrarGoto() {
    fmt.Println("5. GOTO (uso excepcional):")
    x := 0
loop:
    if x < 3 {
        fmt.Printf("x = %d\n", x)
        x++
        goto loop
    }
    fmt.Println()
}

    // ==========================================
    // 6. DEFER
    // ==========================================
func demostrarDefer() {
    fmt.Println("6. DEFER:")
    defer fmt.Println("Esto se ejecuta al final (defer 3)")
    defer fmt.Println("Esto se ejecuta antes (defer 2)")
    defer fmt.Println("Esto se ejecuta primero (defer 1)")
    fmt.Println("Código normal")
    fmt.Println()
}

func demostrarBreakContinueLabels() {
    fmt.Println("7. BREAK/CONTINUE con labels:")
outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                fmt.Printf("  Break outer en [%d,%d]\n", i, j)
                break outer // Sale de ambos loops
            }
            fmt.Printf("  [%d,%d]\n", i, j)
        }
    }
    fmt.Println()
}

    // ==========================================
    // 8. PATRONES COMUNES
    // ==========================================
func demostrarPatronesComunes() {
    fmt.Println("8. PATRONES COMUNES:")

    // Do-while simulado
    fmt.Print("Do-while: ")
    num := 0
    for {
        fmt.Print(num, " ")
        num++
        if num >= 5 {
            break
        }
    }
    fmt.Println()

    // Loop con múltiples variables
    fmt.Print("Loop múltiple: ")
    for i, j := 0, 10; i < j; i, j = i+1, j-1 {
        fmt.Printf("i=%d,j=%d ", i, j)
    }
    fmt.Println()

    // Early return (salida temprana)
    resultado := buscarNumero([]int{1, 2, 3, 4, 5}, 3)
    fmt.Printf("Número encontrado: %v\n", resultado)
}

func buscarNumero(numeros []int, objetivo int) bool {
    for _, num := range numeros {
        if num == objetivo {
            return true // Early return
        }
    }
    return false
}