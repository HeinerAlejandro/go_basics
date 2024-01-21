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

	// Slice (Rebanada)
	// It has any number between brackets. It's a dynamic structure

	var slice []int // or var slice = []int{1,2,3,4,...,n}
	fmt.Println(slice)

	var sliceTwo = []string{
		"Lunes",
		"Martes",
		"Miercoles",
		"Jueves",
		"Viernes",
		"Sabado",
		"Domingo",
	}
	fmt.Println("My First Slice:")
	fmt.Println(sliceTwo)

	// Slice from another

	sliceFromSliceTwo := sliceTwo[0 : cap(sliceTwo)-2]
	fmt.Println("Creating a Slice from Another")
	fmt.Println(sliceFromSliceTwo)

	//Adding Elements

	sliceFromSliceTwo = append(sliceFromSliceTwo, "Word 1", "Word 2", "Word 3")
	println(sliceFromSliceTwo)
	println(len(sliceFromSliceTwo))
	println(cap(sliceFromSliceTwo))

	fmt.Println("Eliminando un elemento")

	sliceFromSliceTwo = append(sliceFromSliceTwo[:2], sliceFromSliceTwo[3:]...)

	fmt.Println(sliceFromSliceTwo)

	// crear rebanada con make (Type, long, capacity)
	sliceX := make([]int, 5)
	sliceX[0] = 10
	sliceX[1] = 11
	sliceX[2] = 12
	sliceX[3] = 18
	sliceX[4] = 20

	sliceY := []int{1, 2, 3, 4, 5}

	fmt.Println(sliceX)
	fmt.Println(sliceY)

	composedSlice := append(sliceY, sliceX...)
	// Con copy se copia el contenido de un slice a otro(se sobreescriben)
	copy(sliceX, sliceY)
	fmt.Println(composedSlice)
	fmt.Println("Using Copy")
	fmt.Println(sliceX)
}
