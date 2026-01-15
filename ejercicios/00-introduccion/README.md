# 00 - Introducción a Go

Esta carpeta contiene una guía completa de los fundamentos de Go, organizada en archivos independientes que puedes ejecutar y modificar para aprender.

## 📚 Contenido Detallado

### 1. **01-syntax-basico.go** - Sintaxis Fundamental de Go
**¿Qué aprenderás?**
- Estructura básica de un programa Go
- Declaración de variables (var, :=, múltiples formas)
- Constantes y bloques de constantes
- Comentarios de una línea y multilínea
- Ámbito de variables (scope)

**¿Para qué sirve?**
- Entender cómo está organizado un programa en Go
- Aprender las diferentes formas de declarar variables
- Conocer las convenciones de código en Go
- Base fundamental antes de avanzar a temas complejos

---

### 2. **02-tipos-datos.go** - Sistema de Tipos Completo
**¿Qué aprenderás?**
- Tipos booleanos (bool)
- Enteros con signo (int8, int16, int32, int64, int)
- Enteros sin signo (uint8, uint16, uint32, uint64, uint)
- Punto flotante (float32, float64)
- Números complejos (complex64, complex128)
- Strings y operaciones con texto
- Arrays (tamaño fijo) y Slices (tamaño dinámico)
- Maps (diccionarios/hash tables)
- Punteros y referencias de memoria
- Interface{} (cualquier tipo)
- Valores cero de cada tipo

**¿Para qué sirve?**
- Elegir el tipo de dato correcto para cada situación
- Entender las diferencias entre arrays y slices
- Trabajar con colecciones de datos (slices, maps)
- Optimizar memoria usando tipos apropiados
- Manejar punteros para modificar valores por referencia

---

### 3. **03-conversiones-tipos.go** - Conversiones y Casteo
**¿Qué aprenderás?**
- Conversiones entre tipos numéricos (int ↔ float)
- String a número (strconv.Atoi, ParseInt, ParseFloat)
- Número a string (strconv.Itoa, FormatInt)
- Conversiones byte/rune ↔ string
- Type assertions (verificar tipos en interfaces)
- Type switch (diferentes acciones según el tipo)

**¿Para qué sirve?**
- Convertir datos de entrada del usuario (strings a números)
- Trabajar con APIs que retornan diferentes tipos
- Validar tipos en tiempo de ejecución
- Evitar errores de tipo en operaciones

---

### 4. **04-funciones.go** - Funciones y Métodos
**¿Qué aprenderás?**
- Funciones básicas y con parámetros
- Funciones con retorno simple y múltiple
- Retornos nombrados
- Funciones variádicas (número variable de argumentos)
- Funciones como valores (first-class functions)
- Closures (funciones anónimas)
- Defer (ejecución diferida)
- Métodos con receivers (valor y puntero)

**¿Para qué sirve?**
- Organizar código en bloques reutilizables
- Manejar múltiples valores de retorno (ej: resultado + error)
- Crear funciones flexibles con parámetros variables
- Implementar callbacks y funciones de orden superior
- Asociar comportamiento a tipos personalizados
- Garantizar limpieza de recursos con defer

---

### 5. **05-structs.go** - Estructuras de Datos
**¿Qué aprenderás?**
- Definición y creación de structs
- Campos exportados vs no exportados
- Structs anidados
- Embedding (composición sin herencia)
- Métodos asociados a structs
- Constructores (patrón factory)
- Tags para metadatos (JSON, DB)
- Comparación de structs

**¿Para qué sirve?**
- Modelar entidades del mundo real (Usuario, Producto, etc.)
- Agrupar datos relacionados
- Crear tipos de datos personalizados
- Implementar composición en lugar de herencia
- Serializar/deserializar datos (JSON, XML)
- Mapear structs a bases de datos

---

### 6. **06-interfaces.go** - Interfaces y Polimorfismo
**¿Qué aprenderás?**
- Definición de interfaces
- Implementación implícita (duck typing)
- Interface vacía (interface{} / any)
- Type assertions seguras
- Type switch para múltiples tipos
- Interfaces múltiples y composición
- Interfaces estándar (Stringer, Error, Reader, Writer)

