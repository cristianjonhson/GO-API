// Package main contiene el punto de entrada de nuestra API REST
package main

// Importamos las librerías necesarias
import (
	"encoding/json" // Para codificar/decodificar JSON
	"fmt"           // Para formatear strings
	"log"           // Para registrar errores
	"net/http"      // Para crear el servidor HTTP
)

// Response define la estructura estándar de respuesta de la API
// Los tags `json` indican cómo se serializarán los campos en JSON
type Response struct {
	Message string `json:"message"` // Mensaje de respuesta
	Status  string `json:"status"`  // Estado de la operación
}

// main es el punto de entrada de la aplicación
func main() {
	// Registramos los manejadores (handlers) para cada ruta
	// Cada ruta se asocia con una función que procesará las peticiones
	http.HandleFunc("/", homeHandler)           // Ruta raíz
	http.HandleFunc("/api/health", healthHandler) // Verificación de salud
	http.HandleFunc("/api/hello", helloHandler)   // Saludo personalizado

	// Definimos el puerto donde escuchará el servidor
	port := ":8080"
	fmt.Printf("🚀 Servidor corriendo en http://localhost%s\n", port)
	
	// Iniciamos el servidor HTTP
	// log.Fatal registrará cualquier error y detendrá el programa si falla
	log.Fatal(http.ListenAndServe(port, nil))
}

// homeHandler maneja las peticiones a la ruta principal "/"
// w: ResponseWriter para escribir la respuesta HTTP
// r: Request contiene los datos de la petición entrante
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Establecemos el tipo de contenido como JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Creamos la estructura de respuesta
	response := Response{
		Message: "Bienvenido a la API de Go",
		Status:  "success",
	}
	
	// Codificamos la respuesta en JSON y la enviamos al cliente
	json.NewEncoder(w).Encode(response)
}

// healthHandler verifica el estado de la API
// Útil para monitoreo y health checks en producción
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// Establecemos la cabecera de respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Respuesta indicando que la API está funcionando
	response := Response{
		Message: "API funcionando correctamente",
		Status:  "healthy",
	}
	
	// Enviamos la respuesta en formato JSON
	json.NewEncoder(w).Encode(response)
}

// helloHandler genera un saludo personalizado
// Acepta un parámetro "name" en la URL query string
// Ejemplo: /api/hello?name=Juan
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Establecemos el tipo de contenido de la respuesta
	w.Header().Set("Content-Type", "application/json")
	
	// Extraemos el parámetro "name" de la URL
	// Si no se proporciona, usamos "Mundo" como valor predeterminado
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Mundo"
	}
	
	// Creamos una respuesta personalizada con el nombre
	response := Response{
		Message: fmt.Sprintf("¡Hola, %s!", name),
		Status:  "success",
	}
	
	// Convertimos la respuesta a JSON y la enviamos
	json.NewEncoder(w).Encode(response)
}
