package main

import "fmt"

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
}