**¿Para qué sirve?**
- Crear código desacoplado y testeable
- Implementar polimorfismo
- Definir contratos de comportamiento
- Trabajar con tipos genéricos
- Facilitar mocking en tests
- Seguir principios SOLID (D - Dependency Inversion)

---

### 7. **07-paquetes-modulos.go** - Organización del Código
**¿Qué aprenderás?**
- Concepto de paquete (package)
- Sistema de módulos (go.mod, go.sum)
- Estructura de proyectos recomendada
- Imports y aliases
- Visibilidad (mayúscula = público, minúscula = privado)
- Comandos go mod (init, tidy, download)
- Mejores prácticas de organización

**¿Para qué sirve?**
- Organizar proyectos grandes en módulos
- Gestionar dependencias externas
- Crear bibliotecas reutilizables
- Controlar la visibilidad del código
- Versionar y publicar paquetes
- Trabajar en equipo con estructura clara

---

### 8. **08-control-flujo.go** - Estructuras de Control
**¿Qué aprenderás?**
- if/else (con y sin inicialización)
- switch (básico, múltiples casos, sin condición)
- for loops (clásico, while, infinito)
- for range (iteración sobre colecciones)
- break y continue
- Labels para loops anidados
- defer (ejecución al final)

**¿Para qué sirve?**
- Controlar el flujo de ejecución del programa
- Tomar decisiones basadas en condiciones
- Iterar sobre colecciones de datos
- Implementar lógica compleja de negocio
- Optimizar código evitando repetición
- Manejar casos especiales en loops

---

### 9. **09-errores.go** - Manejo de Errores
**¿Qué aprenderás?**
- Patrón básico (if err != nil)
- Crear errores con errors.New()
- Errores formateados con fmt.Errorf()
- Errores personalizados (implementar Error())
- Wrapping de errores (%w)
- Panic y recover (casos excepcionales)
- Verificación de errores (errors.Is, errors.As)

**¿Para qué sirve?**
- Manejar situaciones de error de forma explícita
- Crear mensajes de error descriptivos
- Propagar errores con contexto adicional
- Recuperarse de fallos críticos
- Debugging y logging efectivo
- Construir aplicaciones robustas y confiables

---

### 10. **10-concurrencia.go** - Programación Concurrente
**¿Qué aprenderás?**
- Goroutines (funciones concurrentes)
- Channels (comunicación entre goroutines)
- Buffered channels (canales con buffer)
- Select (multiplexing de channels)
- WaitGroup (esperar a múltiples goroutines)
- Mutex (sincronización de acceso)
- Patrones comunes (worker pool, fan-out/fan-in)

**¿Para qué sirve?**
- Ejecutar múltiples tareas simultáneamente
- Aprovechar múltiples núcleos del CPU
- Mejorar rendimiento de aplicaciones
- Procesar datos en paralelo
- Implementar servidores concurrentes
- Manejar operaciones I/O de forma eficiente

---

### 11. **11-testing.go** - Testing y Calidad de Código
**¿Qué aprenderás?**
- Tests básicos (TestXxx)
- Table-driven tests (múltiples casos)
- Subtests (organización de tests)
- Benchmarks (medición de rendimiento)
- Examples (documentación ejecutable)
- Helpers y setup/teardown
- Cobertura de código
- Mocking de dependencias

**¿Para qué sirve?**
- Verificar que el código funciona correctamente
- Prevenir regresiones (bugs que vuelven)
- Documentar comportamiento esperado
- Medir y optimizar rendimiento
- Facilitar refactoring con confianza
- Desarrollo guiado por tests (TDD)

---

### 12. **12-context.go** - Context y Cancelación ⭐
**¿Qué aprenderás?**
- Context básico (Background, TODO)
- WithCancel (cancelación manual)
- WithTimeout (timeout automático)
- WithDeadline (deadline absoluto)
- WithValue (propagación de valores)
- Propagación de context en cadena
- Worker pools con context
- Mejores prácticas

**¿Para qué sirve?**
- Cancelar operaciones largas o innecesarias
- Implementar timeouts en requests HTTP
- Propagar deadlines a través de llamadas
- Pasar valores request-scoped (userID, traceID)
- Coordinar shutdown graceful de servicios
- Evitar goroutines zombies
- Manejo eficiente de recursos

---

