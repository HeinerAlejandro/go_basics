package book

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func (book *Book) PrintInfo() {
	fmt.Printf("'%s'-Author:%s-Pages:%d\n", book.Title, book.Author, book.Pages)
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Pages:  pages,
	}
}
