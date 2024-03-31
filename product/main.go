package main

import "fmt"

type Product[T ~string | ~uint] struct {
	Id    T
	Desc  string
	Price float32
}

func main() {
	product_1 := Product[uint]{
		Id:    1,
		Desc:  "Lavadora",
		Price: 1000,
	}

	product_2 := Product[string]{
		Id:    "MI ID: 2",
		Desc:  "TV",
		Price: 400,
	}

	fmt.Printf("Este es mi producto 1: %v \n Este es mi producto 2: %v\n`", product_1, product_2)

}
