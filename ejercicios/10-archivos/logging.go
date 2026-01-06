// Ejercicio 10.2: Sistema de logging de operaciones
// Registra cada acción del usuario en un archivo log.txt con timestamp.

package main

import (
	"fmt"
	"os"
	"time"
)

// Logger maneja el registro de operaciones en archivo
type Logger struct {
	archivoRuta string
	archivo     *os.File
}

// NuevoLogger crea un nuevo logger y abre el archivo en modo append
func NuevoLogger(archivoRuta string) (*Logger, error) {
	// Abrimos el archivo en modo append (O_APPEND) o lo creamos si no existe
	archivo, err := os.OpenFile(archivoRuta, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("error al abrir archivo de log: %v", err)
	}

	logger := &Logger{
		archivoRuta: archivoRuta,
		archivo:     archivo,
	}

	// Registramos el inicio del programa
	logger.Registrar("INICIO", "Sistema de logging iniciado")

	return logger, nil
}

// Registrar escribe una entrada en el log con timestamp
func (l *Logger) Registrar(operacion, detalle string) error {
	// Obtenemos el timestamp actual
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Formateamos la línea del log
	linea := fmt.Sprintf("[%s] %s: %s\n", timestamp, operacion, detalle)

	// Escribimos en el archivo
	_, err := l.archivo.WriteString(linea)
	if err != nil {
		return fmt.Errorf("error al escribir en log: %v", err)
	}

	// También mostramos en consola para feedback
	fmt.Printf("📝 Log: %s", linea)

	return nil
}

// Cerrar cierra el archivo de log
func (l *Logger) Cerrar() error {
	l.Registrar("FIN", "Sistema de logging cerrado")
	return l.archivo.Close()
}

// Calculadora con logging de operaciones
type Calculadora struct {
	logger *Logger
}

// NuevaCalculadora crea una calculadora con logging
func NuevaCalculadora(logger *Logger) *Calculadora {
	return &Calculadora{logger: logger}
}

// Sumar realiza la suma y la registra
func (c *Calculadora) Sumar(a, b float64) float64 {
	resultado := a + b
	c.logger.Registrar("SUMA", fmt.Sprintf("%.2f + %.2f = %.2f", a, b, resultado))
	return resultado
}

// Restar realiza la resta y la registra
func (c *Calculadora) Restar(a, b float64) float64 {
	resultado := a - b
	c.logger.Registrar("RESTA", fmt.Sprintf("%.2f - %.2f = %.2f", a, b, resultado))
	return resultado
}

// Multiplicar realiza la multiplicación y la registra
func (c *Calculadora) Multiplicar(a, b float64) float64 {
	resultado := a * b
	c.logger.Registrar("MULTIPLICACIÓN", fmt.Sprintf("%.2f × %.2f = %.2f", a, b, resultado))
	return resultado
}

// Dividir realiza la división y la registra (con manejo de errores)
func (c *Calculadora) Dividir(a, b float64) (float64, error) {
	if b == 0 {
		c.logger.Registrar("ERROR", fmt.Sprintf("Intento de división por cero: %.2f ÷ 0", a))
		return 0, fmt.Errorf("no se puede dividir por cero")
	}

	resultado := a / b
	c.logger.Registrar("DIVISIÓN", fmt.Sprintf("%.2f ÷ %.2f = %.2f", a, b, resultado))
	return resultado, nil
}

// MostrarLog lee y muestra el contenido completo del log
func MostrarLog(archivoRuta string) error {
	contenido, err := os.ReadFile(archivoRuta)
	if err != nil {
		return fmt.Errorf("error al leer log: %v", err)
	}

	fmt.Println("\n=== HISTORIAL DE OPERACIONES ===")
	fmt.Println(string(contenido))
	return nil
}

// LimpiarLog borra el contenido del archivo de log
func LimpiarLog(archivoRuta string) error {
	err := os.Truncate(archivoRuta, 0)
	if err != nil {
		return fmt.Errorf("error al limpiar log: %v", err)
	}
	fmt.Println("✓ Log limpiado exitosamente")
	return nil
}

func main() {
	// Creamos el logger
	logger, err := NuevoLogger("operaciones.log")
	if err != nil {
		fmt.Printf("Error al crear logger: %v\n", err)
		return
	}
	defer logger.Cerrar()

	// Creamos la calculadora con logging
	calc := NuevaCalculadora(logger)

	for {
		fmt.Println("\n=== CALCULADORA CON LOGGING ===")
		fmt.Println("1. Sumar")
		fmt.Println("2. Restar")
		fmt.Println("3. Multiplicar")
		fmt.Println("4. Dividir")
		fmt.Println("5. Ver historial de operaciones")
		fmt.Println("6. Limpiar historial")
		fmt.Println("7. Salir")

		var opcion int
		fmt.Print("\nSelecciona una opción: ")
		fmt.Scanln(&opcion)

		if opcion == 7 {
			logger.Registrar("SALIR", "Usuario cerró la aplicación")
			fmt.Println("¡Hasta luego! 👋")
			return
		}

		switch opcion {
		case 1, 2, 3, 4:
			// Operaciones matemáticas
			var a, b float64
			fmt.Print("Primer número: ")
			fmt.Scanln(&a)
			fmt.Print("Segundo número: ")
			fmt.Scanln(&b)

			var resultado float64
			var err error

			switch opcion {
			case 1:
				resultado = calc.Sumar(a, b)
				fmt.Printf("\n✓ Resultado: %.2f + %.2f = %.2f\n", a, b, resultado)

			case 2:
				resultado = calc.Restar(a, b)
				fmt.Printf("\n✓ Resultado: %.2f - %.2f = %.2f\n", a, b, resultado)

			case 3:
				resultado = calc.Multiplicar(a, b)
				fmt.Printf("\n✓ Resultado: %.2f × %.2f = %.2f\n", a, b, resultado)

			case 4:
				resultado, err = calc.Dividir(a, b)
				if err != nil {
					fmt.Printf("\n❌ Error: %v\n", err)
				} else {
					fmt.Printf("\n✓ Resultado: %.2f ÷ %.2f = %.2f\n", a, b, resultado)
				}
			}

		case 5:
			// Ver historial
			if err := MostrarLog("operaciones.log"); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}

		case 6:
			// Limpiar historial
			logger.Registrar("LIMPIAR", "Usuario limpió el historial")
			logger.Cerrar()

			if err := LimpiarLog("operaciones.log"); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
				return
			}

			// Reabrimos el logger
			logger, err = NuevoLogger("operaciones.log")
			if err != nil {
				fmt.Printf("Error al reabrir logger: %v\n", err)
				return
			}
			calc = NuevaCalculadora(logger)

		default:
			logger.Registrar("ERROR", "Opción inválida seleccionada")
			fmt.Println("❌ Opción inválida")
		}
	}
}
