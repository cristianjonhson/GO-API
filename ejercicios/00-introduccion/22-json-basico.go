// Este programa muestra cómo trabajar con JSON en Go.
// Incluye ejemplos de serialización (convertir estructuras a JSON) y deserialización (convertir JSON a estructuras).

package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func main() {
	// Serializar a JSON
	persona := Persona{Nombre: "Juan", Edad: 30}
	jsonData, _ := json.Marshal(persona)
	fmt.Println("JSON:", string(jsonData))

	// Deserializar desde JSON
	jsonStr := `{"nombre":"Ana","edad":25}`
	var nuevaPersona Persona
	json.Unmarshal([]byte(jsonStr), &nuevaPersona)
	fmt.Println("Persona deserializada:", nuevaPersona)
}