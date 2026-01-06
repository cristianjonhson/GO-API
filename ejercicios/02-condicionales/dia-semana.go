// Ejercicio 2.3: Día de la semana
// Convierte un número (1-7) en el día de la semana correspondiente usando switch.

package main

import "fmt"

func main() {
	var dia int

	// Solicitamos el número del día
	fmt.Print("Ingresa un número del 1 al 7: ")
	fmt.Scanln(&dia)

	// Usamos switch para determinar el día de la semana
	var nombreDia string

	switch dia {
	case 1:
		nombreDia = "Lunes"
	case 2:
		nombreDia = "Martes"
	case 3:
		nombreDia = "Miércoles"
	case 4:
		nombreDia = "Jueves"
	case 5:
		nombreDia = "Viernes"
	case 6:
		nombreDia = "Sábado"
	case 7:
		nombreDia = "Domingo"
	default:
		// Manejo de valores inválidos
		fmt.Printf("Error: %d no es un día válido. Ingresa un número entre 1 y 7.\n", dia)
		return
	}

	// Imprimimos el resultado
	fmt.Printf("El día %d es: %s\n", dia, nombreDia)

	// Extra: Indicamos si es fin de semana o día laboral
	if dia >= 1 && dia <= 5 {
		fmt.Println("Es un día laboral 💼")
	} else {
		fmt.Println("Es fin de semana 🎉")
	}
}
