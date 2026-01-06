// Ejercicio 10.1: Agenda con persistencia en JSON
// Guarda y carga contactos desde un archivo JSON para mantener los datos entre ejecuciones.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Contacto representa la información de un contacto
type Contacto struct {
	Nombre   string `json:"nombre"`
	Telefono string `json:"telefono"`
	Email    string `json:"email"`
}

// Agenda almacena una colección de contactos con persistencia
type Agenda struct {
	contactos     []Contacto
	archivoRuta   string
	logOperaciones bool
}

// NuevaAgenda crea una nueva agenda y carga los datos si existen
func NuevaAgenda(archivoRuta string) *Agenda {
	agenda := &Agenda{
		contactos:     []Contacto{},
		archivoRuta:   archivoRuta,
		logOperaciones: true,
	}
	agenda.Cargar()
	return agenda
}

// Guardar escribe los contactos en un archivo JSON
func (a *Agenda) Guardar() error {
	// Convertimos los contactos a JSON con formato legible
	datos, err := json.MarshalIndent(a.contactos, "", "  ")
	if err != nil {
		return fmt.Errorf("error al convertir a JSON: %v", err)
	}

	// Escribimos el archivo
	err = ioutil.WriteFile(a.archivoRuta, datos, 0644)
	if err != nil {
		return fmt.Errorf("error al escribir archivo: %v", err)
	}

	fmt.Printf("💾 Agenda guardada en %s\n", a.archivoRuta)
	return nil
}

// Cargar lee los contactos desde el archivo JSON
func (a *Agenda) Cargar() error {
	// Verificamos si el archivo existe
	if _, err := os.Stat(a.archivoRuta); os.IsNotExist(err) {
		fmt.Println("📄 No se encontró archivo previo. Se creará uno nuevo.")
		return nil
	}

	// Leemos el archivo
	datos, err := ioutil.ReadFile(a.archivoRuta)
	if err != nil {
		return fmt.Errorf("error al leer archivo: %v", err)
	}

	// Convertimos de JSON a la estructura
	err = json.Unmarshal(datos, &a.contactos)
	if err != nil {
		return fmt.Errorf("error al parsear JSON: %v", err)
	}

	fmt.Printf("✓ Se cargaron %d contacto(s) desde %s\n", len(a.contactos), a.archivoRuta)
	return nil
}

// Agregar añade un nuevo contacto y guarda automáticamente
func (a *Agenda) Agregar(c Contacto) error {
	a.contactos = append(a.contactos, c)
	fmt.Printf("✓ Contacto '%s' agregado exitosamente\n", c.Nombre)

	// Guardamos automáticamente después de agregar
	return a.Guardar()
}

// Listar muestra todos los contactos
func (a *Agenda) Listar() {
	if len(a.contactos) == 0 {
		fmt.Println("📭 La agenda está vacía")
		return
	}

	fmt.Println("\n=== LISTA DE CONTACTOS ===")
	for i, c := range a.contactos {
		fmt.Printf("\n[%d] %s\n", i+1, c.Nombre)
		fmt.Printf("    📞 Teléfono: %s\n", c.Telefono)
		fmt.Printf("    📧 Email: %s\n", c.Email)
	}
	fmt.Printf("\nTotal: %d contacto(s)\n", len(a.contactos))
}

// BuscarPorNombre busca contactos por nombre
func (a *Agenda) BuscarPorNombre(nombre string) []Contacto {
	var encontrados []Contacto
	nombreBusqueda := strings.ToLower(nombre)

	for _, c := range a.contactos {
		if strings.Contains(strings.ToLower(c.Nombre), nombreBusqueda) {
			encontrados = append(encontrados, c)
		}
	}
	return encontrados
}

// Eliminar elimina un contacto y guarda automáticamente
func (a *Agenda) Eliminar(indice int) error {
	if indice < 0 || indice >= len(a.contactos) {
		return fmt.Errorf("índice inválido")
	}

	nombreEliminado := a.contactos[indice].Nombre
	a.contactos = append(a.contactos[:indice], a.contactos[indice+1:]...)
	fmt.Printf("✓ Contacto '%s' eliminado exitosamente\n", nombreEliminado)

	// Guardamos automáticamente después de eliminar
	return a.Guardar()
}

// ExportarCSV exporta la agenda a formato CSV
func (a *Agenda) ExportarCSV(archivoCSV string) error {
	var contenido strings.Builder
	contenido.WriteString("Nombre,Teléfono,Email\n")

	for _, c := range a.contactos {
		contenido.WriteString(fmt.Sprintf("%s,%s,%s\n", c.Nombre, c.Telefono, c.Email))
	}

	err := ioutil.WriteFile(archivoCSV, []byte(contenido.String()), 0644)
	if err != nil {
		return fmt.Errorf("error al exportar CSV: %v", err)
	}

	fmt.Printf("✓ Agenda exportada a %s\n", archivoCSV)
	return nil
}

func main() {
	// Creamos la agenda con persistencia
	agenda := NuevaAgenda("contactos.json")

	for {
		fmt.Println("\n=== AGENDA DE CONTACTOS (CON PERSISTENCIA) ===")
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Listar contactos")
		fmt.Println("3. Buscar por nombre")
		fmt.Println("4. Eliminar contacto")
		fmt.Println("5. Exportar a CSV")
		fmt.Println("6. Salir")

		var opcion int
		fmt.Print("\nSelecciona una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			// Agregar contacto
			var nombre, telefono, email string
			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)
			fmt.Print("Teléfono: ")
			fmt.Scanln(&telefono)
			fmt.Print("Email: ")
			fmt.Scanln(&email)

			if err := agenda.Agregar(Contacto{
				Nombre:   nombre,
				Telefono: telefono,
				Email:    email,
			}); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}

		case 2:
			// Listar contactos
			agenda.Listar()

		case 3:
			// Buscar por nombre
			var busqueda string
			fmt.Print("Ingresa el nombre a buscar: ")
			fmt.Scanln(&busqueda)

			encontrados := agenda.BuscarPorNombre(busqueda)
			if len(encontrados) == 0 {
				fmt.Println("❌ No se encontraron contactos")
			} else {
				fmt.Printf("\n✓ Se encontraron %d contacto(s):\n", len(encontrados))
				for i, c := range encontrados {
					fmt.Printf("\n[%d] %s\n", i+1, c.Nombre)
					fmt.Printf("    📞 %s\n", c.Telefono)
					fmt.Printf("    📧 %s\n", c.Email)
				}
			}

		case 4:
			// Eliminar contacto
			agenda.Listar()
			if len(agenda.contactos) > 0 {
				var indice int
				fmt.Print("\nIngresa el número del contacto a eliminar: ")
				fmt.Scanln(&indice)

				if err := agenda.Eliminar(indice - 1); err != nil {
					fmt.Printf("❌ Error: %v\n", err)
				}
			}

		case 5:
			// Exportar a CSV
			if err := agenda.ExportarCSV("contactos.csv"); err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			}

		case 6:
			// Salir
			fmt.Println("¡Hasta luego! 👋")
			return

		default:
			fmt.Println("❌ Opción inválida")
		}
	}
}
