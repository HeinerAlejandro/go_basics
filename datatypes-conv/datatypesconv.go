package main

import (
	"fmt"
	"strconv"
)

func main() {
	var myInt32 int32 = 56
	var myInt16 int16 = 12

	var myFloat32 float32 = 49.9

	fmt.Println("Suma de diferentes tipos de enteros", int32(myInt16)+myInt32)
	fmt.Println("Suma de diferentes de entero y floatante", float32(myInt16)+myFloat32)

	myStrInt := "804"
	myStrToIntConversion, _ := strconv.Atoi(myStrInt)

	fmt.Println("Conversion de String a Entero", myStrToIntConversion)
	fmt.Println("Suma con valor de str a entero convertido", myStrToIntConversion+int(myInt32))

	// int != int32 != int64

	myIntToStr := strconv.Itoa(int(myInt32))

	fmt.Println("Convertir de Int a Str", myIntToStr)
	fmt.Println("Concatenando valores", myIntToStr+myIntToStr)
}
