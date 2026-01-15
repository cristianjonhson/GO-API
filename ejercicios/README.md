# Ejercicios de Go para Principiantes 🚀

Colección de ejercicios prácticos para aprender los fundamentos de programación en Go.

## 📚 Contenido

### 1. Variables, Tipos y Entrada/Salida
- **perfil.go** - Pide nombre, edad y ciudad, imprime resumen formateado
- **conversor.go** - Convierte °C→°F, km→millas, CLP→USD
- **calculadora.go** - Calculadora básica con validación de división por cero

### 2. Condicionales (if / switch)
- **par-impar.go** - Clasifica números como par/impar y positivo/negativo
- **clasificador-notas.go** - Clasifica notas de 1.0-7.0 (Reprobado/Aprobado/Excelente)
- **dia-semana.go** - Convierte números 1-7 en días de la semana

### 3. Bucles (for)
- **tabla-multiplicar.go** - Genera tabla de multiplicar del 1 al 10
- **suma-hasta-cero.go** - Lee números hasta que ingrese 0, calcula suma y promedio
- **adivina-numero.go** - Juego de adivinar número con pistas

### 4. Strings
- **vocales-consonantes.go** - Cuenta vocales y consonantes en un texto
- **palindromo.go** - Verifica si una palabra/frase es palíndromo
- **frecuencia-palabras.go** - Cuenta frecuencia de cada palabra en un texto

### 5. Arrays/Slices y Algoritmos
- **estadisticas.go** - Calcula máximo, mínimo y promedio de números
- **ordenamiento.go** - Implementa Bubble Sort
- **busqueda.go** - Implementa búsqueda lineal y binaria

### 6. Funciones (Modularidad)
- **refactorizacion.go** - Refactoriza ejercicios usando funciones reutilizables
- **validador.go** - Sistema de validación con funciones de rango

### 7. Structs y Métodos
- **agenda.go** - Agenda de contactos con CRUD completo
- **carrito.go** - Carrito de compras con productos, totales y descuentos

### 8. Maps (Diccionarios)
- **inventario.go** - Sistema de inventario con gestión de stock

### 9. Manejo de Errores
- **parser.go** - Parser seguro con conversión de strings a números
- **divisor.go** - Calculadora con manejo robusto de errores

### 10. Archivos (Persistencia)
- **agenda-json.go** - Agenda de contactos con persistencia en JSON
- **logging.go** - Sistema de logging de operaciones con timestamps

### 11. Testing (Confiabilidad)
- **funciones.go** - Funciones puras para testing
- **funciones_test.go** - Tests unitarios con casos borde y benchmarks

### 12. Concurrencia (Goroutines + Channels)
- **descarga-simulada.go** - Simulación de descargas concurrentes
- **procesador-numeros.go** - Pipeline producer-workers-consumer
- **pool-workers.go** - Pool de workers procesando tareas

### 13. Introducción al paquete os
- **13-os.go** - Aprende a interactuar con el sistema operativo: leer variables de entorno, listar archivos y crear archivos temporales.

## 🎯 Objetivos de Aprendizaje

Estos ejercicios te ayudarán a practicar:
- ✅ Variables y tipos de datos
- ✅ Control de flujo (if, switch, for)
- ✅ Entrada/salida de datos
- ✅ Manejo de strings
- ✅ Arrays y slices
- ✅ Funciones y modularidad
- ✅ Structs y métodos
- ✅ Maps (diccionarios)
- ✅ Manejo de errores
- ✅ Persistencia de datos (archivos)
- ✅ Testing unitario
- ✅ Concurrencia (goroutines y channels)
- ✅ Algoritmos básicos
- ✅ Validación de datos
- ✅ Sistemas CRUD
- ✅ Patrones de diseño (producer-consumer, worker pool)

## 🏃 Cómo Ejecutar

Navega a la carpeta del ejercicio que quieras ejecutar:

```bash
# Ejemplo: Ejecutar el ejercicio de perfil
cd ejercicios/01-variables-tipos
go run perfil.go

# Ejemplo: Ejecutar el juego de adivinar número
cd ejercicios/03-bucles
go run adivina-numero.go

# Ejemplo: Ejecutar tests unitarios
cd ejercicios/11-testing
go test -v

# Ejemplo: Ejecutar con coverage
go test -v -cover

# Ejemplo: Ejecutar benchmarks
go test -bench=.
```

## 💡 Consejos

1. **Lee los comentarios** - Cada archivo tiene comentarios explicativos detallados
2. **Experimenta** - Modifica el código y observa los cambios
3. **Progresivo** - Sigue el orden sugerido para mejor comprensión
4. **Práctica** - Intenta resolver cada ejercicio antes de ver la solución
5. **Testing** - Ejecuta los tests para validar tu código
6. **Concurrencia** - Los ejercicios de goroutines son avanzados, tómate tu tiempo

## 🎓 Niveles de Dificultad

- **Básico** (01-05): Variables, condicionales, bucles, strings, arrays
- **Intermedio** (06-09): Funciones, structs, maps, manejo de errores
- **Avanzado** (10-12): Persistencia, testing, concurrencia

## 📖 Estructura del Proyecto

```
ejercicios/
├── 01-variables-tipos/
│   ├── perfil.go
│   ├── conversor.go
│   └── calculadora.go
├── 02-condicionales/
│   ├── par-impar.go
│   ├── clasificador-notas.go
│   └── dia-semana.go
├── 03-bucles/
│   ├── tabla-multiplicar.go
│   ├── suma-hasta-cero.go
│   └── adivina-numero.go
├── 04-strings/
│   ├── vocales-consonantes.go
│   ├── palindromo.go
│   └── frecuencia-palabras.go
├── 05-arrays-slices/
│   ├── estadisticas.go
│   ├── ordenamiento.go
│   └── busqueda.go
├── 06-funciones/
│   ├── refactorizacion.go
│   └── validador.go
├── 07-structs/
│   ├── agenda.go
│   └── carrito.go
├── 08-maps/
│   └── inventario.go
├── 09-errores/
│   ├── parser.go
│   └── divisor.go
├── 10-archivos/
│   ├── agenda-json.go
│   └── logging.go
├── 11-testing/
│   ├── funciones.go
│   └── funciones_test.go
├── 12-concurrencia/
│   ├── descarga-simulada.go
│   ├── procesador-numeros.go
│   └── pool-workers.go
└── 13-os/
    └── 13-os.go
```

## 🎓 Recursos Adicionales

- [Documentación oficial de Go](https://golang.org/doc/)
- [Tour de Go](https://tour.golang.org/)
- [Go by Example](https://gobyexample.com/)

## 🤝 Contribuir

¡Siéntete libre de agregar más ejercicios o mejorar los existentes!

---

**Happy Coding!** 💻✨
