package storage

import "task1/book"

type SliceStorage struct {
	Books []book.Book
}

func (storage *SliceStorage) GetById(id int) (*book.Book, bool) {
	for _, curBook := range storage.Books {
		if curBook.Id == id {
			return &curBook, true
		}
	}
	return nil, false
}

func (storage *SliceStorage) GetByTitle(title string) *[]book.Book {
	var foundBooks []book.Book
	for _, curBook := range storage.Books {
		if curBook.Title == title {
			foundBooks = append(foundBooks, curBook)
		}
	}
	return &foundBooks
}

func (storage *SliceStorage) AddBook(book *book.Book) {
	storage.Books = append(storage.Books, *book)
}

func (storage *SliceStorage) UpdateIds(generator func() int) {
	for i := range storage.Books {
		storage.Books[i].Id = generator()
	}
}
