package main

import "fmt"

// To create a na structure we use type word

type Persona struct {
	nombre       string
	edad         int
	correo       string
	nacionalidad string
}

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

	// Maps

	colors := map[string]string{
		"red":   "ff0000",
		"blue":  "0000FF",
		"green": "#008000",
	}

	fmt.Printf("Available Colors: %v\n", colors)

	// Get Value

	valor, ok := colors["red"]

	// if ok is false, valor will be an empty string

	if ok {
		fmt.Printf("Obtained Color: %v - %v\n", ok, valor)
	} else {
		fmt.Printf("El Color: %v No existe (%v)\n", ok, valor)
	}

	// Deleting Key

	delete(colors, "blue")

	fmt.Printf("After Delete: %v\n\n\n", colors)

	fmt.Printf("Iterating by key\n\n")
	// Iter through map

	for index, valor := range colors {
		fmt.Printf("Key: %s - Valor: %s\n", index, valor)
	}

	// using structure

	fmt.Printf("\nUsing structures\n")

	persona := Persona{"", 0, "", "Venezolana"}
	persona2 := Persona{"Maria Del Rey", 23, "mariadelrey@gmail.com", "Venezolana"}

	persona.nombre = "Heiner Enis"
	persona.edad = 25
	persona.correo = "alguncorreo@gmail.com"

	fmt.Println(persona)
	fmt.Println(persona2)

}
