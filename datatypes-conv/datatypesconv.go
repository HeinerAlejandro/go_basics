package main

import "fmt"

func main() {
	var myInt32 int32 = 56
	var myInt16 int16 = 12

	var myFloat32 float32 = 49.9

	fmt.Println("Suma de diferentes tipos de enteros", int32(myInt16)+myInt32)
	fmt.Println("Suma de diferentes de entero y floatante", float32(myInt16)+myFloat32)

}
