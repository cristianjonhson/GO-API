// Ejemplo avanzado de interfaces en Go
package main

import "fmt"

type Forma interface {
	Area() float64
}

type Circulo struct {
	Radio float64
}

type Rectangulo struct {
	Ancho, Alto float64
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Radio * c.Radio
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}
}

func imprimirArea(f Forma) {
	fmt.Println("Área:", f.Area())
}

func main() {
	c := Circulo{Radio: 5}
	r := Rectangulo{Ancho: 4, Alto: 6}

	imprimirArea(c)
	imprimirArea(r)
}