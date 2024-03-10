package main

import (
	"library/animal"
	"library/book"
)

func main() {
	var myBook = book.NewBook(
		"Pensar bien, Sentirse bien",
		"Walter Riso",
		250,
	)

	myBook.SetTitle("Una vida pacifica")
	myBook.SetAuthor("Heiner Enis")
	myBook.SetPages(100)

	myBook.PrintInfo()

	myTextBook := book.NewTextBook(
		"El Arte de Amar",
		"Desconocido",
		250,
		"Planeta",
		"2",
	)

	myTextBook.PrintInfo()

	book.Print(myBook)
	book.Print(myTextBook)

	dog := animal.Dog{
		Nombre: "Floppy",
	}

	cat := animal.Cat{
		Nombre: "Bombonkan",
	}

	animal.DoSound(&dog)
	animal.DoSound(&cat)
}
