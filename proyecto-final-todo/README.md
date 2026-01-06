# 📝 Sistema de Gestión de Tareas (ToDo CLI)

## Descripción

Proyecto final que integra todos los conceptos aprendidos en el curso de Go. Es un sistema completo de gestión de tareas desde la línea de comandos con persistencia en JSON, validaciones, manejo de errores, tests unitarios y concurrencia.

## 🎯 Características

### Funcionalidades Principales
- ✅ **CRUD completo**: Crear, listar, completar, buscar y eliminar tareas
- 💾 **Persistencia en JSON**: Las tareas se guardan automáticamente en archivo
- 🔍 **Búsqueda avanzada**: Por ID o por texto (case-insensitive)
- 📊 **Estadísticas**: Total, completadas y pendientes
- ✨ **Autoguardado**: Goroutine que guarda cambios cada 30 segundos
- ✅ **Validaciones**: Títulos de 3-100 caracteres
- 🛡️ **Manejo de errores**: Control robusto de errores en todas las operaciones
- 🧪 **Tests completos**: Suite de tests unitarios y benchmarks

### Conceptos Aplicados
- **Structs**: `Tarea` y `GestorTareas` con métodos
- **Concurrencia**: Goroutine con ticker para autoguardado
- **Channels**: Canal para señal de cierre del autoguardado
- **Sync**: Mutex para sincronizar acceso a tareas
- **Encoding/JSON**: Marshal/Unmarshal para persistencia
- **Testing**: Tests unitarios con tabla de tests y benchmarks
- **Manejo de errores**: Retornos de error en todas las operaciones críticas
- **Time**: Gestión de fechas y timers

## 📦 Instalación

```bash
cd proyecto-final-todo
go build -o todo
```

## 🚀 Uso

### Ejecutar la aplicación
```bash
./todo
```

### Menú interactivo
```
=== GESTOR DE TAREAS ===
1. Crear tarea
2. Listar todas
3. Listar pendientes
4. Listar completadas
5. Buscar por ID
6. Buscar por texto
7. Completar tarea
8. Eliminar tarea
9. Ver estadísticas
0. Salir
```

### Ejemplos de uso

**Crear una tarea:**
```
Selecciona una opción: 1
Título de la tarea: Estudiar concurrencia en Go
✓ Tarea creada con ID: 1
```

**Listar tareas:**
```
Selecciona una opción: 2
ID: 1 | Título: Estudiar concurrencia en Go
       Estado: [ ] Pendiente
       Creada: 2024-01-15 10:30:00
```

**Buscar por texto:**
```
Selecciona una opción: 6
Texto a buscar: Go
Se encontraron 3 tareas
ID: 1 | Título: Estudiar concurrencia en Go
...
```

**Completar tarea:**
```
Selecciona una opción: 7
ID de la tarea: 1
✓ Tarea completada
```

**Ver estadísticas:**
```
Selecciona una opción: 9
📊 Estadísticas:
   Total: 5 tareas
   ✓ Completadas: 2
   ⏳ Pendientes: 3
```

## 🧪 Tests

### Ejecutar todos los tests
```bash
go test -v
```

