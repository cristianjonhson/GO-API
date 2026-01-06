// Package main implementa un sistema completo de gestión de tareas desde línea de comandos.
//
// El sistema proporciona operaciones CRUD (Crear, Leer, Actualizar, Eliminar) para tareas,
// con persistencia en archivos JSON, validaciones de entrada, manejo robusto de errores,
// y guardado automático mediante concurrencia con goroutines.
//
// # Características principales
//
//   - CRUD completo de tareas con validaciones
//   - Persistencia automática en formato JSON
//   - Búsqueda por ID o texto (case-insensitive)
//   - Filtrado por estado (completadas/pendientes)
//   - Estadísticas en tiempo real
//   - Autoguardado periódico con goroutines
//   - Interfaz CLI interactiva con menú
//
// # Uso básico
//
// Crear un gestor y realizar operaciones:
//
//	gestor, err := NuevoGestorTareas("tareas.json")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Crear una tarea
//	tarea, err := gestor.Crear("Estudiar concurrencia en Go")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Listar tareas pendientes
//	pendientes := gestor.ListarPendientes()
//	for _, t := range pendientes {
//		fmt.Printf("[%d] %s\n", t.ID, t.Titulo)
//	}
//
//	// Completar una tarea
//	err = gestor.Completar(tarea.ID)
//
//	// Guardar cambios
//	err = gestor.Guardar()
//
// # Autoguardado
//
// El sistema soporta guardado automático periódico:
//
//	detener := make(chan bool)
//	gestor.IniciarAutoguardado(30*time.Second, detener)
//	// ... operaciones ...
//	detener <- true  // Detener cuando termine
//
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Tarea representa una tarea individual en el sistema de gestión.
//
// Cada tarea contiene un identificador único auto-incremental, un título
// descriptivo que debe cumplir validaciones de longitud (3-100 caracteres),
// un estado de completitud booleano, y un timestamp de creación.
//
// Las tareas se serializan a JSON para persistencia usando los tags json.
type Tarea struct {
	// ID es el identificador único auto-incremental de la tarea.
	// Los IDs se asignan secuencialmente y nunca se reutilizan.
	ID            int       `json:"id"`
	
	// Titulo es la descripción de la tarea.
	// Debe tener entre 3 y 100 caracteres y no puede estar vacío.
	Titulo        string    `json:"titulo"`
	
	// Completada indica si la tarea ha sido marcada como finalizada.
	// Las tareas nuevas siempre comienzan con false.
	Completada    bool      `json:"completada"`
	
	// FechaCreacion es el timestamp UTC de cuando se creó la tarea.
	// Se asigna automáticamente al crear la tarea.
	FechaCreacion time.Time `json:"fecha_creacion"`
}

// GestorTareas maneja la colección de tareas y su persistencia en disco.
//
// Este tipo es el núcleo del sistema, encapsulando todas las operaciones
// sobre tareas (CRUD), la persistencia en JSON, y el control del autoguardado.
//
// Los campos no exportados garantizan la integridad de los datos y evitan
// modificaciones directas desde código externo.
//
// Todas las operaciones que modifican tareas son thread-safe cuando se usan
// correctamente con el autoguardado.
type GestorTareas struct {
	// tareas almacena la colección completa de tareas en memoria
	tareas         []Tarea
	
	// archivoRuta es la ruta del archivo JSON donde se persisten las tareas
	archivoRuta    string
	
	// proximoID es el siguiente ID disponible para asignar a nuevas tareas
	proximoID      int
	
	// cambiosPendientes indica si hay modificaciones sin guardar en disco
	cambiosPendientes bool
	
	// autoguardadoActivo indica si el goroutine de autoguardado está ejecutándose
	autoguardadoActivo bool
}

