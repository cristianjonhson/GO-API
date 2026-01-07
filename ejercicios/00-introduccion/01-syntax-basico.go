// 01. SINTAXIS BÁSICA DE GO
// Este archivo explica la estructura fundamental de un programa en Go

package main // Todo programa ejecutable debe estar en el package main

import "fmt" // Importamos paquetes con import

// FUNCIÓN MAIN: Punto de entrada del programa
func main() {
    // Los comentarios de una línea usan //
    /* Los comentarios 
       multilínea usan /* y */ */

    fmt.Println("Hola, Go!") // Las sentencias NO necesitan ; al final

    // DECLARACIÓN DE VARIABLES
    // Forma 1: var nombre tipo
    var mensaje string
    mensaje = "Hola"

    // Forma 2: var nombre tipo = valor
    var numero int = 42

    // Forma 3: var nombre = valor (tipo inferido)
    var pi = 3.14

    // Forma 4: := (declaración corta, solo dentro de funciones)
    edad := 25

    // Múltiples declaraciones
    var x, y, z int = 1, 2, 3
    a, b := "Go", "Golang"

    fmt.Println(mensaje, numero, pi, edad, x, y, z, a, b)

    // CONSTANTES: valores inmutables
    const PI = 3.14159
    const (
        Lunes = 1
        Martes = 2
        Miercoles = 3
    )

    // BLOQUES DE CÓDIGO: se definen con { }
    {
        // Este es un bloque interno
        variable := "local"
        fmt.Println(variable)
    }
    // 'variable' no existe fuera del bloque

    // ESTRUCTURA GENERAL DE UN ARCHIVO .go:
    // 1. Declaración de package
    // 2. Imports
    // 3. Constantes globales
    // 4. Variables globales
    // 5. Funciones (main va primero por convención)
}