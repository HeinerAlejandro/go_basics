package main

import (
	"fmt"
)

func main() {
	// Declaracion de Variables
	var firstName, lastName string
	var age int

	firstName = "Heiner"
	lastName = "Enis"
	age = 25

	// Tambien se pueden asignar en un unico bloque var

	var (
		firstName2 string = "Alejandro"
		lastName2  string = "Caicedo"
		age2       int    = 25
	)

	// se puede omitir el var por el operador de asignacion := (solo para efectos de declaracion)
	// Se puede omitir tipado cuando se hace asigmacion multiple

	// Hay un diferencia entre Var y :=.
	// 1. Permite declarar variables que estaran en el scope global, es decir, pueden ser accedidas desde fuera y dentro de funciones
	// 2. Permite declarar variables dentro del scope de la funcion (es como un let de javascript)

	charge, business, years := "Something", "Some place", 10

	fmt.Println(
		firstName,
		lastName,
		age,
	)

	fmt.Println(
		firstName2,
		lastName2,
		age2,
	)

	fmt.Println(
		charge,
		business,
		years,
	)

}