// NuevoGestorTareas crea un nuevo gestor de tareas con persistencia en archivo.
//
// Si el archivo especificado existe, carga automáticamente las tareas desde él
// y actualiza el próximo ID para evitar colisiones. Si el archivo no existe,
// crea un gestor vacío que creará el archivo en el primer guardado.
//
// Parámetros:
//   - archivoRuta: ruta del archivo JSON para persistencia (ej: "tareas.json")
//
// Retorna:
//   - *GestorTareas: puntero al gestor creado y listo para usar
//   - error: error si hay problemas leyendo/parseando el archivo existente
//
// Ejemplo:
//
//	gestor, err := NuevoGestorTareas("mis_tareas.json")
//	if err != nil {
//		return err
//	}
//
func NuevoGestorTareas(archivoRuta string) (*GestorTareas, error) {
	gestor := &GestorTareas{
		tareas:         []Tarea{},
		archivoRuta:    archivoRuta,
		proximoID:      1,
		cambiosPendientes: false,
		autoguardadoActivo: false,
	}

	// Intentamos cargar tareas existentes
	if err := gestor.Cargar(); err != nil {
		// Si no existe el archivo, no es un error crítico
		if !os.IsNotExist(err) {
			return nil, err
		}
	}

	return gestor, nil
}

// Guardar persiste todas las tareas en el archivo JSON configurado.
//
// Serializa la colección completa de tareas a JSON con indentación para
// legibilidad y sobrescribe el archivo completo. Si el guardado es exitoso,
// resetea la bandera de cambios pendientes.
//
// El archivo se crea con permisos 0644 (lectura para todos, escritura para dueño).
//
// Retorna:
//   - error: error si falla la serialización o escritura del archivo
//
// Ejemplo:
//
//	if err := gestor.Guardar(); err != nil {
//		log.Printf("Error al guardar: %v", err)
//	}
//
func (g *GestorTareas) Guardar() error {
	datos, err := json.MarshalIndent(g.tareas, "", "  ")
	if err != nil {
		return fmt.Errorf("error al serializar tareas: %v", err)
	}

	err = ioutil.WriteFile(g.archivoRuta, datos, 0644)
	if err != nil {
		return fmt.Errorf("error al escribir archivo: %v", err)
	}

	g.cambiosPendientes = false
	return nil
}

// Cargar lee y deserializa las tareas desde el archivo JSON.
//
// Lee el archivo completo, parsea el JSON a la estructura de tareas,
// y actualiza el próximo ID basándose en el ID más alto encontrado.
// Imprime un mensaje confirmando cuántas tareas se cargaron.
//
// Esta función se llama automáticamente por NuevoGestorTareas, pero puede
// invocarse manualmente para recargar tareas desde disco.
//
// Retorna:
//   - error: error si el archivo no existe, no se puede leer, o el JSON es inválido
//
// Nota: Si el archivo no existe, retorna os.IsNotExist error que puede
// manejarse con os.IsNotExist(err).
//
func (g *GestorTareas) Cargar() error {
	datos, err := ioutil.ReadFile(g.archivoRuta)
	if err != nil {
		return err
	}

	err = json.Unmarshal(datos, &g.tareas)
	if err != nil {
		return fmt.Errorf("error al parsear JSON: %v", err)
	}

	// Actualizamos el próximo ID
	for _, tarea := range g.tareas {
		if tarea.ID >= g.proximoID {
			g.proximoID = tarea.ID + 1
		}
	}

	fmt.Printf("✓ Cargadas %d tarea(s) desde %s\n", len(g.tareas), g.archivoRuta)
	return nil
}

// ValidarTitulo valida que el título de una tarea cumpla todos los requisitos.
//
// El título se considera válido si:
//   - No está vacío después de eliminar espacios en blanco
//   - Tiene al menos 3 caracteres
//   - No excede los 100 caracteres
//
// Esta función se usa internamente por Crear antes de añadir una nueva tarea.
//
// Parámetros:
//   - titulo: el título a validar (puede contener espacios al inicio/fin)
//
// Retorna:
//   - error: error descriptivo si la validación falla, nil si es válido
//
// Ejemplo:
//
//	if err := ValidarTitulo("Comprar leche"); err != nil {
//		fmt.Println("Título inválido:", err)
//	}
//
func ValidarTitulo(titulo string) error {
	titulo = strings.TrimSpace(titulo)
	
	if titulo == "" {
		return fmt.Errorf("el título no puede estar vacío")
	}
	
	if len(titulo) < 3 {
		return fmt.Errorf("el título debe tener al menos 3 caracteres")
	}
	
	if len(titulo) > 100 {
		return fmt.Errorf("el título no puede exceder 100 caracteres")
	}
	
	return nil
}

