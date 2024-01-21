package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		catetoA float64
		catetoB float64
	)

	fmt.Println("Logintud del cateto a (cm):")
	fmt.Scanln(&catetoA)
	fmt.Println("Logintud del cateto b (cm):")
	fmt.Scanln(&catetoB)

	hipotenusa := math.Sqrt(math.Pow(catetoA, 2) + math.Pow(catetoB, 2))
	//also we can use
	hipotenusa = math.Hypot(catetoA, catetoB)
	area := (catetoA * catetoB) / 2
	perimetro := catetoA + catetoB + hipotenusa

	fmt.Printf("Esta es la hipotenusa %.*f cm \n", 4, hipotenusa)
	fmt.Printf("Este es el area %.*f cm \n", 4, area)
	fmt.Printf("Este es el perimetro %.*f cm \n", 4, perimetro)
}
