package main

import (
	"fmt"
)

type interger int

type Numbers interface {
	~int | ~uint | ~float32 | ~float64
}

func Sum[T Numbers](nums ...T) T {

	var sum T

	for _, num := range nums {
		sum += num
	}

	return sum
}

func Includes[T comparable](list []T, value T) bool {
	for _, value := range list {
		if value == value {
			return true
		}
	}

	return false
}

func main() {
	sum_ints := Sum[int](
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
	)

	fmt.Println(sum_ints)

	sum_floats := Sum[float32](
		1.5,
		2.5,
		5.8,
		7.7,
		9.9,
	)

	fmt.Println(sum_floats)

	sum_custom_integers := Sum[interger](
		interger(1),
		interger(2),
		interger(10),
	)

	fmt.Println(sum_custom_integers)

	sum_interface_numbers := Sum[uint](
		5,
		3,
		40,
	)

	fmt.Println(sum_interface_numbers)

	strings := []string{
		"Hola",
		"Estoy",
		"usando",
		"Golang",
		"esta",
		"chevere",
	}

	ints := []int{
		90, 65, 34, 888, 21,
	}

	isInclude1 := Includes[string](strings, "Golang")
	isInclude2 := Includes[int](ints, 34)

	fmt.Printf("Estoy imprimiendo 2 resultados de la funcion include: %b y %b\n", isInclude1, isInclude2)

}
