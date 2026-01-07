// 11. TESTING Y BENCHMARKING
// Go incluye un framework de testing integrado

package main

import "fmt"

/*
==========================================
TESTING EN GO
==========================================

Go tiene soporte nativo para testing en el paquete "testing".

ESTRUCTURA DE ARCHIVOS:
- Código: archivo.go
- Tests: archivo_test.go (mismo directorio)

CONVENCIONES:
- Función de test: func TestNombre(t *testing.T)
- Función de benchmark: func BenchmarkNombre(b *testing.B)
- Función de ejemplo: func ExampleNombre()

COMANDOS:
go test                  # Ejecutar tests
go test -v               # Modo verbose
go test -cover           # Con cobertura
go test -bench=.         # Ejecutar benchmarks
go test -race            # Detector de race conditions
go test ./...            # Tests recursivos

==========================================
EJEMPLO DE TEST BÁSICO
==========================================

// archivo: matematicas.go
package matematicas

func Sumar(a, b int) int {
    return a + b
}

func Dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("división por cero")
    }
    return a / b, nil
}

// archivo: matematicas_test.go
package matematicas

import "testing"

func TestSumar(t *testing.T) {
    resultado := Sumar(2, 3)
    esperado := 5
    
    if resultado != esperado {
        t.Errorf("Sumar(2, 3) = %d; esperado %d", resultado, esperado)
    }
}

func TestDividir(t *testing.T) {
    resultado, err := Dividir(10, 2)
    if err != nil {
        t.Fatalf("Error inesperado: %v", err)
    }
    
    esperado := 5.0
    if resultado != esperado {
        t.Errorf("Dividir(10, 2) = %f; esperado %f", resultado, esperado)
    }
}

func TestDividirPorCero(t *testing.T) {
    _, err := Dividir(10, 0)
    if err == nil {
        t.Error("Se esperaba un error al dividir por cero")
    }
}

==========================================
TABLE-DRIVEN TESTS
==========================================

func TestSumarTabla(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        esperado int
    }{
        {"positivos", 2, 3, 5},
        {"negativos", -2, -3, -5},
        {"mixtos", -2, 3, 1},
        {"ceros", 0, 0, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            resultado := Sumar(tt.a, tt.b)
            if resultado != tt.esperado {
                t.Errorf("Sumar(%d, %d) = %d; esperado %d",
                    tt.a, tt.b, resultado, tt.esperado)
            }
        })
    }
}

==========================================
SUBTESTS
==========================================

func TestOperaciones(t *testing.T) {
    t.Run("Suma", func(t *testing.T) {
        if Sumar(2, 3) != 5 {
            t.Error("Suma falló")
        }
    })
    
    t.Run("Resta", func(t *testing.T) {
        if Restar(5, 3) != 2 {
            t.Error("Resta falló")
        }
    })
}

==========================================
BENCHMARKS
==========================================

func BenchmarkSumar(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Sumar(2, 3)
    }
}

func BenchmarkDividir(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Dividir(10, 2)
    }
}

// Ejecutar: go test -bench=.
// Resultado: BenchmarkSumar-8   1000000000   0.25 ns/op

==========================================
EXAMPLES (DOCUMENTACIÓN)
==========================================

func ExampleSumar() {
    resultado := Sumar(2, 3)
    fmt.Println(resultado)
    // Output: 5
}

func ExampleDividir() {
    resultado, _ := Dividir(10, 2)
    fmt.Printf("%.1f", resultado)
    // Output: 5.0
}

==========================================
HELPERS
==========================================

func TestHelper(t *testing.T) {
    verificarResultado(t, 2, 3, 5)
    verificarResultado(t, 10, 5, 15)
}

func verificarResultado(t *testing.T, a, b, esperado int) {
    t.Helper() // Marca esta función como helper
    resultado := Sumar(a, b)
    if resultado != esperado {
        t.Errorf("Sumar(%d, %d) = %d; esperado %d",
            a, b, resultado, esperado)
    }
}

==========================================
SETUP Y TEARDOWN
==========================================

func TestMain(m *testing.M) {
    // Setup global
    fmt.Println("Setup antes de todos los tests")
    
    // Ejecutar tests
    code := m.Run()
    
    // Teardown global
    fmt.Println("Cleanup después de todos los tests")
    
    os.Exit(code)
}

func TestConSetup(t *testing.T) {
    // Setup del test
    db := setupDatabase()
    defer db.Close() // Teardown del test
    
    // Test...
}

==========================================
COBERTURA
==========================================

go test -cover                    # Porcentaje de cobertura
go test -coverprofile=cover.out   # Generar archivo de cobertura
go tool cover -html=cover.out     # Ver cobertura en HTML

==========================================
MOCKING
==========================================

// Interface para mocking
type DatabaseInterface interface {
    Get(id string) (string, error)
    Save(id, value string) error
}

// Mock para testing
type MockDatabase struct {
    data map[string]string
}

func (m *MockDatabase) Get(id string) (string, error) {
    if val, ok := m.data[id]; ok {
        return val, nil
    }
    return "", errors.New("no encontrado")
}

func (m *MockDatabase) Save(id, value string) error {
    m.data[id] = value
    return nil
}

func TestConMock(t *testing.T) {
    mock := &MockDatabase{
        data: map[string]string{"1": "valor1"},
    }
    
    val, err := mock.Get("1")
    if err != nil {
        t.Fatal(err)
    }
    if val != "valor1" {
        t.Errorf("esperado 'valor1', obtenido '%s'", val)
    }
}

==========================================
MEJORES PRÁCTICAS
==========================================

✅ DO:
  • Usar table-driven tests
  • Tests independientes y determinísticos
  • Nombres descriptivos (TestFunción_Caso_Resultado)
  • Usar t.Helper() en funciones auxiliares
  • Verificar errores (t.Errorf, t.Fatalf)
  • Escribir tests antes de arreglar bugs
  • Mantener tests simples y legibles

❌ DON'T:
  • Tests que dependen de orden de ejecución
  • Tests que modifican estado global
  • Ignorar errores en tests
  • Tests demasiado complejos
  • Depender de timings específicos
  • Tests que requieren recursos externos sin mocks

==========================================
HERRAMIENTAS ÚTILES
==========================================

# Testify (biblioteca popular)
go get github.com/stretchr/testify

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestConTestify(t *testing.T) {
    assert.Equal(t, 5, Sumar(2, 3))
    assert.NoError(t, err)
    assert.True(t, condicion)
}

# GoMock (generación de mocks)
go get github.com/golang/mock/gomock

# Race detector
go test -race

# Coverage profile
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
*/

func main() {
    fmt.Println("=== TESTING EN GO ===\n")
    fmt.Println("Este archivo contiene documentación sobre testing.")
    fmt.Println("Para ejemplos prácticos, crea archivos *_test.go\n")
    
    fmt.Println("Comandos útiles:")
    fmt.Println("  go test              # Ejecutar tests")
    fmt.Println("  go test -v           # Modo verbose")
    fmt.Println("  go test -cover       # Con cobertura")
    fmt.Println("  go test -bench=.     # Benchmarks")
    fmt.Println("  go test -race        # Race detector")
    fmt.Println("\nRevisa los comentarios del código para más detalles.")
}