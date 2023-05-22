package main

import (
	"fmt"


	"github.com/robertobouses/struct_ejercicios/user"
)

func main() {
	fmt.Println("programa de struct ejercicios")

	var Persona1 user.Persona
=======
type Persona struct {
	nombre string
	edad   int
}

func (per *Persona) Scanear() {
	fmt.Println("Introduce los datos de la persona, nombre y edad:")
	fmt.Scanln(&per.nombre)
	fmt.Scanln(&per.edad)
}

func (per Persona) Imprimir() {
	fmt.Println("Se llama", per.nombre)
	fmt.Println("Tiene", per.edad, "años")
}

func main() {
	fmt.Println("Viejo")

	var Persona1 Persona

	Persona1.Scanear()
	Persona1.Imprimir()
}
