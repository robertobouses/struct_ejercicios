package user

import "fmt"

func (per *Persona) Scanear() {
	fmt.Println("Introduce los datos de la persona, nombre y edad:")
	fmt.Println("nombre")
	fmt.Scanln(&per.nombre)
	fmt.Println("edad")
	fmt.Scanln(&per.edad)
}