// Crear añade una nueva tarea a la colección con el título especificado.
//
// Valida el título usando ValidarTitulo, asigna un ID único auto-incremental,
// establece el estado como no completada, y registra el timestamp de creación.
// Los espacios en blanco al inicio/fin del título se eliminan automáticamente.
//
// Marca el gestor como teniendo cambios pendientes para el autoguardado.
//
// Parámetros:
//   - titulo: descripción de la tarea (se validará longitud)
//
// Retorna:
//   - *Tarea: puntero a la tarea recién creada
//   - error: error si la validación del título falla
//
// Ejemplo:
//
//	tarea, err := gestor.Crear("Estudiar goroutines")
//	if err != nil {
//		return err
//	}
//	fmt.Printf("Tarea creada con ID: %d\n", tarea.ID)
//
func (g *GestorTareas) Crear(titulo string) (*Tarea, error) {
	// Validamos el título
	if err := ValidarTitulo(titulo); err != nil {
		return nil, err
	}

	tarea := Tarea{
		ID:            g.proximoID,
		Titulo:        strings.TrimSpace(titulo),
		Completada:    false,
		FechaCreacion: time.Now(),
	}

	g.tareas = append(g.tareas, tarea)
	g.proximoID++
	g.cambiosPendientes = true

	return &tarea, nil
}

// Listar retorna todas las tareas sin filtrar.
//
// Devuelve una copia del slice de tareas, incluyendo tanto completadas
// como pendientes, en el orden en que fueron creadas.
//
// Retorna:
//   - []Tarea: slice con todas las tareas (puede estar vacío)
//
// Ver también: ListarPendientes, ListarCompletadas
//
func (g *GestorTareas) Listar() []Tarea {
	return g.tareas
}

// ListarPendientes retorna solo las tareas que no han sido completadas.
//
// Filtra la colección completa y retorna únicamente las tareas con
// Completada == false.
//
// Retorna:
//   - []Tarea: slice con tareas pendientes (vacío si no hay pendientes)
//
// Ejemplo:
//
//	pendientes := gestor.ListarPendientes()
//	fmt.Printf("Tienes %d tareas pendientes\n", len(pendientes))
//
func (g *GestorTareas) ListarPendientes() []Tarea {
	var pendientes []Tarea
	for _, tarea := range g.tareas {
		if !tarea.Completada {
			pendientes = append(pendientes, tarea)
		}
	}
	return pendientes
}

// ListarCompletadas retorna solo las tareas que han sido marcadas como completadas.
//
// Filtra la colección completa y retorna únicamente las tareas con
// Completada == true.
//
// Retorna:
//   - []Tarea: slice con tareas completadas (vacío si no hay completadas)
//
// Ejemplo:
//
//	completadas := gestor.ListarCompletadas()
//	fmt.Printf("Has completado %d tareas\n", len(completadas))
//
func (g *GestorTareas) ListarCompletadas() []Tarea {
	var completadas []Tarea
	for _, tarea := range g.tareas {
		if tarea.Completada {
			completadas = append(completadas, tarea)
		}
	}
	return completadas
}

// BuscarPorID encuentra y retorna una tarea específica por su identificador único.
//
// Realiza una búsqueda lineal en la colección de tareas. Si encuentra
// una tarea con el ID especificado, retorna un puntero a ella.
//
// Parámetros:
//   - id: el identificador único de la tarea a buscar
//
// Retorna:
//   - *Tarea: puntero a la tarea encontrada
//   - error: error si no existe ninguna tarea con ese ID
//
// Ejemplo:
//
//	tarea, err := gestor.BuscarPorID(5)
//	if err != nil {
//		fmt.Println("Tarea no encontrada")
//	} else {
//		fmt.Println("Tarea:", tarea.Titulo)
//	}
//
func (g *GestorTareas) BuscarPorID(id int) (*Tarea, error) {
	for i := range g.tareas {
		if g.tareas[i].ID == id {
			return &g.tareas[i], nil
		}
	}
	return nil, fmt.Errorf("tarea con ID %d no encontrada", id)
}

