// 02. TIPOS DE DATOS EN GO
// Go es un lenguaje con tipado estático y fuerte

package main

import "fmt"

func main() {
    fmt.Println("=== TIPOS DE DATOS EN GO ===\n")

    // ==========================================
    // 1. TIPOS BOOLEANOS
    // ==========================================
    var activo bool = true
    var inactivo bool = false
    fmt.Printf("bool: %v, %v\n\n", activo, inactivo)

    // ==========================================
    // 2. TIPOS NUMÉRICOS ENTEROS
    // ==========================================
    
    // Enteros con signo (pueden ser negativos)
    var i8 int8 = 127           // -128 a 127 (8 bits = 1 byte)
    var i16 int16 = 32767       // -32768 a 32767 (16 bits = 2 bytes)
    var i32 int32 = 2147483647  // -2^31 a 2^31-1 (32 bits = 4 bytes)
    var i64 int64 = 9223372036854775807 // -2^63 a 2^63-1 (64 bits = 8 bytes)
    
    // int: tamaño depende de la arquitectura (32 o 64 bits)
    var i int = 100
    
    fmt.Println("Enteros con signo:")
    fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d, int: %d\n\n", 
        i8, i16, i32, i64, i)

    // Enteros sin signo (solo positivos)
    var ui8 uint8 = 255         // 0 a 255 (8 bits)
    var ui16 uint16 = 65535     // 0 a 65535 (16 bits)
    var ui32 uint32 = 4294967295 // 0 a 2^32-1 (32 bits)
    var ui64 uint64 = 18446744073709551615 // 0 a 2^64-1 (64 bits)
    
    // uint: tamaño depende de la arquitectura
    var ui uint = 200
    
    fmt.Println("Enteros sin signo:")
    fmt.Printf("uint8: %d, uint16: %d, uint32: %d, uint64: %d, uint: %d\n\n", 
        ui8, ui16, ui32, ui64, ui)

    // Alias especiales
    var b byte = 255    // alias de uint8 (usado para datos binarios)
    var r rune = '🔥'   // alias de int32 (usado para caracteres Unicode)
    
    fmt.Printf("byte (uint8): %d, rune (int32): %d (%c)\n\n", b, r, r)

    // ==========================================
    // 3. TIPOS NUMÉRICOS DE PUNTO FLOTANTE
    // ==========================================
    var f32 float32 = 3.14159   // 32 bits (6-7 decimales de precisión)
    var f64 float64 = 3.141592653589793 // 64 bits (15-16 decimales)
    
    fmt.Println("Punto flotante:")
    fmt.Printf("float32: %.5f, float64: %.15f\n\n", f32, f64)

    // ==========================================
    // 4. TIPOS COMPLEJOS (números complejos)
    // ==========================================
    var c64 complex64 = 1 + 2i      // parte real e imaginaria float32
    var c128 complex128 = 3 + 4i    // parte real e imaginaria float64
    
    fmt.Println("Números complejos:")
    fmt.Printf("complex64: %v, complex128: %v\n", c64, c128)
    fmt.Printf("Real: %.0f, Imaginaria: %.0f\n\n", real(c128), imag(c128))

    // ==========================================
    // 5. STRINGS (cadenas de texto)
    // ==========================================
    var texto string = "Hola, Go!"
    var multilinea string = `Este es un
    string multilínea
    usando backticks`
    
    // Strings son inmutables
    fmt.Println("Strings:")
    fmt.Printf("Texto: %s\n", texto)
    fmt.Printf("Longitud: %d caracteres\n", len(texto))
    fmt.Printf("Primer byte: %c\n", texto[0])
    fmt.Printf("Multilínea:\n%s\n\n", multilinea)

    // Concatenación
    nombre := "Juan"
    apellido := "Pérez"
    nombreCompleto := nombre + " " + apellido
    fmt.Printf("Concatenación: %s\n\n", nombreCompleto)

    // ==========================================
    // 6. ARRAYS (tamaño fijo)
    // ==========================================
    var arr1 [5]int              // Array de 5 enteros (inicializado en 0)
    arr2 := [3]string{"a", "b", "c"} // Array inicializado
    arr3 := [...]int{1, 2, 3, 4} // Tamaño automático según elementos
    
    fmt.Println("Arrays:")
    fmt.Printf("arr1: %v (longitud: %d)\n", arr1, len(arr1))
    fmt.Printf("arr2: %v\n", arr2)
    fmt.Printf("arr3: %v\n\n", arr3)

    // ==========================================
    // 7. SLICES (tamaño dinámico)
    // ==========================================
    var slice1 []int                    // slice vacío
    slice2 := []string{"Go", "Python", "Java"}
    slice3 := make([]int, 3, 5)        // longitud 3, capacidad 5
    
    // Operaciones con slices
    slice2 = append(slice2, "JavaScript") // Agregar elementos
    subslice := slice2[1:3]               // Subslice [Python Java]
    
    fmt.Println("Slices:")
    fmt.Printf("slice1: %v (len: %d, cap: %d)\n", slice1, len(slice1), cap(slice1))
    fmt.Printf("slice2: %v (len: %d, cap: %d)\n", slice2, len(slice2), cap(slice2))
    fmt.Printf("slice3: %v (len: %d, cap: %d)\n", slice3, len(slice3), cap(slice3))
    fmt.Printf("subslice: %v\n\n", subslice)

    // ==========================================
    // 8. MAPS (diccionarios/hash tables)
    // ==========================================
    var map1 map[string]int                    // map vacío (nil)
    map2 := make(map[string]int)               // map inicializado vacío
    map3 := map[string]int{                    // map con valores iniciales
        "edad":   25,
        "altura": 175,
    }
    
    // Operaciones con maps
    map2["go"] = 2009
    map2["python"] = 1991
    
    valor, existe := map2["go"]                // Verificar si existe una clave
    delete(map2, "python")                     // Eliminar elemento
    
    fmt.Println("Maps:")
    fmt.Printf("map1: %v\n", map1)
    fmt.Printf("map2: %v\n", map2)
    fmt.Printf("map3: %v\n", map3)
    fmt.Printf("Clave 'go' existe: %v, valor: %d\n\n", existe, valor)

    // ==========================================
    // 9. PUNTEROS
    // ==========================================
    numero := 42
    var puntero *int = &numero    // & obtiene la dirección de memoria
    
    fmt.Println("Punteros:")
    fmt.Printf("Valor: %d\n", numero)
    fmt.Printf("Dirección: %p\n", puntero)
    fmt.Printf("Valor apuntado: %d\n", *puntero) // * desreferencia el puntero
    
    *puntero = 100 // Modificar el valor a través del puntero
    fmt.Printf("Nuevo valor: %d\n\n", numero)

    // ==========================================
    // 10. INTERFACES (tipo vacío)
    // ==========================================
    var cualquierCosa interface{} // Puede contener cualquier tipo
    
    cualquierCosa = 42
    fmt.Printf("interface{} con int: %v (tipo: %T)\n", cualquierCosa, cualquierCosa)
    
    cualquierCosa = "texto"
    fmt.Printf("interface{} con string: %v (tipo: %T)\n\n", cualquierCosa, cualquierCosa)

    // ==========================================
    // 11. VALORES CERO (zero values)
    // ==========================================
    var zeroBool bool
    var zeroInt int
    var zeroFloat float64
    var zeroString string
    var zeroSlice []int
    var zeroMap map[string]int
    var zeroPtr *int
    var zeroFunc func()
    var zeroInterface interface{}
    
    fmt.Println("Valores cero de cada tipo:")
    fmt.Printf("bool: %v\n", zeroBool)
    fmt.Printf("int: %v\n", zeroInt)
    fmt.Printf("float64: %v\n", zeroFloat)
    fmt.Printf("string: %q\n", zeroString)
    fmt.Printf("[]int: %v\n", zeroSlice)
    fmt.Printf("map: %v\n", zeroMap)
    fmt.Printf("*int: %v\n", zeroPtr)
    fmt.Printf("func(): %v\n", zeroFunc)
    fmt.Printf("interface{}: %v\n", zeroInterface)
}