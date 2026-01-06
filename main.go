package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	// Rutas
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/api/hello", helloHandler)

	// Puerto del servidor
	port := ":8080"
	fmt.Printf("🚀 Servidor corriendo en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// Handler para la ruta principal
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Bienvenido a la API de Go",
		Status:  "success",
	}
	json.NewEncoder(w).Encode(response)
}

// Handler para verificar el estado de la API
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "API funcionando correctamente",
		Status:  "healthy",
	}
	json.NewEncoder(w).Encode(response)
}

// Handler para saludo personalizado
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "Mundo"
	}
	response := Response{
		Message: fmt.Sprintf("¡Hola, %s!", name),
		Status:  "success",
	}
	json.NewEncoder(w).Encode(response)
}
