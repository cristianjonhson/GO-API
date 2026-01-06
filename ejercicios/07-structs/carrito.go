// Ejercicio 7.2: Carrito de compras
// Sistema de carrito con productos, cantidades, totales y descuentos.

package main

import "fmt"

// Producto representa un artículo en la tienda
type Producto struct {
	Nombre string
	Precio float64
}

// Item representa un producto con su cantidad en el carrito
type Item struct {
	Producto Producto
	Cantidad int
}

// Carrito almacena los items de compra
type Carrito struct {
	items []Item
}

// AgregarItem añade un producto al carrito con su cantidad
func (c *Carrito) AgregarItem(producto Producto, cantidad int) {
	// Verificamos si el producto ya existe en el carrito
	for i, item := range c.items {
		if item.Producto.Nombre == producto.Nombre {
			// Si existe, aumentamos la cantidad
			c.items[i].Cantidad += cantidad
			fmt.Printf("✓ Se agregaron %d unidades de '%s'\n", cantidad, producto.Nombre)
			return
		}
	}

	// Si no existe, lo agregamos como nuevo item
	c.items = append(c.items, Item{
		Producto: producto,
		Cantidad: cantidad,
	})
	fmt.Printf("✓ '%s' agregado al carrito (%d unidades)\n", producto.Nombre, cantidad)
}

// Subtotal calcula el precio de un item específico (precio x cantidad)
func (c *Carrito) Subtotal(item Item) float64 {
	return item.Producto.Precio * float64(item.Cantidad)
}

// Total calcula el precio total de todos los items en el carrito
func (c *Carrito) Total() float64 {
	total := 0.0
	for _, item := range c.items {
		total += c.Subtotal(item)
	}
	return total
}

// AplicarDescuento calcula el total con descuento si supera cierto monto
// Retorna el total con descuento aplicado y el porcentaje de descuento
func (c *Carrito) AplicarDescuento(montoMinimo float64, porcentajeDescuento float64) (float64, float64) {
	total := c.Total()

	// Si el total supera el monto mínimo, aplicamos descuento
	if total >= montoMinimo {
		descuento := total * (porcentajeDescuento / 100)
		return total - descuento, porcentajeDescuento
	}

	return total, 0
}

// MostrarCarrito imprime todos los items del carrito con detalles
func (c *Carrito) MostrarCarrito() {
	if len(c.items) == 0 {
		fmt.Println("🛒 El carrito está vacío")
		return
	}

	fmt.Println("\n=== CARRITO DE COMPRAS ===")
	for i, item := range c.items {
		subtotal := c.Subtotal(item)
		fmt.Printf("\n[%d] %s\n", i+1, item.Producto.Nombre)
		fmt.Printf("    Precio unitario: $%.2f\n", item.Producto.Precio)
		fmt.Printf("    Cantidad: %d\n", item.Cantidad)
		fmt.Printf("    Subtotal: $%.2f\n", subtotal)
	}
}

// EliminarItem elimina un item del carrito por su índice
func (c *Carrito) EliminarItem(indice int) bool {
	if indice < 0 || indice >= len(c.items) {
		return false
	}

	nombreEliminado := c.items[indice].Producto.Nombre
	c.items = append(c.items[:indice], c.items[indice+1:]...)
	fmt.Printf("✓ '%s' eliminado del carrito\n", nombreEliminado)
	return true
}

func main() {
	carrito := Carrito{}

	// Catálogo de productos disponibles
	productos := []Producto{
		{Nombre: "Laptop", Precio: 599990.00},
		{Nombre: "Mouse", Precio: 15990.00},
		{Nombre: "Teclado", Precio: 45990.00},
		{Nombre: "Monitor", Precio: 199990.00},
		{Nombre: "Audífonos", Precio: 89990.00},
	}

	for {
		fmt.Println("\n=== TIENDA ONLINE ===")
		fmt.Println("1. Ver catálogo")
		fmt.Println("2. Agregar al carrito")
		fmt.Println("3. Ver carrito")
		fmt.Println("4. Eliminar del carrito")
		fmt.Println("5. Finalizar compra")
		fmt.Println("6. Salir")

		var opcion int
		fmt.Print("\nSelecciona una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			// Mostrar catálogo
			fmt.Println("\n=== CATÁLOGO DE PRODUCTOS ===")
			for i, p := range productos {
				fmt.Printf("[%d] %s - $%.2f\n", i+1, p.Nombre, p.Precio)
			}

		case 2:
			// Agregar al carrito
			fmt.Println("\n=== CATÁLOGO ===")
			for i, p := range productos {
				fmt.Printf("[%d] %s - $%.2f\n", i+1, p.Nombre, p.Precio)
			}

			var indice, cantidad int
			fmt.Print("\nNúmero de producto: ")
			fmt.Scanln(&indice)

			if indice < 1 || indice > len(productos) {
				fmt.Println("❌ Producto inválido")
				continue
			}

			fmt.Print("Cantidad: ")
			fmt.Scanln(&cantidad)

			if cantidad < 1 {
				fmt.Println("❌ Cantidad inválida")
				continue
			}

			carrito.AgregarItem(productos[indice-1], cantidad)

		case 3:
			// Ver carrito
			carrito.MostrarCarrito()
			if len(carrito.items) > 0 {
				fmt.Printf("\n💰 Total: $%.2f\n", carrito.Total())
			}

		case 4:
			// Eliminar del carrito
			carrito.MostrarCarrito()
			if len(carrito.items) > 0 {
				var indice int
				fmt.Print("\nNúmero de item a eliminar: ")
				fmt.Scanln(&indice)

				if !carrito.EliminarItem(indice - 1) {
					fmt.Println("❌ Índice inválido")
				}
			}

		case 5:
			// Finalizar compra
			if len(carrito.items) == 0 {
				fmt.Println("❌ El carrito está vacío")
				continue
			}

			carrito.MostrarCarrito()

			// Aplicamos descuento si el total supera $100,000
			totalConDescuento, porcentaje := carrito.AplicarDescuento(100000, 10)

			fmt.Println("\n=== RESUMEN DE COMPRA ===")
			fmt.Printf("Subtotal: $%.2f\n", carrito.Total())

			if porcentaje > 0 {
				fmt.Printf("Descuento (%.0f%%): -$%.2f\n", porcentaje, carrito.Total()-totalConDescuento)
				fmt.Println("🎉 ¡Felicidades! Obtuviste descuento por compra sobre $100,000")
			}

			fmt.Printf("TOTAL A PAGAR: $%.2f\n", totalConDescuento)
			fmt.Println("\n✓ ¡Gracias por tu compra! 🛍️")
			return

		case 6:
			fmt.Println("¡Hasta luego! 👋")
			return

		default:
			fmt.Println("❌ Opción inválida")
		}
	}
}