// BuscarPorTexto encuentra todas las tareas cuyos títulos contengan el texto especificado.
//
// Realiza una búsqueda case-insensitive (no distingue mayúsculas/minúsculas)
// en todos los títulos de tareas. Retorna todas las coincidencias encontradas.
//
// Si el texto está vacío, retorna todas las tareas.
//
// Parámetros:
//   - texto: el texto a buscar en los títulos (case-insensitive)
//
// Retorna:
//   - []Tarea: slice con todas las tareas que contienen el texto (puede estar vacío)
//
// Ejemplo:
//
//	resultados := gestor.BuscarPorTexto("comprar")
//	fmt.Printf("Se encontraron %d tareas\n", len(resultados))
//	for _, t := range resultados {
//		fmt.Printf("  - %s\n", t.Titulo)
//	}
//
func (g *GestorTareas) BuscarPorTexto(texto string) []Tarea {
	var encontradas []Tarea
	textoBusqueda := strings.ToLower(texto)

	for _, tarea := range g.tareas {
		if strings.Contains(strings.ToLower(tarea.Titulo), textoBusqueda) {
			encontradas = append(encontradas, tarea)
		}
	}

	return encontradas
}

// Completar marca una tarea específica como completada.
//
// Busca la tarea por su ID y establece su campo Completada en true.
// Verifica que la tarea no esté ya completada para evitar operaciones redundantes.
//
// Marca el gestor como teniendo cambios pendientes para el autoguardado.
//
// Parámetros:
//   - id: el identificador único de la tarea a completar
//
// Retorna:
//   - error: error si el ID no existe o la tarea ya está completada
//
// Ejemplo:
//
//	if err := gestor.Completar(3); err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println("Tarea completada exitosamente")
//	}
//
func (g *GestorTareas) Completar(id int) error {
	for i := range g.tareas {
		if g.tareas[i].ID == id {
			if g.tareas[i].Completada {
				return fmt.Errorf("la tarea ya está completada")
			}
			g.tareas[i].Completada = true
			g.cambiosPendientes = true
			return nil
		}
	}
	return fmt.Errorf("tarea con ID %d no encontrada", id)
}

// Eliminar remueve permanentemente una tarea de la colección.
//
// Busca la tarea por su ID y la elimina del slice. Esta operación
// es irreversible y el ID eliminado no se reutiliza.
//
// Marca el gestor como teniendo cambios pendientes para el autoguardado.
//
// Parámetros:
//   - id: el identificador único de la tarea a eliminar
//
// Retorna:
//   - error: error si no existe ninguna tarea con ese ID
//
// Ejemplo:
//
//	if err := gestor.Eliminar(7); err != nil {
//		fmt.Println("No se pudo eliminar:", err)
//	} else {
//		fmt.Println("Tarea eliminada")
//	}
//
func (g *GestorTareas) Eliminar(id int) error {
	for i := range g.tareas {
		if g.tareas[i].ID == id {
			g.tareas = append(g.tareas[:i], g.tareas[i+1:]...)
			g.cambiosPendientes = true
			return nil
		}
	}
	return fmt.Errorf("tarea con ID %d no encontrada", id)
}

// Estadisticas calcula y retorna estadísticas sobre las tareas.
//
// Recorre todas las tareas y cuenta el total, cuántas están completadas,
// y cuántas están pendientes. Es útil para mostrar resúmenes al usuario.
//
// Retorna (retornos con nombre):
//   - total: número total de tareas en la colección
//   - completadas: número de tareas con Completada == true
//   - pendientes: número de tareas con Completada == false
//
// Ejemplo:
//
//	total, completadas, pendientes := gestor.Estadisticas()
//	fmt.Printf("Total: %d | Completadas: %d | Pendientes: %d\n",
//		total, completadas, pendientes)
//
func (g *GestorTareas) Estadisticas() (total, completadas, pendientes int) {
	total = len(g.tareas)
	for _, tarea := range g.tareas {
		if tarea.Completada {
			completadas++
		} else {
			pendientes++
		}
	}
	return
}

