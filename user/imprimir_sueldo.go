package user

import "fmt"

func (emp Empleado) ImprimirSueldo() {
	fmt.Println("el sueldo es:")
	fmt.Println(emp.sueldo)

}
