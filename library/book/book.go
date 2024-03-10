package book

import "fmt"

type Book struct {
	title  string
	author string
	pages  int
}

func (book *Book) PrintInfo() {
	fmt.Printf("'%s'-Author:%s-Pages:%d\n", book.title, book.author, book.pages)
}

func (book *Book) SetTitle(title string) {
	book.title = title
}

func (book *Book) SetAuthor(author string) {
	book.author = author
}

func (book *Book) GetAuthor() string {
	return book.author
}

func (book *Book) SetPages(pages int) {
	book.pages = pages
}

func (book *Book) GetPages() string {
	return book.title
}

func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}
