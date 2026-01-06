// Tests unitarios para el sistema de gestión de tareas

package main

import (
	"os"
	"testing"
	"time"
)

// TestValidarTitulo prueba la validación de títulos
func TestValidarTitulo(t *testing.T) {
	tests := []struct {
		nombre    string
		titulo    string
		debeErrar bool
	}{
		{"título válido", "Comprar leche", false},
		{"título vacío", "", true},
		{"título con espacios", "   ", true},
		{"título muy corto", "ab", true},
		{"título exacto 3 chars", "abc", false},
		{"título largo válido", "Esta es una tarea con un título bastante largo pero válido", false},
		{"título muy largo", string(make([]byte, 101)), true},
		{"título con 100 chars", string(make([]byte, 100)), false},
	}

	for _, tt := range tests {
		t.Run(tt.nombre, func(t *testing.T) {
			err := ValidarTitulo(tt.titulo)
			
			if tt.debeErrar && err == nil {
				t.Error("Se esperaba un error pero no se obtuvo ninguno")
			}
			
			if !tt.debeErrar && err != nil {
				t.Errorf("No se esperaba error pero se obtuvo: %v", err)
			}
		})
	}
}

// TestCrearTarea prueba la creación de tareas
func TestCrearTarea(t *testing.T) {
	// Creamos un archivo temporal para las pruebas
	archivoTemp := "test_tareas.json"
	defer os.Remove(archivoTemp)

	gestor, err := NuevoGestorTareas(archivoTemp)
	if err != nil {
		t.Fatalf("Error al crear gestor: %v", err)
	}

	// Test: Crear tarea válida
	tarea, err := gestor.Crear("Tarea de prueba")
	if err != nil {
		t.Errorf("Error al crear tarea válida: %v", err)
	}

	if tarea.ID != 1 {
		t.Errorf("ID esperado: 1, obtenido: %d", tarea.ID)
	}

	if tarea.Titulo != "Tarea de prueba" {
		t.Errorf("Título esperado: 'Tarea de prueba', obtenido: '%s'", tarea.Titulo)
	}

	if tarea.Completada {
		t.Error("La tarea nueva debe estar incompleta")
	}

	// Test: Crear tarea inválida (título corto)
	_, err = gestor.Crear("ab")
	if err == nil {
		t.Error("Se esperaba error para título muy corto")
	}

	// Test: Incremento de ID
	tarea2, _ := gestor.Crear("Segunda tarea")
	if tarea2.ID != 2 {
		t.Errorf("ID esperado: 2, obtenido: %d", tarea2.ID)
	}
}

// TestBuscarPorID prueba la búsqueda por ID
func TestBuscarPorID(t *testing.T) {
	archivoTemp := "test_buscar.json"
	defer os.Remove(archivoTemp)

	gestor, _ := NuevoGestorTareas(archivoTemp)
	
	// Creamos algunas tareas
	gestor.Crear("Tarea 1")
	gestor.Crear("Tarea 2")
	gestor.Crear("Tarea 3")

	// Test: Buscar tarea existente
	tarea, err := gestor.BuscarPorID(2)
	if err != nil {
		t.Errorf("Error al buscar tarea existente: %v", err)
	}

	if tarea.Titulo != "Tarea 2" {
		t.Errorf("Título esperado: 'Tarea 2', obtenido: '%s'", tarea.Titulo)
	}

	// Test: Buscar tarea inexistente
	_, err = gestor.BuscarPorID(999)
	if err == nil {
		t.Error("Se esperaba error para ID inexistente")
	}
}

