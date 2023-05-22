package main

import (
	"fmt"

	"github.com/robertobouses/struct_ejercicios/user"
)

func main() {
	fmt.Println("programa de struct ejercicios")

	var Persona1 user.Persona
	Persona1.Scanear()
	Persona1.Imprimir()
}
