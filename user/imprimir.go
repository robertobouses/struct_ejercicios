package user

import "fmt"

func (per Persona) Imprimir() {
	fmt.Println("Se llama", per.nombre)
	fmt.Println("Tiene", per.edad, "años")
}
