package main

import (
	"fmt"
	"time"
)

func main() {

	for_func(
		"Esto es con continue %d\n",
		15,
		true,
	)

	for_func(
		"Esto es con continue %d\n",
		21,
		false,
	)

	is_morning := check_if_morning(time.Now().Hour())

	fmt.Printf("Es de maniana?: %v\n\n", is_morning)

	sum, mult := sum_mult_values(12, 75)

	fmt.Printf("Sum: %d\nMult: %d\n\n", sum, mult)

	sum_2, mult_2 := sum_mult_values_return_var(10, 21)
	fmt.Printf("Sum 2: %d\nMult 2: %d\n\n", sum_2, mult_2)
}

func for_func(prefix string, upto int, is_continue bool) {

	for index := 0; index < upto; index++ {
		if index == int(upto/2) {
			if is_continue {
				continue
			} else {
				break
			}
		} else {
			fmt.Printf(prefix, index)
		}

	}
}

func check_if_morning(hour int) bool {
	return hour < 12
}

func sum_mult_values(a, b int) (int, int) { // we can also return variables: (sum, mult int)
	sum := a + b
	mult := a * b

	return sum, mult
}

func sum_mult_values_return_var(a, b int) (sum int, mult int) {
	sum = a + b
	mult = a * b

	return // Here de function knows how many and which variables to return based on de named declaration
}
