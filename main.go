package main

import (
	"fmt"
	"task1/book"
	library2 "task1/library"
	"task1/storage"
)

func main() {
	books := []book.Book{
		{Title: "Понедельник начинается в субботу", Author: "А.Н.Стругацкий, Б.Н.Стругацкий"},
		{Title: "Смерть этажом ниже", Author: "Кир Булычёв"},
		{Title: "Детство", Author: "М.Горький"},
		{Title: "Детство", Author: "Л.Н.Толстой"},
		{Title: "Хищные вещи века", Author: "А.Н.Стругацкий, Б.Н.Стругацкий"},
	}

	slcStorage := storage.SliceStorage{}
	library := library2.Library{Storage: &slcStorage, IdGen: library2.GenerateAllId()}
	for _, curBook := range books {
		library.AddBook(curBook.Title, curBook.Author)
	}

	fmt.Println("Slice as storage")
	fmt.Println("Поиск по названию 'Смерть этажом ниже': ", library.GetByTitle("Смерть этажом ниже"))
	fmt.Println("Поиск по названию 'Детство' (несколько книг с одним названием): ", library.GetByTitle("Детство"))
	fmt.Println("Поиск по несуществующему названию:", library.GetByTitle("Стажеры"))

	library.UpdateIdGen(library2.GenerateEvenId())
	fmt.Println("\nChanged id generator")
	fmt.Println("Поиск по названию 'Смерть этажом ниже': ", library.GetByTitle("Смерть этажом ниже"))
	fmt.Println("Поиск по названию 'Понедельник начинается в субботу'", library.GetByTitle("Детство"))

	mpStorage := storage.MapStorage{}
	library.Storage = &mpStorage
	for _, curBook := range books {
		library.AddBook(curBook.Title, curBook.Author)
	}
	fmt.Println("\nMap as storage")
	fmt.Println("Поиск по названию 'Смерть этажом ниже': ", library.GetByTitle("Смерть этажом ниже"))
	fmt.Println("Поиск по названию 'Детство' (несколько книг с одним названием): ", library.GetByTitle("Детство"))
	fmt.Println("Поиск по несуществующему названию:", library.GetByTitle("Стажеры"))
}
