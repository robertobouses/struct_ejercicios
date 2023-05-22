package user

import "fmt"

func (emp *Empleado) ScanearSueldo() {
	emp.Persona.Scanear()
	fmt.Println("indique el sueldo")
	fmt.Scanln(&emp.sueldo)
	fmt.Println("acceso a nombre", emp.Persona.nombre)
	fmt.Println("indica la edad")
	fmt.Scanln(&emp.edad)
	fmt.Println("la edad es", emp.edad)
	fmt.Scanln(&emp.Persona.edad)
	fmt.Println("la edad es", emp.Persona.edad)
}
