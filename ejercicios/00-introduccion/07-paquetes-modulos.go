// 07. PAQUETES Y MÓDULOS
// Organización y modularidad del código en Go

package main

import (
    "fmt"
    "math"
    "math/rand"
    "strings"
    "time"
)

/*
==========================================
CONCEPTOS FUNDAMENTALES
==========================================

1. PAQUETE (package):
   - Unidad básica de organización del código
   - Todos los archivos en el mismo directorio deben tener el mismo package
   - package main: ejecutable (debe tener func main())
   - Otros nombres: bibliotecas/librerías

2. MÓDULO (module):
   - Colección de paquetes relacionados
   - Definido por go.mod (en la raíz del proyecto)
   - Gestiona dependencias y versiones
   - Identificador único: module path (ej: github.com/usuario/proyecto)

==========================================
ESTRUCTURA DE DIRECTORIOS
==========================================

mi-proyecto/
├── go.mod                    # Define el módulo
├── go.sum                    # Checksums de dependencias
├── main.go                   # package main (ejecutable)
├── utils/
│   ├── helpers.go           # package utils
│   └── validators.go        # package utils
├── models/
│   ├── user.go              # package models
│   └── product.go           # package models
└── internal/                 # Privado al módulo
    └── config/
        └── config.go        # package config

==========================================
VISIBILIDAD / EXPORTACIÓN
==========================================

En Go NO existen palabras clave public/private.
La visibilidad se controla con la primera letra:

- Mayúscula inicial: EXPORTADO (público)
  func SumarNumeros()    ✅ Accesible desde otros paquetes
  type Usuario struct    ✅ Exportado
  const MaxSize          ✅ Exportado

- Minúscula inicial: NO EXPORTADO (privado al paquete)
  func calcularTotal()   ❌ Solo accesible dentro del paquete
  type datosInternos     ❌ Privado
  const limiteInterno    ❌ Privado

==========================================
IMPORTS
==========================================

// Import simple
import "fmt"

// Múltiples imports
import (
    "fmt"
    "strings"
    "math"
)

// Import con alias
import m "math"

// Import para efectos secundarios (init)
import _ "image/png"

// Import de submódulo
import "mi-proyecto/utils"

// Import punto (importa al namespace actual - NO RECOMENDADO)
import . "fmt"

==========================================
COMANDOS GO MOD
==========================================

go mod init <module-path>    # Crear nuevo módulo
go mod tidy                  # Limpiar dependencias no usadas
go mod download              # Descargar dependencias
go mod verify                # Verificar integridad
go mod vendor                # Copiar dependencias a vendor/
go get <package>             # Añadir/actualizar dependencia
go list -m all               # Listar todas las dependencias

==========================================
EJEMPLO DE go.mod
==========================================

module github.com/usuario/mi-proyecto

go 1.21

require (
    github.com/gorilla/mux v1.8.0
    github.com/joho/godotenv v1.5.1
)

==========================================
FUNCIÓN init()
==========================================

Se ejecuta automáticamente antes de main()
Útil para inicialización de paquetes
*/

// init se ejecuta automáticamente al importar el paquete
func init() {
    fmt.Println("⚡ Ejecutando init()...")
    rand.Seed(time.Now().UnixNano())
}

func main() {
    fmt.Println("\n=== PAQUETES Y MÓDULOS EN GO ===\n")

    // ==========================================
    // USO DE PAQUETES ESTÁNDAR
    // ==========================================
    fmt.Println("1. Paquete fmt (formateo):")
    nombre := "Go"
    version := 1.21
    fmt.Printf("Lenguaje: %s, Versión: %.2f\n\n", nombre, version)

    fmt.Println("2. Paquete strings (manipulación de texto):")
    texto := "Hola, Mundo de Go"
    fmt.Printf("Mayúsculas: %s\n", strings.ToUpper(texto))
    fmt.Printf("¿Contiene 'Go'? %v\n", strings.Contains(texto, "Go"))
    fmt.Printf("Reemplazar: %s\n\n", strings.Replace(texto, "Go", "Golang", 1))

    fmt.Println("3. Paquete math (matemáticas):")
    fmt.Printf("Pi: %.5f\n", math.Pi)
    fmt.Printf("Raíz cuadrada de 16: %.0f\n", math.Sqrt(16))
    fmt.Printf("Potencia 2^10: %.0f\n\n", math.Pow(2, 10))

    fmt.Println("4. Paquete time (fecha y hora):")
    ahora := time.Now()
    fmt.Printf("Fecha actual: %s\n", ahora.Format("02/01/2006 15:04:05"))
    fmt.Printf("Unix timestamp: %d\n\n", ahora.Unix())

    // ==========================================
    // ORGANIZACIÓN RECOMENDADA
    // ==========================================
    fmt.Println("=== ESTRUCTURA DE PROYECTO RECOMENDADA ===\n")
    fmt.Println(`
mi-proyecto/
├── go.mod                    # Módulo principal
├── go.sum                    # Checksums
├── cmd/                      # Ejecutables
│   ├── api/
│   │   └── main.go          # package main
│   └── cli/
│       └── main.go          # package main
├── internal/                 # Código privado al módulo
│   ├── auth/
│   ├── database/
│   └── handlers/
├── pkg/                      # Código reutilizable público
│   ├── models/
│   └── utils/
├── api/                      # Definiciones API
├── web/                      # Assets web
├── configs/                  # Archivos de configuración
├── scripts/                  # Scripts de automatización
├── test/                     # Tests adicionales
└── docs/                     # Documentación
    `)

    // ==========================================
    // MEJORES PRÁCTICAS
    // ==========================================
    fmt.Println("\n=== MEJORES PRÁCTICAS ===\n")
    fmt.Println(`
✅ DO:
  • Usar nombres de paquete descriptivos y cortos
  • Un paquete = una responsabilidad
  • Exportar solo lo necesario
  • Usar internal/ para código privado
  • Documentar funciones exportadas
  • go mod tidy regularmente

❌ DON'T:
  • Nombres de paquete genéricos (util, common, misc)
  • Paquetes con muchas responsabilidades
  • Importar paquetes no usados
  • Dependencias circulares
  • Import punto (import . "fmt")
    `)

    // ==========================================
    // EJEMPLO DE DOCUMENTACIÓN
    // ==========================================
    fmt.Println("\n=== DOCUMENTACIÓN DE PAQUETES ===\n")
    fmt.Println(`
// Package utils proporciona utilidades comunes para el proyecto.
//
// Este paquete incluye funciones helper para validación,
// formateo y manipulación de datos.
package utils

// Validar verifica si un email es válido.
//
// Retorna true si el email tiene formato correcto,
// false en caso contrario.
//
// Ejemplo:
//   if utils.ValidarEmail("test@example.com") {
//       // email válido
//   }
func ValidarEmail(email string) bool {
    // implementación...
}
    `)

    fmt.Println("\n💡 Usa 'go doc <paquete>' para ver la documentación")
    fmt.Println("💡 Usa 'godoc -http=:6060' para servidor de docs local")
}