package book

import "fmt"

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (book *Book) PrintInfo() {
	fmt.Printf("'%s'-Author:%s-Pages:%d\n", book.Title, book.Author, book.Pages)
}
