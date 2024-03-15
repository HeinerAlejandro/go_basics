package main

import "fmt"

func Sum(nums ...int) int {

	sum := 0

	for num := range nums {
		sum += num
	}

	return sum
}

func main() {
	sum := Sum(
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

	fmt.Println(sum)
}