### 13. **13-os.go** - Introducción al paquete os
**¿Qué aprenderás?**
- Leer variables de entorno
- Listar archivos en un directorio
- Crear y manejar archivos temporales

**¿Para qué sirve?**
- Entender cómo interactuar con el sistema operativo desde Go
- Aprender a manejar archivos y directorios
- Trabajar con variables de entorno para configuraciones dinámicas

---

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

### Modificar y experimentar:
```bash
# Abre el archivo en tu editor favorito
code 04-funciones.go

# Modifica el código y ejecuta
go run 04-funciones.go
```

---

## 📖 Orden recomendado de estudio

### 🟢 Nivel Principiante (Días 1-3)
1. **01-syntax-basico.go** - Empieza aquí siempre
2. **02-tipos-datos.go** - Tipos fundamentales
3. **03-conversiones-tipos.go** - Trabajar con tipos
4. **08-control-flujo.go** - if, for, switch

### 🟡 Nivel Intermedio (Días 4-6)
5. **04-funciones.go** - Organizar código
6. **05-structs.go** - Estructuras de datos
7. **06-interfaces.go** - Abstracciones
8. **07-paquetes-modulos.go** - Organización de proyectos
9. **09-errores.go** - Manejo robusto de errores

### 🔴 Nivel Avanzado (Días 7-10)
10. **10-concurrencia.go** - Programación paralela
11. **12-context.go** - Cancelación y timeouts
12. **11-testing.go** - Calidad de código

---

## 💡 Consejos de estudio

### Para cada archivo:
1. **Lee primero** todos los comentarios
2. **Ejecuta** el código sin modificar
3. **Modifica** valores y observa cambios
4. **Experimenta** creando tus propios ejemplos
5. **Rompe** el código intencionalmente para entender errores

### Ejercicios sugeridos:
- **01-syntax-basico.go**: Crea variables de diferentes formas
- **02-tipos-datos.go**: Implementa una calculadora simple
- **04-funciones.go**: Crea tus propias funciones helper
- **05-structs.go**: Modela tu propia entidad (Coche, Casa, etc.)
- **10-concurrencia.go**: Implementa un worker pool personalizado

---

## 🔗 Recursos adicionales

- [Go Tour](https://go.dev/tour/) - Tutorial interactivo oficial
- [Go by Example](https://gobyexample.com/) - Ejemplos prácticos
- [Effective Go](https://go.dev/doc/effective_go) - Guía de mejores prácticas
- [Go Playground](https://go.dev/play/) - Ejecuta código en el navegador
- [Context Package](https://pkg.go.dev/context) - Documentación oficial de context
- [Go Standard Library](https://pkg.go.dev/std) - Biblioteca estándar completa

---

## ✅ Verificación de aprendizaje

### Después de completar estos archivos, deberías poder:

#### Fundamentos
- [ ] Declarar variables de múltiples formas
- [ ] Usar todos los tipos de datos básicos
- [ ] Convertir entre tipos de forma segura
- [ ] Crear y usar constantes

#### Estructuras de código
- [ ] Escribir funciones con múltiples retornos
- [ ] Definir structs y métodos
- [ ] Implementar interfaces implícitamente
- [ ] Organizar código en paquetes

#### Control y errores
- [ ] Usar if, for, switch correctamente
- [ ] Manejar errores de forma explícita
- [ ] Crear errores personalizados
- [ ] Usar defer para cleanup

#### Avanzado
- [ ] Lanzar goroutines y usar channels
- [ ] Implementar cancelación con context
- [ ] Escribir tests básicos y table-driven tests
- [ ] Medir rendimiento con benchmarks

---

## 🎯 Proyecto final sugerido

Después de completar todos los archivos, intenta crear:

**Sistema de procesamiento de pedidos concurrente:**
- Structs para Pedido, Cliente, Producto
- Interfaces para Procesador, Notificador
- Goroutines para procesar múltiples pedidos
- Context para cancelación y timeouts
- Tests completos con mocks
- Manejo robusto de errores

---

**¡Buena suerte en tu viaje aprendiendo Go!** 🚀

💡 **Tip Pro**: Go es un lenguaje simple pero poderoso. No intentes aplicar patrones de otros lenguajes. Abraza la simplicidad de Go.