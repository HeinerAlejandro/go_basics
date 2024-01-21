package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 10; i++ {

		if i == 6 {
			continue
		}
		fmt.Println(i)
	}

	for j := 1; j <= 10; j++ {

		if j == 6 {
			break
		}
		fmt.Println(j)
	}
}
