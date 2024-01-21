package main

import (
	"fmt"
	"time"
)

func main() {

	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Estas en la maniana")
	} else if t.Hour() < 18 {
		println("Estas en la tarde")
	} else {
		println("Estas en la noche")
	}
}
