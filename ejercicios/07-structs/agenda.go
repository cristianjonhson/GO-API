// Ejercicio 7.1: Agenda de contactos
// Sistema CRUD (Create, Read, Update, Delete) en memoria usando structs.

package main

import (
	"fmt"
	"strings"
)

// Contacto representa la información de un contacto
type Contacto struct {
	Nombre   string
	Telefono string
	Email    string
}

// Agenda almacena una colección de contactos
type Agenda struct {
	contactos []Contacto
}

// Agregar añade un nuevo contacto a la agenda
func (a *Agenda) Agregar(c Contacto) {
	a.contactos = append(a.contactos, c)
	fmt.Printf("✓ Contacto '%s' agregado exitosamente\n", c.Nombre)
}

// Listar muestra todos los contactos de la agenda
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

// BuscarPorNombre busca contactos por nombre (búsqueda parcial, ignora mayúsculas)
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

// Eliminar elimina un contacto por su posición en la agenda
func (a *Agenda) Eliminar(indice int) bool {
	if indice < 0 || indice >= len(a.contactos) {
		return false
	}

	nombreEliminado := a.contactos[indice].Nombre
	// Eliminamos el elemento del slice
	a.contactos = append(a.contactos[:indice], a.contactos[indice+1:]...)
	fmt.Printf("✓ Contacto '%s' eliminado exitosamente\n", nombreEliminado)
	return true
}

func main() {
	agenda := Agenda{}

	// Agregamos algunos contactos de ejemplo
	agenda.Agregar(Contacto{
		Nombre:   "Juan Pérez",
		Telefono: "+56912345678",
		Email:    "juan.perez@email.com",
	})
	agenda.Agregar(Contacto{
		Nombre:   "María González",
		Telefono: "+56987654321",
		Email:    "maria.gonzalez@email.com",
	})
	agenda.Agregar(Contacto{
		Nombre:   "Pedro Silva",
		Telefono: "+56955555555",
		Email:    "pedro.silva@email.com",
	})

	// Menú interactivo
	for {
		fmt.Println("\n=== AGENDA DE CONTACTOS ===")
		fmt.Println("1. Agregar contacto")
		fmt.Println("2. Listar contactos")
		fmt.Println("3. Buscar por nombre")
		fmt.Println("4. Eliminar contacto")
		fmt.Println("5. Salir")

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

			agenda.Agregar(Contacto{
				Nombre:   nombre,
				Telefono: telefono,
				Email:    email,
			})

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

				if !agenda.Eliminar(indice - 1) {
					fmt.Println("❌ Índice inválido")
				}
			}

		case 5:
			// Salir
			fmt.Println("¡Hasta luego! 👋")
			return

		default:
			fmt.Println("❌ Opción inválida")
		}
	}
}
