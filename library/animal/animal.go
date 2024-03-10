package animal

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct {
	Nombre string
}

type Cat struct {
	Nombre string
}

func (dog *Dog) Sound() {
	fmt.Println("Hago guau guauu")
}

func (cat *Cat) Sound() {
	fmt.Println("Hago miau miau")
}

func DoSound(animal Animal) {
	animal.Sound()
}
