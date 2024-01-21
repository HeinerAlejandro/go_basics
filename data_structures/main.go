package main

import "fmt"

func main() {
	//First we start from matrixes
	//var matriz [5]int for declaration
	//var matriz = [5]int{...n} for inizialization
	//var matriz = [...]int{...n} We we don't know how many elmements will have
	var matriz = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Valor inicial")
	fmt.Println(matriz)

	matriz[0] = 4
	matriz[1] = 7
	matriz[2] = 41

	fmt.Println("Nuevo valor de la matriz")
	fmt.Println(matriz)

	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}

	// Other way is using range, it's like enumarate of python (return index and value)

	for index, value := range matriz {
		fmt.Printf("Indice: %d - Value: %d\n", index, value)
	}

	// declare matrix 2d

	var matriz_2d = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(matriz_2d)
}
