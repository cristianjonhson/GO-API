// 05. STRUCTS (ESTRUCTURAS)
// Los structs son tipos de datos compuestos que agrupan campos relacionados

package main

import "fmt"

// ==========================================
// 1. DEFINICIÓN BÁSICA DE STRUCT
// ==========================================
type Persona struct {
    Nombre string
    Edad   int
    Email  string
}

// ==========================================
// 2. STRUCT CON CAMPOS EXPORTADOS Y NO EXPORTADOS
// ==========================================
type Usuario struct {
    ID       int    // Exportado (mayúscula inicial)
    Username string // Exportado
    password string // No exportado (minúscula inicial, privado al paquete)
}

// ==========================================
// 3. STRUCT ANIDADO
// ==========================================
type Direccion struct {
    Calle  string
    Ciudad string
    CP     string
}

type Empleado struct {
    Nombre    string
    Edad      int
    Direccion Direccion // Struct dentro de otro struct
}

// ==========================================
// 4. STRUCT CON EMBEDDING (COMPOSICIÓN)
// ==========================================
type Animal struct {
    Nombre string
    Edad   int
}

type Perro struct {
    Animal // Embedding anónimo (hereda campos)
    Raza   string
}

// ==========================================
// 5. STRUCT VACÍO
// ==========================================
type Vacio struct{} // No ocupa memoria (0 bytes)

// ==========================================
// 6. MÉTODOS ASOCIADOS A STRUCTS
// ==========================================
type Rectangulo struct {
    Base   float64
    Altura float64
}

// Método con receiver por valor
func (r Rectangulo) Area() float64 {
    return r.Base * r.Altura
}

// Método con receiver por puntero (puede modificar)
func (r *Rectangulo) Escalar(factor float64) {
    r.Base *= factor
    r.Altura *= factor
}

// ==========================================
// 7. CONSTRUCTOR (PATRÓN COMÚN)
// ==========================================
func NuevaPersona(nombre string, edad int) *Persona {
    return &Persona{
        Nombre: nombre,
        Edad:   edad,
        Email:  "", // valor por defecto
    }
}

// ==========================================
// 8. STRUCT CON TAGS (metadatos)
// ==========================================
type Producto struct {
    ID     int     `json:"id" db:"product_id"`
    Nombre string  `json:"name" db:"product_name"`
    Precio float64 `json:"price" db:"product_price"`
}