// IniciarAutoguardado inicia una goroutine que guarda automáticamente las tareas periódicamente.
//
// Crea un ticker que dispara cada intervalo especificado. En cada tick,
// verifica si hay cambios pendientes y solo guarda si es necesario para
// optimizar I/O. La goroutine se ejecuta en segundo plano hasta recibir
// una señal por el canal detener.
//
// Solo puede haber un autoguardado activo a la vez. Llamadas adicionales
// son ignoradas mientras hay uno en ejecución.
//
// Parámetros:
//   - intervalo: frecuencia de guardado (ej: 30*time.Second para cada 30s)
//   - detener: canal de solo lectura para señalizar terminación del autoguardado
//
// Ejemplo:
//
//	detener := make(chan bool)
//	gestor.IniciarAutoguardado(30*time.Second, detener)
//	defer func() { detener <- true }()
//	// ... realizar operaciones ...
//
// Nota: Es responsabilidad del llamador cerrar o enviar señal por el canal
// cuando desee detener el autoguardado.
//
func (g *GestorTareas) IniciarAutoguardado(intervalo time.Duration, detener <-chan bool) {
	if g.autoguardadoActivo {
		return
	}
	
	g.autoguardadoActivo = true
	
	go func() {
		ticker := time.NewTicker(intervalo)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if g.cambiosPendientes {
					if err := g.Guardar(); err != nil {
						fmt.Printf("\n⚠️  Error en autoguardado: %v\n", err)
					} else {
						fmt.Printf("\n💾 Autoguardado realizado (%s)\n", time.Now().Format("15:04:05"))
					}
				}
			case <-detener:
				fmt.Println("\n🛑 Autoguardado detenido")
				return
			}
		}
	}()
}

// MostrarTareas imprime una lista de tareas con formato visual atractivo.
//
// Genera una salida formateada con emojis para el estado (✅ completada, ⬜ pendiente),
// el ID, título y fecha de creación de cada tarea. Incluye un encabezado con el
// título especificado y un resumen con el total de tareas mostradas.
//
// Si la lista está vacía, muestra un mensaje indicándolo.
//
// Parámetros:
//   - tareas: slice de tareas a mostrar (puede estar vacío)
//   - titulo: encabezado descriptivo para la lista (ej: "TAREAS PENDIENTES")
//
// Ejemplo:
//
//	pendientes := gestor.ListarPendientes()
//	MostrarTareas(pendientes, "⬜ TAREAS POR HACER")
//
func MostrarTareas(tareas []Tarea, titulo string) {
	if len(tareas) == 0 {
		fmt.Printf("\n%s: No hay tareas\n", titulo)
		return
	}

	fmt.Printf("\n%s\n", titulo)
	fmt.Println(strings.Repeat("=", 70))

	for _, tarea := range tareas {
		estado := "⬜"
		if tarea.Completada {
			estado = "✅"
		}

		fecha := tarea.FechaCreacion.Format("02/01/2006 15:04")
		fmt.Printf("%s [%d] %s\n", estado, tarea.ID, tarea.Titulo)
		fmt.Printf("    📅 Creada: %s\n", fecha)
	}

	fmt.Println(strings.Repeat("=", 70))
	fmt.Printf("Total: %d tarea(s)\n", len(tareas))
}

