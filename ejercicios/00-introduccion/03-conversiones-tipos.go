// 03. CONVERSIONES Y CASTEO DE TIPOS
// Go requiere conversiones explícitas entre tipos

package main

import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println("=== CONVERSIONES DE TIPOS ===\n")

    // ==========================================
    // 1. CONVERSIONES NUMÉRICAS
    // ==========================================
    var entero int = 42
    var flotante float64 = float64(entero)  // int a float64
    var entero2 int = int(flotante)         // float64 a int (pierde decimales)
    
    fmt.Println("Conversiones numéricas:")
    fmt.Printf("int: %d -> float64: %f -> int: %d\n\n", entero, flotante, entero2)

    // Conversiones entre tipos enteros
    var i8 int8 = 127
    var i16 int16 = int16(i8)
    var ui uint = uint(i16)
    
    fmt.Printf("int8: %d -> int16: %d -> uint: %d\n\n", i8, i16, ui)

    // ==========================================
    // 2. STRING A NÚMERO (strconv)
    // ==========================================
    str := "123"
    
    // String a int
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Printf("String '%s' a int: %d\n", str, num)

    // String a int64
    num64, _ := strconv.ParseInt("9876543210", 10, 64)
    fmt.Printf("String a int64: %d\n", num64)

    // String a float
    pi, _ := strconv.ParseFloat("3.14159", 64)
    fmt.Printf("String a float64: %f\n", pi)

    // String a bool
    activo, _ := strconv.ParseBool("true")
    fmt.Printf("String a bool: %v\n\n", activo)

    // ==========================================
    // 3. NÚMERO A STRING
    // ==========================================
    numero := 42
    
    // int a string
    strNum := strconv.Itoa(numero)
    fmt.Printf("int %d a string: '%s'\n", numero, strNum)

    // int64 a string
    var grande int64 = 9876543210
    strGrande := strconv.FormatInt(grande, 10) // base 10
    fmt.Printf("int64 a string: '%s'\n", strGrande)

    // float a string
    piValue := 3.14159265359
    strPi := strconv.FormatFloat(piValue, 'f', 2, 64) // formato, precisión, bits
    fmt.Printf("float64 a string: '%s'\n", strPi)

    // bool a string
    verdadero := true
    strBool := strconv.FormatBool(verdadero)
    fmt.Printf("bool a string: '%s'\n\n", strBool)

    // ==========================================
    // 4. BYTE/RUNE Y STRING
    // ==========================================
    texto := "Hola"
    
    // String a []byte (slice de bytes)
    bytes := []byte(texto)
    fmt.Printf("String '%s' a []byte: %v\n", texto, bytes)

    // []byte a string
    nuevoTexto := string(bytes)
    fmt.Printf("[]byte %v a string: '%s'\n", bytes, nuevoTexto)

    // Rune (int32) a string
    emoji := '😀'
    strEmoji := string(emoji)
    fmt.Printf("Rune %d (%c) a string: '%s'\n\n", emoji, emoji, strEmoji)

    // ==========================================
    // 5. CONVERSIONES CON fmt.Sprintf
    // ==========================================
    edad := 25
    altura := 1.75
    
    mensaje := fmt.Sprintf("Edad: %d, Altura: %.2f", edad, altura)
    fmt.Printf("Usando fmt.Sprintf: '%s'\n\n", mensaje)

    // ==========================================
    // 6. TYPE ASSERTIONS (interfaces)
    // ==========================================
    var valor interface{} = "Hola, Go!"
    
    // Aserción de tipo segura (con verificación)
    if str, ok := valor.(string); ok {
        fmt.Printf("El valor es un string: '%s'\n", str)
    } else {
        fmt.Println("El valor NO es un string")
    }

    // Aserción de tipo directa (puede causar panic si falla)
    strDirecto := valor.(string)
    fmt.Printf("Aserción directa: '%s'\n\n", strDirecto)

    // ==========================================
    // 7. SWITCH DE TIPOS
    // ==========================================
    var datos interface{} = 42
    
    switch v := datos.(type) {
    case int:
        fmt.Printf("Es un int: %d\n", v)
    case string:
        fmt.Printf("Es un string: %s\n", v)
    case bool:
        fmt.Printf("Es un bool: %v\n", v)
    default:
        fmt.Printf("Tipo desconocido: %T\n", v)
    }
}