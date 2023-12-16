package repository

import (


	"github.com/google/uuid"
	desc "ubulllka.com/rschir10-book-server/pkg/api"
)

var array []*desc.Book
func init() {
	array = append(array, &desc.Book{Id: "a68b823c-7ca6-44bc-b721-fb4d5312cafc", Name: "1984", Author: "George Orwell"})
	array = append(array, &desc.Book{Id: "34415c7c-f82d-4e44-88ca-ae2a1aaa92b7", Name: "The Green Mile", Author: "Stephen King"})
}



func GetAll() (*desc.BookList, error) {
	return &desc.BookList{
		Books: array,
	}, nil
}

func Get(in *desc.BookId) (*desc.Book, error) {
	var product *desc.Book
	for _, e := range array {
		if e.Id == in.Id {
			product = e
			break
		}
	}
	return product, nil
}

func Insert(newBook *desc.Book) (*desc.Book, error) {
	uuid := uuid.NewString()
	newBook.Id = uuid
	array = append(array, newBook)
	return newBook, nil
}

func Update(updateBook *desc.Book) (*desc.Book, error) {
	var book *desc.Book
	for _, e := range array {
		if e.Id == updateBook.Id {
			e.Name = updateBook.Name
			e.Author = updateBook.Author
			book = e
			break
		}
	}
	return book, nil
}

func Remove(in *desc.BookId) (*desc.Book, error) {
	var book *desc.Book
	var index int
	for i, e := range array {
		if e.Id == in.Id {
			book = e
			index = i
			break
		}
	}
	array = append(array[:index], array[index+1:]...)
	return book, nil
}