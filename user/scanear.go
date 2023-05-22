package user

import "fmt"

func (per *Persona) Scanear() {
	fmt.Println("Introduce los datos de la persona, nombre y edad:")
	fmt.Scanln(&per.nombre)
	fmt.Scanln(&per.edad)
}