// TestBuscarPorTexto prueba la búsqueda por texto
func TestBuscarPorTexto(t *testing.T) {
	archivoTemp := "test_buscar_texto.json"
	defer os.Remove(archivoTemp)

	gestor, _ := NuevoGestorTareas(archivoTemp)
	
	// Creamos tareas con diferentes títulos
	gestor.Crear("Comprar leche")
	gestor.Crear("Comprar pan")
	gestor.Crear("Estudiar Go")
	gestor.Crear("Leer libro de programación")

	// Test: Búsqueda que encuentra múltiples resultados
	encontradas := gestor.BuscarPorTexto("Comprar")
	if len(encontradas) != 2 {
		t.Errorf("Se esperaban 2 resultados, se obtuvieron: %d", len(encontradas))
	}

	// Test: Búsqueda case-insensitive
	encontradas = gestor.BuscarPorTexto("go")
	if len(encontradas) != 1 {
		t.Errorf("Se esperaba 1 resultado, se obtuvieron: %d", len(encontradas))
	}

	// Test: Búsqueda sin resultados
	encontradas = gestor.BuscarPorTexto("inexistente")
	if len(encontradas) != 0 {
		t.Errorf("Se esperaban 0 resultados, se obtuvieron: %d", len(encontradas))
	}

	// Test: Búsqueda vacía retorna nada
	encontradas = gestor.BuscarPorTexto("")
	if len(encontradas) != 4 {
		t.Errorf("Búsqueda vacía debería retornar todas (4), se obtuvieron: %d", len(encontradas))
	}
}

// TestCompletarTarea prueba completar tareas
func TestCompletarTarea(t *testing.T) {
	archivoTemp := "test_completar.json"
	defer os.Remove(archivoTemp)

	gestor, _ := NuevoGestorTareas(archivoTemp)
	tarea, _ := gestor.Crear("Tarea por completar")

	// Test: Completar tarea existente
	err := gestor.Completar(tarea.ID)
	if err != nil {
		t.Errorf("Error al completar tarea: %v", err)
	}

	// Verificamos que esté completada
	tareaActualizada, _ := gestor.BuscarPorID(tarea.ID)
	if !tareaActualizada.Completada {
		t.Error("La tarea debería estar completada")
	}

	// Test: Completar tarea ya completada
	err = gestor.Completar(tarea.ID)
	if err == nil {
		t.Error("Se esperaba error al completar tarea ya completada")
	}

	// Test: Completar tarea inexistente
	err = gestor.Completar(999)
	if err == nil {
		t.Error("Se esperaba error para tarea inexistente")
	}
}

// TestEliminarTarea prueba eliminar tareas
func TestEliminarTarea(t *testing.T) {
	archivoTemp := "test_eliminar.json"
	defer os.Remove(archivoTemp)

	gestor, _ := NuevoGestorTareas(archivoTemp)
	
	tarea1, _ := gestor.Crear("Tarea 1")
	gestor.Crear("Tarea 2")
	gestor.Crear("Tarea 3")

	// Test: Eliminar tarea existente
	err := gestor.Eliminar(tarea1.ID)
	if err != nil {
		t.Errorf("Error al eliminar tarea: %v", err)
	}

	// Verificamos que ya no existe
	_, err = gestor.BuscarPorID(tarea1.ID)
	if err == nil {
		t.Error("La tarea eliminada no debería existir")
	}

	// Verificamos que quedan 2 tareas
	if len(gestor.tareas) != 2 {
		t.Errorf("Deberían quedar 2 tareas, hay: %d", len(gestor.tareas))
	}

	// Test: Eliminar tarea inexistente
	err = gestor.Eliminar(999)
	if err == nil {
		t.Error("Se esperaba error para tarea inexistente")
	}
}

// TestEstadisticas prueba las estadísticas
func TestEstadisticas(t *testing.T) {
	archivoTemp := "test_stats.json"
	defer os.Remove(archivoTemp)

	gestor, _ := NuevoGestorTareas(archivoTemp)

	// Test: Estadísticas vacías
	total, completadas, pendientes := gestor.Estadisticas()
	if total != 0 || completadas != 0 || pendientes != 0 {
		t.Error("Estadísticas iniciales deberían ser todas 0")
	}

	// Creamos tareas
	t1, _ := gestor.Crear("Tarea 1")
	gestor.Crear("Tarea 2")
	t3, _ := gestor.Crear("Tarea 3")

	// Completamos algunas
	gestor.Completar(t1.ID)
	gestor.Completar(t3.ID)

	// Test: Estadísticas con datos
	total, completadas, pendientes = gestor.Estadisticas()
	
	if total != 3 {
		t.Errorf("Total esperado: 3, obtenido: %d", total)
	}
	if completadas != 2 {
		t.Errorf("Completadas esperadas: 2, obtenidas: %d", completadas)
	}
	if pendientes != 1 {
		t.Errorf("Pendientes esperadas: 1, obtenidas: %d", pendientes)
	}
}

