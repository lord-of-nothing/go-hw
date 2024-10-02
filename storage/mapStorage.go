package storage

import "task1/book"

type MapStorage struct {
	Books map[int]book.Book
}

func (storage *MapStorage) GetById(id int) (*book.Book, bool) {
	for _, curBook := range storage.Books {
		if curBook.Id == id {
			return &curBook, true
		}
	}
	return nil, false
}

func (storage *MapStorage) GetByTitle(title string) *[]book.Book {
	var foundBooks []book.Book
	for _, curBook := range storage.Books {
		if curBook.Title == title {
			foundBooks = append(foundBooks, curBook)
		}
	}
	return &foundBooks
}

func (storage *MapStorage) AddBook(book *book.Book) {
	if storage.Books == nil {
		storage.Books = make(map[int]book.Book)
	}
	storage.Books[book.Id] = *book
}

func (storage *MapStorage) GetAllBooks() *map[int]book.Book {
	return &storage.Books
}

func (storage *MapStorage) UpdateIds(generator func() int) {
	for _, curBook := range storage.Books {
		curBook.Id = generator()
	}
}
