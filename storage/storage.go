package storage

import "task1/book"

type Storage interface {
	GetById(id int) (*book.Book, bool)
	AddBook(b *book.Book)
	GetByTitle(title string) *[]book.Book
	UpdateIds(generator func() int)
}
