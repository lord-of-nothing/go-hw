package library

import (
	"task1/book"
	"task1/storage"
)

type Library struct {
	Storage storage.Storage
	IdGen   func() int
}

func (library *Library) AddBook(title, author string) {
	newBook := book.Book{Id: library.IdGen(), Title: title, Author: author}
	library.Storage.AddBook(&newBook)
}

func (library *Library) GetByTitle(title string) []book.Book {
	return *library.Storage.GetByTitle(title)
}

func (library *Library) UpdateIdGen(newIdGen func() int) {
	library.IdGen = newIdGen
	library.Storage.UpdateIds(library.IdGen)
}