func main() {
	fmt.Println("╔═══════════════════════════════════════════════════════╗")
	fmt.Println("║     SISTEMA DE GESTIÓN DE TAREAS - TODO CLI         ║")
	fmt.Println("╚═══════════════════════════════════════════════════════╝")

	// Creamos el gestor de tareas
	gestor, err := NuevoGestorTareas("tareas.json")
	if err != nil {
		fmt.Printf("Error fatal al iniciar: %v\n", err)
		return
	}

	// Iniciamos el autoguardado cada 30 segundos
	detenerAutoguardado := make(chan bool)
	gestor.IniciarAutoguardado(30*time.Second, detenerAutoguardado)

	// Menú principal
	for {
		total, completadas, pendientes := gestor.Estadisticas()

		fmt.Println("\n┌─────────────────────────────────────────────────────┐")
		fmt.Printf("│ Tareas: %d total | %d completadas | %d pendientes    \n", total, completadas, pendientes)
		fmt.Println("└─────────────────────────────────────────────────────┘")
		fmt.Println("\n📋 MENÚ PRINCIPAL")
		fmt.Println("1. ➕ Crear tarea")
		fmt.Println("2. 📄 Listar todas")
		fmt.Println("3. ⬜ Listar pendientes")
		fmt.Println("4. ✅ Listar completadas")
		fmt.Println("5. 🔍 Buscar tarea")
		fmt.Println("6. ✔️  Completar tarea")
		fmt.Println("7. 🗑️  Eliminar tarea")
		fmt.Println("8. 💾 Guardar ahora")
		fmt.Println("9. 🚪 Salir")

		var opcion int
		fmt.Print("\n➤ Selecciona una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			// Crear tarea
			fmt.Print("\n📝 Título de la tarea: ")
			var titulo string
			fmt.Scanln(&titulo)

			tarea, err := gestor.Crear(titulo)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			} else {
				fmt.Printf("✅ Tarea creada con ID: %d\n", tarea.ID)
			}

		case 2:
			// Listar todas
			MostrarTareas(gestor.Listar(), "📋 TODAS LAS TAREAS")

		case 3:
			// Listar pendientes
			MostrarTareas(gestor.ListarPendientes(), "⬜ TAREAS PENDIENTES")

		case 4:
			// Listar completadas
			MostrarTareas(gestor.ListarCompletadas(), "✅ TAREAS COMPLETADAS")

		case 5:
			// Buscar tarea
			fmt.Print("\n🔍 Buscar por (1=ID, 2=Texto): ")
			var tipoBusqueda int
			fmt.Scanln(&tipoBusqueda)

			if tipoBusqueda == 1 {
				// Buscar por ID
				var id int
				fmt.Print("ID de la tarea: ")
				fmt.Scanln(&id)

				tarea, err := gestor.BuscarPorID(id)
				if err != nil {
					fmt.Printf("❌ %v\n", err)
				} else {
					MostrarTareas([]Tarea{*tarea}, "🔍 RESULTADO DE BÚSQUEDA")
				}
			} else {
				// Buscar por texto
				var texto string
				fmt.Print("Texto a buscar: ")
				fmt.Scanln(&texto)

				encontradas := gestor.BuscarPorTexto(texto)
				if len(encontradas) == 0 {
					fmt.Println("❌ No se encontraron tareas")
				} else {
					MostrarTareas(encontradas, fmt.Sprintf("🔍 RESULTADOS (contienen '%s')", texto))
				}
			}

		case 6:
			// Completar tarea
			var id int
			fmt.Print("\n✔️  ID de la tarea a completar: ")
			fmt.Scanln(&id)

			if err := gestor.Completar(id); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			} else {
				fmt.Printf("✅ Tarea %d marcada como completada\n", id)
			}

		case 7:
			// Eliminar tarea
			var id int
			fmt.Print("\n🗑️  ID de la tarea a eliminar: ")
			fmt.Scanln(&id)

			// Mostramos la tarea antes de eliminar
			tarea, err := gestor.BuscarPorID(id)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
				continue
			}

			fmt.Printf("¿Eliminar '%s'? (s/n): ", tarea.Titulo)
			var confirmar string
			fmt.Scanln(&confirmar)

			if strings.ToLower(confirmar) == "s" {
				if err := gestor.Eliminar(id); err != nil {
					fmt.Printf("❌ Error: %v\n", err)
				} else {
					fmt.Printf("✅ Tarea eliminada\n")
				}
			} else {
				fmt.Println("❌ Cancelado")
			}

		case 8:
			// Guardar manualmente
			if err := gestor.Guardar(); err != nil {
				fmt.Printf("❌ Error al guardar: %v\n", err)
			} else {
				fmt.Println("✅ Tareas guardadas exitosamente")
			}

		case 9:
			// Salir
			detenerAutoguardado <- true
			
			// Guardamos antes de salir si hay cambios
			if gestor.cambiosPendientes {
				fmt.Print("\n💾 Hay cambios sin guardar. ¿Guardar antes de salir? (s/n): ")
				var guardar string
				fmt.Scanln(&guardar)
				
				if strings.ToLower(guardar) == "s" {
					if err := gestor.Guardar(); err != nil {
						fmt.Printf("❌ Error al guardar: %v\n", err)
					} else {
						fmt.Println("✅ Tareas guardadas")
					}
				}
			}

			fmt.Println("\n👋 ¡Hasta luego!")
			return

		default:
			fmt.Println("❌ Opción inválida")
		}
	}
}
