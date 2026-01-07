# 00 - Introducción a Go

Esta carpeta contiene una guía completa de los fundamentos de Go, organizada en archivos independientes.

## 📚 Contenido

### Archivos Principales

1. **01-syntax-basico.go** - Sintaxis fundamental
   - Estructura de programas
   - Declaración de variables
   - Constantes
   - Comentarios

2. **02-tipos-datos.go** - Tipos de datos
   - Booleanos
   - Enteros (int8, int16, int32, int64, uint, etc.)
   - Flotantes (float32, float64)
   - Complex numbers
   - Strings
   - Arrays
   - Slices
   - Maps
   - Punteros
   - Interfaces
   - Valores cero

3. **03-conversiones-tipos.go** - Conversiones y casteo
   - Conversiones numéricas
   - String ↔ Número
   - Byte/Rune ↔ String
   - Type assertions
   - Type switch

4. **04-funciones.go** - Funciones
   - Funciones básicas
   - Parámetros y retornos
   - Múltiples retornos
   - Retornos nombrados
   - Funciones variádicas
   - First-class functions
   - Closures
   - Defer
   - Métodos (receivers)

5. **05-structs.go** - Estructuras
   - Definición de structs
   - Campos exportados/no exportados
   - Structs anidados
   - Embedding (composición)
   - Métodos asociados
   - Constructores
   - Tags
   - Comparación

6. **06-interfaces.go** - Interfaces
   - Definición de interfaces
   - Implementación implícita
   - Interface vacía (interface{})
   - Type assertions
   - Type switch
   - Interfaces múltiples
   - Interfaces estándar (Stringer, Error)

7. **07-paquetes-modulos.go** - Paquetes y Módulos
   - Concepto de paquete
   - Módulos (go.mod)
   - Estructura de proyectos
   - Imports
   - Visibilidad (exportación)
   - Comandos go mod
   - Mejores prácticas

8. **08-control-flujo.go** - Control de flujo
   - if/else
   - switch
   - for (único loop)
   - for range
   - break/continue
   - goto
   - defer

9. **09-errores.go** - Manejo de errores
   - Error básico
   - Error formateado
   - Errores personalizados
   - Múltiples errores
   - Panic y Recover
   - Error wrapping
   - Mejores prácticas

10. **10-concurrencia.go** - Concurrencia
    - Goroutines
    - Channels
    - Buffered channels
    - Select
    - WaitGroup
    - Mutex
    - Patrones comunes

11. **11-testing.go** - Testing
    - Tests básicos
    - Table-driven tests
    - Subtests
    - Benchmarks
    - Examples
    - Mocking
    - Cobertura

## 🚀 Cómo usar estos archivos

### Ejecutar un archivo específico:
```bash
cd ejercicios/00-introduccion
go run 01-syntax-basico.go
```

### Ejecutar todos los archivos:
```bash
for file in *.go; do
    echo "=== Ejecutando $file ==="
    go run "$file"
    echo ""
done
```

### Ver el contenido sin ejecutar:
```bash
cat 02-tipos-datos.go
```

## 📖 Orden recomendado de estudio

1. **Fundamentos básicos:**
   - 01-syntax-basico.go
   - 02-tipos-datos.go
   - 03-conversiones-tipos.go

2. **Estructuras de código:**
   - 04-funciones.go
   - 05-structs.go
   - 06-interfaces.go

3. **Organización:**
   - 07-paquetes-modulos.go

4. **Control y errores:**
   - 08-control-flujo.go
   - 09-errores.go

5. **Avanzado:**
   - 10-concurrencia.go
   - 11-testing.go

## 💡 Consejos

- **Lee los comentarios:** Cada archivo tiene explicaciones detalladas
- **Ejecuta el código:** Modifica y experimenta con los ejemplos
- **Practica:** Crea tus propias variaciones de los ejemplos
- **Consulta:** Usa `go doc` para ver documentación oficial

## 🔗 Recursos adicionales

- [Go Tour](https://go.dev/tour/) - Tutorial interactivo oficial
- [Go by Example](https://gobyexample.com/) - Ejemplos prácticos
- [Effective Go](https://go.dev/doc/effective_go) - Guía de mejores prácticas
- [Go Playground](https://go.dev/play/) - Ejecuta código en el navegador

## ✅ Verificación de aprendizaje

Después de estudiar estos archivos, deberías poder:

- [ ] Declarar variables y constantes
- [ ] Usar todos los tipos de datos básicos
- [ ] Crear y usar funciones
- [ ] Definir structs e interfaces
- [ ] Organizar código en paquetes
- [ ] Manejar errores correctamente
- [ ] Usar goroutines y channels
- [ ] Escribir tests básicos

---

**¡Buena suerte en tu viaje aprendiendo Go!** 🚀