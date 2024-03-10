package main

import "library/book"

func main() {
	var myBook = book.Book{
		Title:  "El Arte de Amar",
		Author: "Walter Riso",
		Pages:  250,
	}

	myBook.PrintInfo()
}