// TestPersistencia prueba guardar y cargar tareas
func TestPersistencia(t *testing.T) {
	archivoTemp := "test_persistencia.json"
	defer os.Remove(archivoTemp)

	// Creamos gestor y tareas
	gestor1, _ := NuevoGestorTareas(archivoTemp)
	gestor1.Crear("Tarea persistente 1")
	gestor1.Crear("Tarea persistente 2")
	gestor1.Completar(1)

	// Guardamos
	err := gestor1.Guardar()
	if err != nil {
		t.Fatalf("Error al guardar: %v", err)
	}

	// Creamos nuevo gestor y cargamos
	gestor2, err := NuevoGestorTareas(archivoTemp)
	if err != nil {
		t.Fatalf("Error al crear segundo gestor: %v", err)
	}

	// Verificamos que se cargaron correctamente
	if len(gestor2.tareas) != 2 {
		t.Errorf("Se esperaban 2 tareas cargadas, se obtuvieron: %d", len(gestor2.tareas))
	}

	// Verificamos que la tarea completada sigue completada
	tarea, _ := gestor2.BuscarPorID(1)
	if !tarea.Completada {
		t.Error("La tarea cargada debería estar completada")
	}

	// Verificamos que el próximo ID es correcto
	nuevaTarea, _ := gestor2.Crear("Nueva tarea")
	if nuevaTarea.ID != 3 {
		t.Errorf("El próximo ID debería ser 3, es: %d", nuevaTarea.ID)
	}
}

// TestListarPendientesYCompletadas prueba los filtros de listado
func TestListarPendientesYCompletadas(t *testing.T) {
	archivoTemp := "test_listar.json"
	defer os.Remove(archivoTemp)

	gestor, _ := NuevoGestorTareas(archivoTemp)
	
	t1, _ := gestor.Crear("Tarea 1")
	gestor.Crear("Tarea 2")
	t3, _ := gestor.Crear("Tarea 3")
	gestor.Crear("Tarea 4")

	// Completamos algunas
	gestor.Completar(t1.ID)
	gestor.Completar(t3.ID)

	// Test: Listar pendientes
	pendientes := gestor.ListarPendientes()
	if len(pendientes) != 2 {
		t.Errorf("Se esperaban 2 pendientes, se obtuvieron: %d", len(pendientes))
	}

	// Test: Listar completadas
	completadas := gestor.ListarCompletadas()
	if len(completadas) != 2 {
		t.Errorf("Se esperaban 2 completadas, se obtuvieron: %d", len(completadas))
	}
}

// TestFechaCreacion verifica que se guarde correctamente la fecha
func TestFechaCreacion(t *testing.T) {
	archivoTemp := "test_fecha.json"
	defer os.Remove(archivoTemp)

	gestor, _ := NuevoGestorTareas(archivoTemp)
	
	antes := time.Now()
	tarea, _ := gestor.Crear("Tarea con fecha")
	despues := time.Now()

	// Verificamos que la fecha esté en el rango correcto
	if tarea.FechaCreacion.Before(antes) || tarea.FechaCreacion.After(despues) {
		t.Error("La fecha de creación no está en el rango esperado")
	}
}

// Benchmark para crear tareas
func BenchmarkCrearTarea(b *testing.B) {
	archivoTemp := "bench_crear.json"
	defer os.Remove(archivoTemp)

	gestor, _ := NuevoGestorTareas(archivoTemp)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gestor.Crear("Tarea de benchmark")
	}
}

// Benchmark para buscar por ID
func BenchmarkBuscarPorID(b *testing.B) {
	archivoTemp := "bench_buscar.json"
	defer os.Remove(archivoTemp)

	gestor, _ := NuevoGestorTareas(archivoTemp)
	
	// Creamos 1000 tareas
	for i := 0; i < 1000; i++ {
		gestor.Crear("Tarea de prueba")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gestor.BuscarPorID(500)
	}
}