func main() {
    fmt.Println("=== STRUCTS EN GO ===\n")

    // ==========================================
    // CREACIÓN E INICIALIZACIÓN
    // ==========================================

    // Forma 1: Declaración y asignación separada
    var p1 Persona
    p1.Nombre = "Ana"
    p1.Edad = 25
    p1.Email = "ana@example.com"
    fmt.Printf("Persona 1: %+v\n", p1)

    // Forma 2: Literal de struct con nombres de campos
    p2 := Persona{
        Nombre: "Carlos",
        Edad:   30,
        Email:  "carlos@example.com",
    }
    fmt.Printf("Persona 2: %+v\n", p2)

    // Forma 3: Literal de struct sin nombres (orden importa)
    p3 := Persona{"Luis", 28, "luis@example.com"}
    fmt.Printf("Persona 3: %+v\n", p3)

    // Forma 4: Inicialización parcial (resto son valores cero)
    p4 := Persona{Nombre: "María"} // Edad=0, Email=""
    fmt.Printf("Persona 4: %+v\n\n", p4)

    // ==========================================
    // ACCESO A CAMPOS
    // ==========================================
    fmt.Println("Acceso a campos:")
    fmt.Printf("Nombre: %s, Edad: %d\n", p2.Nombre, p2.Edad)
    
    // Modificar campos
    p2.Edad = 31
    fmt.Printf("Nueva edad: %d\n\n", p2.Edad)

    // ==========================================
    // PUNTEROS A STRUCTS
    // ==========================================
    fmt.Println("Punteros a structs:")
    puntero := &p2
    fmt.Printf("A través del puntero: %s\n", puntero.Nombre) // Go desreferencia automáticamente
    fmt.Printf("Dirección de memoria: %p\n\n", puntero)

    // ==========================================
    // STRUCT ANIDADO
    // ==========================================
    fmt.Println("Struct anidado:")
    empleado := Empleado{
        Nombre: "Pedro",
        Edad:   35,
        Direccion: Direccion{
            Calle:  "Av. Siempre Viva 123",
            Ciudad: "Springfield",
            CP:     "12345",
        },
    }
    fmt.Printf("Empleado: %+v\n", empleado)
    fmt.Printf("Ciudad: %s\n\n", empleado.Direccion.Ciudad)

    // ==========================================
    // EMBEDDING (COMPOSICIÓN)
    // ==========================================
    fmt.Println("Embedding:")
    perro := Perro{
        Animal: Animal{
            Nombre: "Rex",
            Edad:   3,
        },
        Raza: "Pastor Alemán",
    }
    
    // Acceso directo a campos del Animal embebido
    fmt.Printf("Perro: %s, Edad: %d, Raza: %s\n", perro.Nombre, perro.Edad, perro.Raza)
    
    // También se puede acceder explícitamente
    fmt.Printf("Nombre (explícito): %s\n\n", perro.Animal.Nombre)

    // ==========================================
    // COMPARACIÓN DE STRUCTS
    // ==========================================
    fmt.Println("Comparación:")
    persona1 := Persona{"Juan", 25, "juan@example.com"}
    persona2 := Persona{"Juan", 25, "juan@example.com"}
    persona3 := Persona{"Ana", 25, "ana@example.com"}
    
    fmt.Printf("persona1 == persona2: %v\n", persona1 == persona2) // true
    fmt.Printf("persona1 == persona3: %v\n\n", persona1 == persona3) // false

    // ==========================================
    // STRUCTS ANÓNIMOS
    // ==========================================
    fmt.Println("Struct anónimo:")
    config := struct {
        Host string
        Port int
    }{
        Host: "localhost",
        Port: 8080,
    }
    fmt.Printf("Config: %+v\n\n", config)

    // ==========================================
    // MÉTODOS
    // ==========================================
    fmt.Println("Métodos:")
    rect := Rectangulo{Base: 5, Altura: 3}
    fmt.Printf("Rectángulo: %+v\n", rect)
    fmt.Printf("Área: %.2f\n", rect.Area())
    
    rect.Escalar(2)
    fmt.Printf("Rectángulo escalado x2: %+v\n", rect)
    fmt.Printf("Nueva área: %.2f\n\n", rect.Area())

    // ==========================================
    // CONSTRUCTOR
    // ==========================================
    fmt.Println("Constructor:")
    nuevaPersona := NuevaPersona("Elena", 28)
    fmt.Printf("Nueva persona: %+v\n\n", *nuevaPersona)

    // ==========================================
    // COPIA VS REFERENCIA
    // ==========================================
    fmt.Println("Copia vs Referencia:")
    original := Persona{Nombre: "Original", Edad: 20}
    
    // Copia por valor
    copia := original
    copia.Nombre = "Copia"
    fmt.Printf("Original: %s, Copia: %s\n", original.Nombre, copia.Nombre)
    
    // Referencia (puntero)
    referencia := &original
    referencia.Nombre = "Modificado"
    fmt.Printf("Original modificado: %s\n\n", original.Nombre)

    // ==========================================
    // STRUCT CON SLICES Y MAPS
    // ==========================================
    fmt.Println("Struct con colecciones:")
    type Curso struct {
        Nombre     string
        Estudiantes []string
        Notas      map[string]int
    }
    
    curso := Curso{
        Nombre:     "Go Avanzado",
        Estudiantes: []string{"Ana", "Carlos", "Luis"},
        Notas: map[string]int{
            "Ana":    95,
            "Carlos": 88,
            "Luis":   92,
        },
    }
    fmt.Printf("Curso: %+v\n", curso)
    fmt.Printf("Nota de Ana: %d\n\n", curso.Notas["Ana"])

    // ==========================================
    // VALORES CERO DE STRUCTS
    // ==========================================
    fmt.Println("Valores cero:")
    var vacio Persona
    fmt.Printf("Struct vacío: %+v\n", vacio)
    fmt.Printf("Nombre vacío: '%s' (len: %d)\n", vacio.Nombre, len(vacio.Nombre))
    fmt.Printf("Edad cero: %d\n", vacio.Edad)
}