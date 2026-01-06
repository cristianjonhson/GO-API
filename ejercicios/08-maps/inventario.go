// Ejercicio 8: Sistema de inventario con Maps
// Gestiona productos y su stock usando un diccionario (map).

package main

import "fmt"

// Inventario es un mapa que relaciona productos con su cantidad en stock
type Inventario map[string]int

// AgregarStock aumenta el stock de un producto
// Si el producto no existe, lo crea con el stock indicado
func (inv Inventario) AgregarStock(producto string, cantidad int) {
	inv[producto] += cantidad
	fmt.Printf("✓ Se agregaron %d unidades de '%s'. Stock actual: %d\n",
		cantidad, producto, inv[producto])
}

// RestarStock disminuye el stock de un producto
// No permite que el stock sea negativo
// Retorna true si la operación fue exitosa, false si no hay suficiente stock
func (inv Inventario) RestarStock(producto string, cantidad int) bool {
	stockActual, existe := inv[producto]

	// Verificamos si el producto existe
	if !existe {
		fmt.Printf("❌ El producto '%s' no existe en el inventario\n", producto)
		return false
	}

	// Verificamos si hay suficiente stock
	if stockActual < cantidad {
		fmt.Printf("❌ Stock insuficiente. Disponible: %d, Solicitado: %d\n",
			stockActual, cantidad)
		return false
	}

	// Restamos el stock
	inv[producto] -= cantidad
	fmt.Printf("✓ Se restaron %d unidades de '%s'. Stock actual: %d\n",
		cantidad, producto, inv[producto])
	return true
}

// ListarInventario muestra todos los productos y su stock
func (inv Inventario) ListarInventario() {
	if len(inv) == 0 {
		fmt.Println("📦 El inventario está vacío")
		return
	}

	fmt.Println("\n=== INVENTARIO COMPLETO ===")
	fmt.Println("Producto\t\tStock")
	fmt.Println("---------------------------")

	for producto, stock := range inv {
		// Indicador visual según nivel de stock
		indicador := "✓"
		if stock == 0 {
			indicador = "⚠️ "
		} else if stock < 10 {
			indicador = "⚡"
		}

		fmt.Printf("%s %-20s\t%d\n", indicador, producto, stock)
	}

	fmt.Printf("\nTotal de productos: %d\n", len(inv))
}

// ConsultarStock muestra el stock de un producto específico
func (inv Inventario) ConsultarStock(producto string) {
	stock, existe := inv[producto]

	if !existe {
		fmt.Printf("❌ El producto '%s' no existe en el inventario\n", producto)
		return
	}

	fmt.Printf("📦 Stock de '%s': %d unidades\n", producto, stock)

	// Alertas de stock
	if stock == 0 {
		fmt.Println("⚠️  ¡SIN STOCK! Necesita reposición urgente")
	} else if stock < 10 {
		fmt.Println("⚡ Stock bajo. Considera reponer pronto")
	}
}

// EliminarProducto elimina un producto del inventario
func (inv Inventario) EliminarProducto(producto string) bool {
	if _, existe := inv[producto]; !existe {
		fmt.Printf("❌ El producto '%s' no existe\n", producto)
		return false
	}

	delete(inv, producto)
	fmt.Printf("✓ Producto '%s' eliminado del inventario\n", producto)
	return true
}

// ProductosBajoStock retorna una lista de productos con stock menor al mínimo
func (inv Inventario) ProductosBajoStock(minimo int) []string {
	var bajoStock []string

	for producto, stock := range inv {
		if stock < minimo {
			bajoStock = append(bajoStock, producto)
		}
	}

	return bajoStock
}

func main() {
	// Creamos el inventario inicial
	inventario := make(Inventario)

	// Agregamos algunos productos de ejemplo
	inventario.AgregarStock("Laptop", 15)
	inventario.AgregarStock("Mouse", 50)
	inventario.AgregarStock("Teclado", 30)
	inventario.AgregarStock("Monitor", 8)

	// Menú interactivo
	for {
		fmt.Println("\n=== SISTEMA DE INVENTARIO ===")
		fmt.Println("1. Agregar stock")
		fmt.Println("2. Restar stock")
		fmt.Println("3. Consultar stock de producto")
		fmt.Println("4. Listar todo el inventario")
		fmt.Println("5. Productos con stock bajo")
		fmt.Println("6. Eliminar producto")
		fmt.Println("7. Salir")

		var opcion int
		fmt.Print("\nSelecciona una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			// Agregar stock
			var producto string
			var cantidad int

			fmt.Print("Nombre del producto: ")
			fmt.Scanln(&producto)
			fmt.Print("Cantidad a agregar: ")
			fmt.Scanln(&cantidad)

			if cantidad < 1 {
				fmt.Println("❌ La cantidad debe ser mayor a 0")
				continue
			}

			inventario.AgregarStock(producto, cantidad)

		case 2:
			// Restar stock
			var producto string
			var cantidad int

			fmt.Print("Nombre del producto: ")
			fmt.Scanln(&producto)
			fmt.Print("Cantidad a restar: ")
			fmt.Scanln(&cantidad)

			if cantidad < 1 {
				fmt.Println("❌ La cantidad debe ser mayor a 0")
				continue
			}

			inventario.RestarStock(producto, cantidad)

		case 3:
			// Consultar stock
			var producto string
			fmt.Print("Nombre del producto: ")
			fmt.Scanln(&producto)

			inventario.ConsultarStock(producto)

		case 4:
			// Listar inventario
			inventario.ListarInventario()

		case 5:
			// Productos con stock bajo
			var minimo int
			fmt.Print("Stock mínimo a considerar: ")
			fmt.Scanln(&minimo)

			bajoStock := inventario.ProductosBajoStock(minimo)

			if len(bajoStock) == 0 {
				fmt.Printf("✓ Todos los productos tienen stock >= %d\n", minimo)
			} else {
				fmt.Printf("\n⚠️  Productos con stock menor a %d:\n", minimo)
				for _, producto := range bajoStock {
					fmt.Printf("  • %s (Stock: %d)\n", producto, inventario[producto])
				}
			}

		case 6:
			// Eliminar producto
			var producto string
			fmt.Print("Nombre del producto a eliminar: ")
			fmt.Scanln(&producto)

			inventario.EliminarProducto(producto)

		case 7:
			fmt.Println("¡Hasta luego! 👋")
			return

		default:
			fmt.Println("❌ Opción inválida")
		}
	}
}
