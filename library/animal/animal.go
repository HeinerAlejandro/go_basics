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
	fmt.Printf("%s hace guau guauu", dog.Nombre)
}

func (cat *Cat) Sound() {
	fmt.Printf("%s hace miau miau\n", cat.Nombre)
}

func DoSound(animal Animal) {
	animal.Sound()
}
