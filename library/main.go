package main

import "library/book"

func main() {
	var myBook = book.NewBook(
		"El Arte de Amar",
		"Walter Riso",
		250,
	)

	myBook.PrintInfo()
}
