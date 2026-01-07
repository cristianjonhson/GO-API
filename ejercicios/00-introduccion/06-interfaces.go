// 06. INTERFACES
// Las interfaces definen comportamientos (conjuntos de métodos)

package main

import (
    "fmt"
    "math"
)

// ==========================================
// 1. DEFINICIÓN BÁSICA DE INTERFACE
// ==========================================
type Forma interface {
    Area() float64
    Perimetro() float64
}

// ==========================================
// 2. TIPOS QUE IMPLEMENTAN LA INTERFACE
// ==========================================
type Rectangulo struct {
    Base   float64
    Altura float64
}

func (r Rectangulo) Area() float64 {
    return r.Base * r.Altura
}

func (r Rectangulo) Perimetro() float64 {
    return 2 * (r.Base + r.Altura)
}

type Circulo struct {
    Radio float64
}

func (c Circulo) Area() float64 {
    return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
    return 2 * math.Pi * c.Radio
}

// ==========================================
// 3. FUNCIÓN QUE ACEPTA INTERFACE
// ==========================================
func imprimirInfo(f Forma) {
    fmt.Printf("Área: %.2f, Perímetro: %.2f\n", f.Area(), f.Perimetro())
}

// ==========================================
// 4. INTERFACE VACÍA (interface{} o any)
// ==========================================
func imprimirCualquierCosa(valor interface{}) {
    fmt.Printf("Valor: %v (Tipo: %T)\n", valor, valor)
}

// ==========================================
// 5. TYPE ASSERTION Y TYPE SWITCH
// ==========================================
func describir(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Es un entero: %d\n", v)
    case string:
        fmt.Printf("Es un string: %s\n", v)
    case Forma:
        fmt.Printf("Es una Forma con área: %.2f\n", v.Area())
    default:
        fmt.Printf("Tipo desconocido: %T\n", v)
    }
}

// ==========================================
// 6. INTERFACES MÚLTIPLES
// ==========================================
type Figura interface {
    Dibujar()
}

type FormaCompleta interface {
    Forma  // Embebiendo interface Forma
    Figura // Embebiendo interface Figura
}

func (r Rectangulo) Dibujar() {
    fmt.Println("Dibujando rectángulo...")
}

func (c Circulo) Dibujar() {
    fmt.Println("Dibujando círculo...")
}

// ==========================================
// 7. INTERFACES ESTÁNDAR DE GO
// ==========================================

// Stringer: permite personalizar fmt.Print
type Punto struct {
    X, Y int
}

