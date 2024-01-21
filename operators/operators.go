package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	potencia := int(math.Pow(4, 6))

	fmt.Printf("Esta es la potencia 4 ala 6: %s.\n", strconv.Itoa(potencia))
	fmt.Printf("Este es PI: %f.\n", math.Pi)
	fmt.Printf("Este es Euler: %f.\n", math.E)
	fmt.Printf("Otra constante: %f.\n", math.Log10E)

	fmt.Println("Algunas constantes mas:")
	fmt.Printf(
		"|%5s     |     %.4f     |\n|%5s     |     %.4f     |\n|%5s     |     %.4f     |\n",
		"Phi", math.Phi,
		"SQRt2", math.Sqrt2,
		"SQRtPhi", math.SqrtPhi,
	)
}