### Ver cobertura
```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

### Ejecutar benchmarks
```bash
go test -bench=.
go test -bench=. -benchmem
```

### Tests incluidos
- `TestValidarTitulo`: Validación de títulos (vacíos, cortos, largos, válidos)
- `TestCrearTarea`: Creación con validación e incremento de IDs
- `TestBuscarPorID`: Búsqueda existente e inexistente
- `TestBuscarPorTexto`: Búsqueda case-insensitive y múltiples resultados
- `TestCompletarTarea`: Completar tareas y validar estados
- `TestEliminarTarea`: Eliminación y verificación
- `TestEstadisticas`: Cálculo de totales, completadas y pendientes
- `TestPersistencia`: Guardar y cargar desde JSON
- `TestListarPendientesYCompletadas`: Filtros de listado
- `TestFechaCreacion`: Verificación de timestamps
- `BenchmarkCrearTarea`: Rendimiento de creación
- `BenchmarkBuscarPorID`: Rendimiento de búsqueda

## 📁 Estructura del Código

### Struct Tarea
```go
type Tarea struct {
    ID            int       // Identificador único
    Titulo        string    // Descripción de la tarea
    Completada    bool      // Estado (completada/pendiente)
    FechaCreacion time.Time // Timestamp de creación
}
```

### Struct GestorTareas
```go
type GestorTareas struct {
    tareas      []Tarea      // Slice de tareas
    proximoID   int          // Contador para IDs únicos
    archivo     string       // Ruta del archivo JSON
    mu          sync.Mutex   // Mutex para concurrencia
    cerrarAuto  chan bool    // Canal para cerrar autoguardado
}
```

### Métodos Principales

| Método | Descripción |
|--------|-------------|
| `ValidarTitulo(titulo string) error` | Valida longitud del título (3-100 chars) |
| `NuevoGestorTareas(archivo string) (*GestorTareas, error)` | Constructor, carga tareas si existen |
| `Crear(titulo string) (*Tarea, error)` | Crea nueva tarea con validación |
| `Listar() []Tarea` | Retorna todas las tareas |
| `ListarPendientes() []Tarea` | Filtra tareas pendientes |
| `ListarCompletadas() []Tarea` | Filtra tareas completadas |
| `BuscarPorID(id int) (*Tarea, error)` | Búsqueda por ID exacto |
| `BuscarPorTexto(texto string) []Tarea` | Búsqueda case-insensitive en títulos |
| `Completar(id int) error` | Marca tarea como completada |
| `Eliminar(id int) error` | Elimina tarea por ID |
| `Estadisticas() (int, int, int)` | Retorna total, completadas, pendientes |
| `Guardar() error` | Persiste tareas en JSON |
| `Cargar() error` | Carga tareas desde JSON |
| `IniciarAutoguardado(intervalo time.Duration)` | Goroutine para guardado automático |
| `DetenerAutoguardado()` | Detiene el autoguardado |

## 🔧 Detalles Técnicos

### Persistencia
Las tareas se guardan en `tareas.json` con el siguiente formato:
```json
[
  {
    "ID": 1,
    "Titulo": "Estudiar concurrencia en Go",
    "Completada": false,
    "FechaCreacion": "2024-01-15T10:30:00Z"
  }
]
```

### Concurrencia
El autoguardado se implementa con:
- **Goroutine**: Ejecuta en segundo plano
- **time.Ticker**: Dispara cada 30 segundos
- **Channel**: Señal para terminar el guardado
- **sync.Mutex**: Protege el slice de tareas

```go
func (g *GestorTareas) IniciarAutoguardado(intervalo time.Duration) {
    ticker := time.NewTicker(intervalo)
    go func() {
        for {
            select {
            case <-ticker.C:
                g.Guardar()
            case <-g.cerrarAuto:
                ticker.Stop()
                return
            }
        }
    }()
}
```

### Validaciones
- **Título vacío**: Error
- **Título < 3 caracteres**: Error
- **Título > 100 caracteres**: Error
- **Tarea ya completada**: Error al intentar completar nuevamente
- **ID inexistente**: Error en búsqueda, completar o eliminar

## 📚 Aprendizajes

Este proyecto demuestra:
1. ✅ **Estructuras de datos** complejas con structs anidados
2. ✅ **Métodos** asociados a structs
3. ✅ **Punteros** para modificar estado
4. ✅ **Slices** para colecciones dinámicas
5. ✅ **Maps** implícitos en búsquedas
6. ✅ **Manejo de errores** con retornos múltiples
7. ✅ **Persistencia** con JSON marshaling
8. ✅ **Concurrencia** segura con mutex
9. ✅ **Goroutines y channels** para tareas en segundo plano
10. ✅ **Testing** completo con cobertura

## 🎓 Conceptos Avanzados

- **Table-driven tests**: Patrón idiomático de Go para tests parametrizados
- **Benchmarks**: Medición de rendimiento
- **Defer**: Limpieza de archivos temporales en tests
- **Select**: Multiplexación de canales en autoguardado
- **Mutex**: Sincronización de acceso concurrente
- **JSON tags**: Serialización personalizada (si se necesitara)

## 🔜 Posibles Mejoras

- [ ] Prioridades para tareas (alta, media, baja)
- [ ] Fechas de vencimiento
- [ ] Categorías o etiquetas
- [ ] Exportar a CSV
- [ ] Interfaz web con net/http
- [ ] Base de datos SQLite en lugar de JSON
- [ ] Ordenamiento personalizado
- [ ] Historial de cambios (log)
- [ ] Recordatorios con notificaciones

## 📝 Licencia

Proyecto educativo - Curso de Go

---

**Autor**: [Tu nombre]  
**Fecha**: Enero 2024  
**Go Version**: 1.25+