func (p Punto) String() string {
    return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Error: interface para manejo de errores
type MiError struct {
    Codigo  int
    Mensaje string
}

func (e MiError) Error() string {
    return fmt.Sprintf("[Error %d] %s", e.Codigo, e.Mensaje)
}

// ==========================================
// 8. INTERFACE CON MÚLTIPLES MÉTODOS
// ==========================================
type Almacenamiento interface {
    Guardar(dato string) error
    Recuperar(id string) (string, error)
    Eliminar(id string) error
}

type AlmacenamientoMemoria struct {
    datos map[string]string
}

func (a *AlmacenamientoMemoria) Guardar(dato string) error {
    if a.datos == nil {
        a.datos = make(map[string]string)
    }
    id := fmt.Sprintf("id_%d", len(a.datos)+1)
    a.datos[id] = dato
    return nil
}

func (a *AlmacenamientoMemoria) Recuperar(id string) (string, error) {
    if dato, existe := a.datos[id]; existe {
        return dato, nil
    }
    return "", fmt.Errorf("dato no encontrado")
}

func (a *AlmacenamientoMemoria) Eliminar(id string) error {
    delete(a.datos, id)
    return nil
}

// ==========================================
// 9. INTERFACE VACÍA CON REFLEXIÓN
// ==========================================
func esNil(i interface{}) bool {
    return i == nil
}

func main() {
    fmt.Println("=== INTERFACES EN GO ===\n")

    // ==========================================
    // IMPLEMENTACIÓN IMPLÍCITA
    // ==========================================
    fmt.Println("1. Implementación implícita:")
    rect := Rectangulo{Base: 5, Altura: 3}
    circ := Circulo{Radio: 4}
    
    // No necesitamos declarar explícitamente que implementan Forma
    var forma1 Forma = rect
    var forma2 Forma = circ
    
    fmt.Print("Rectángulo - ")
    imprimirInfo(forma1)
    fmt.Print("Círculo - ")
    imprimirInfo(forma2)
    fmt.Println()

    // ==========================================
    // SLICE DE INTERFACES
    // ==========================================
    fmt.Println("2. Slice de interfaces:")
    formas := []Forma{
        Rectangulo{Base: 4, Altura: 6},
        Circulo{Radio: 3},
        Rectangulo{Base: 2, Altura: 8},
    }
    
    for i, forma := range formas {
        fmt.Printf("Forma %d - ", i+1)
        imprimirInfo(forma)
    }
    fmt.Println()

    // ==========================================
    // INTERFACE VACÍA (interface{} / any)
    // ==========================================
    fmt.Println("3. Interface vacía:")
    imprimirCualquierCosa(42)
    imprimirCualquierCosa("Hola")
    imprimirCualquierCosa(3.14)
    imprimirCualquierCosa([]int{1, 2, 3})
    fmt.Println()

    // ==========================================
    // TYPE ASSERTION
    // ==========================================
    fmt.Println("4. Type assertion:")
    var cualquierCosa interface{} = "Hola, Go!"
    
    // Type assertion segura
    if str, ok := cualquierCosa.(string); ok {
        fmt.Printf("Es un string: %s\n", str)
    } else {
        fmt.Println("No es un string")
    }
    
    // Type assertion directa (panic si falla)
    str := cualquierCosa.(string)
    fmt.Printf("String directo: %s\n\n", str)

    // ==========================================
    // TYPE SWITCH
    // ==========================================
    fmt.Println("5. Type switch:")
    describir(42)
    describir("texto")
    describir(rect)
    describir(true)
    fmt.Println()

    // ==========================================
    // INTERFACES MÚLTIPLES
    // ==========================================
    fmt.Println("6. Interfaces múltiples:")
    var fc FormaCompleta = rect
    fc.Dibujar()
    fmt.Printf("Área: %.2f\n\n", fc.Area())

    // ==========================================
    // STRINGER INTERFACE
    // ==========================================
    fmt.Println("7. Stringer interface:")
    punto := Punto{X: 10, Y: 20}
    fmt.Println("Punto:", punto) // Usa String() automáticamente
    fmt.Println()

    // ==========================================
    // ERROR INTERFACE
    // ==========================================
    fmt.Println("8. Error interface:")
    err := MiError{Codigo: 404, Mensaje: "Recurso no encontrado"}
    fmt.Println("Error:", err)
    fmt.Println()

    // ==========================================
    // INTERFACE CON MÚLTIPLES MÉTODOS
    // ==========================================
    fmt.Println("9. Interface con múltiples métodos:")
    var almacen Almacenamiento = &AlmacenamientoMemoria{}
    
    almacen.Guardar("Dato 1")
    almacen.Guardar("Dato 2")
    
    if dato, err := almacen.Recuperar("id_1"); err == nil {
        fmt.Printf("Recuperado: %s\n", dato)
    }
    fmt.Println()

    // ==========================================
    // VERIFICAR NIL
    // ==========================================
    fmt.Println("10. Verificar nil:")
    var formaVacia Forma
    fmt.Printf("¿Forma es nil? %v\n", formaVacia == nil)
    fmt.Printf("¿42 es nil? %v\n", esNil(42))
    fmt.Println()

    // ==========================================
    // CONVERSIÓN ENTRE INTERFACES
    // ==========================================
    fmt.Println("11. Conversión entre interfaces:")
    var cualquier interface{} = circ
    
    if forma, ok := cualquier.(Forma); ok {
        fmt.Printf("Convertido a Forma - Área: %.2f\n", forma.Area())
    }
    
    if figura, ok := cualquier.(Figura); ok {
        figura.Dibujar()
    }
}