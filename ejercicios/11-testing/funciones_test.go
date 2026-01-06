// Ejercicio 11: Tests unitarios con casos borde
// Archivo de tests para las funciones implementadas en funciones.go

package main

import (
	"math"
	"testing"
)

// Test para esPalindromo con múltiples casos
func TestEsPalindromo(t *testing.T) {
	// Tabla de casos de prueba
	tests := []struct {
		nombre   string
		input    string
		esperado bool
	}{
		// Casos básicos
		{"palindromo simple", "aba", true},
		{"palindromo con espacios", "anita lava la tina", true},
		{"palindromo mixto", "A man a plan a canal Panama", true},
		{"no palindromo", "hola", false},
		{"palabra vacía", "", true},
		{"un caracter", "a", true},
		{"dos caracteres iguales", "aa", true},
		{"dos caracteres diferentes", "ab", false},
		{"palindromo numeros", "12321", true},
		{"palindromo con mayúsculas", "RaCeCaR", true},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			resultado := esPalindromo(tt.input)
			if resultado != tt.esperado {
				t.Errorf("esPalindromo(%q) = %v, esperado %v", tt.input, resultado, tt.esperado)
			}
		})
	}
}

// Test para max con casos borde
func TestMax(t *testing.T) {
	tests := []struct {
		nombre      string
		input       []float64
		esperado    float64
		debeErrar   bool
	}{
		{"números positivos", []float64{1, 5, 3, 9, 2}, 9, false},
		{"números negativos", []float64{-10, -5, -20, -3}, -3, false},
		{"números mixtos", []float64{-5, 0, 5, -10, 15}, 15, false},
		{"un solo elemento", []float64{42}, 42, false},
		{"dos elementos", []float64{10, 20}, 20, false},
		{"todos iguales", []float64{5, 5, 5, 5}, 5, false},
		{"con decimales", []float64{3.14, 2.71, 1.41, 9.99}, 9.99, false},
		{"slice vacío", []float64{}, 0, true},
		{"con cero", []float64{0, -1, -2}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			resultado, err := max(tt.input)

			if tt.debeErrar {
				if err == nil {
					t.Error("Se esperaba un error pero no se obtuvo ninguno")
				}
				return
			}

			if err != nil {
				t.Errorf("Error inesperado: %v", err)
				return
			}

			if resultado != tt.esperado {
				t.Errorf("max(%v) = %f, esperado %f", tt.input, resultado, tt.esperado)
			}
		})
	}
}

// Test para promedio con casos borde
func TestPromedio(t *testing.T) {
	tests := []struct {
		nombre    string
		input     []float64
		esperado  float64
		debeErrar bool
	}{
		{"números positivos", []float64{2, 4, 6, 8}, 5, false},
		{"números negativos", []float64{-2, -4, -6}, -4, false},
		{"números mixtos", []float64{-10, 0, 10}, 0, false},
		{"un elemento", []float64{42}, 42, false},
		{"con decimales", []float64{1.5, 2.5, 3.5}, 2.5, false},
		{"slice vacío", []float64{}, 0, true},
		{"todos ceros", []float64{0, 0, 0}, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			resultado, err := promedio(tt.input)

			if tt.debeErrar {
				if err == nil {
					t.Error("Se esperaba un error pero no se obtuvo ninguno")
				}
				return
			}

			if err != nil {
				t.Errorf("Error inesperado: %v", err)
				return
			}

			// Comparamos con tolerancia para decimales
			if math.Abs(resultado-tt.esperado) > 0.0001 {
				t.Errorf("promedio(%v) = %f, esperado %f", tt.input, resultado, tt.esperado)
			}
		})
	}
}

// Test para dividir con casos borde
func TestDividir(t *testing.T) {
	tests := []struct {
		nombre    string
		a         float64
		b         float64
		esperado  float64
		debeErrar bool
	}{
		{"división normal", 10, 2, 5, false},
		{"división con decimales", 7, 2, 3.5, false},
		{"dividendo cero", 0, 5, 0, false},
		{"división por cero", 10, 0, 0, true},
		{"números negativos", -10, 2, -5, false},
		{"ambos negativos", -10, -2, 5, false},
		{"división de uno", 5, 1, 5, false},
		{"resultado menor a 1", 1, 2, 0.5, false},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			resultado, err := dividir(tt.a, tt.b)

			if tt.debeErrar {
				if err == nil {
					t.Errorf("dividir(%f, %f): Se esperaba error por división por cero", tt.a, tt.b)
				}
				return
			}

			if err != nil {
				t.Errorf("Error inesperado: %v", err)
				return
			}

			if math.Abs(resultado-tt.esperado) > 0.0001 {
				t.Errorf("dividir(%f, %f) = %f, esperado %f", tt.a, tt.b, resultado, tt.esperado)
			}
		})
	}
}

// Test para factorial
func TestFactorial(t *testing.T) {
	tests := []struct {
		nombre    string
		input     int
		esperado  int
		debeErrar bool
	}{
		{"factorial de 0", 0, 1, false},
		{"factorial de 1", 1, 1, false},
		{"factorial de 5", 5, 120, false},
		{"factorial de 10", 10, 3628800, false},
		{"número negativo", -5, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			resultado, err := factorial(tt.input)

			if tt.debeErrar {
				if err == nil {
					t.Error("Se esperaba un error pero no se obtuvo ninguno")
				}
				return
			}

			if err != nil {
				t.Errorf("Error inesperado: %v", err)
				return
			}

			if resultado != tt.esperado {
				t.Errorf("factorial(%d) = %d, esperado %d", tt.input, resultado, tt.esperado)
			}
		})
	}
}

// Test para esPrimo
func TestEsPrimo(t *testing.T) {
	tests := []struct {
		nombre   string
		input    int
		esperado bool
	}{
		{"número primo pequeño", 2, true},
		{"número primo", 17, true},
		{"número primo grande", 97, true},
		{"no primo", 4, false},
		{"no primo compuesto", 15, false},
		{"cero", 0, false},
		{"uno", 1, false},
		{"número negativo", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			resultado := esPrimo(tt.input)
			if resultado != tt.esperado {
				t.Errorf("esPrimo(%d) = %v, esperado %v", tt.input, resultado, tt.esperado)
			}
		})
	}
}

// Benchmark para esPalindromo
func BenchmarkEsPalindromo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		esPalindromo("A man a plan a canal Panama")
	}
}

// Benchmark para factorial
func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factorial(10)
	}
}
