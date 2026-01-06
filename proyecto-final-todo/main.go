// Proyecto Final: Sistema de Gestión de Tareas (ToDo CLI)
// Integra todos los conceptos: structs, persistencia, errores, testing, concurrencia

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Tarea representa una tarea individual en el sistema
type Tarea struct {
	ID            int       `json:"id"`
	Titulo        string    `json:"titulo"`
	Completada    bool      `json:"completada"`
	FechaCreacion time.Time `json:"fecha_creacion"`
}

// GestorTareas maneja la colección de tareas y su persistencia
type GestorTareas struct {
	tareas         []Tarea
	archivoRuta    string
	proximoID      int
	cambiosPendientes bool
	autoguardadoActivo bool
}

// NuevoGestorTareas crea un nuevo gestor de tareas
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

// Guardar escribe las tareas en el archivo JSON
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

// Cargar lee las tareas desde el archivo JSON
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

// ValidarTitulo valida que el título de una tarea sea correcto
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

// Crear añade una nueva tarea
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

// Listar retorna todas las tareas
func (g *GestorTareas) Listar() []Tarea {
	return g.tareas
}

// ListarPendientes retorna solo las tareas no completadas
func (g *GestorTareas) ListarPendientes() []Tarea {
	var pendientes []Tarea
	for _, tarea := range g.tareas {
		if !tarea.Completada {
			pendientes = append(pendientes, tarea)
		}
	}
	return pendientes
}

// ListarCompletadas retorna solo las tareas completadas
func (g *GestorTareas) ListarCompletadas() []Tarea {
	var completadas []Tarea
	for _, tarea := range g.tareas {
		if tarea.Completada {
			completadas = append(completadas, tarea)
		}
	}
	return completadas
}

// BuscarPorID encuentra una tarea por su ID
func (g *GestorTareas) BuscarPorID(id int) (*Tarea, error) {
	for i := range g.tareas {
		if g.tareas[i].ID == id {
			return &g.tareas[i], nil
		}
	}
	return nil, fmt.Errorf("tarea con ID %d no encontrada", id)
}

// BuscarPorTexto encuentra tareas que contengan el texto en su título
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

// Completar marca una tarea como completada
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

// Eliminar remueve una tarea de la lista
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

// Estadisticas retorna estadísticas de las tareas
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

// IniciarAutoguardado inicia un goroutine que guarda automáticamente cada X segundos
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

// MostrarTareas imprime una lista de tareas formateada
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
