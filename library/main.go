package main

import "library/book"

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
}
