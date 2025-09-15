/*

// Create In-memory Store:
	-> For now we will use in-memory storage (map) later on, we can update with any database like sql or postgresql..
	-> We have create a book struct say BS:{Id, Title, Author, Year}, basically model how our book will look like...
	-> But we are designing our database structure where we will place any x number of books & want to perform operations like add, delete update, So all those operations should reelect our storage{map}
	-> We will create a data structure like:
			map{
				bookID : model.book{}
				bookID : model.book{}
			}

			map[string]model.book type


	// Data Model
		-> We define "model.Book" struct (ID, Title, Author, Year) to represent a book entity.
		-> This keeps the business model independent from the storage mechanism.


	// Storage Data Structure
		-> We use "map[string]model.Book" where:
			Key   = unique BookID (string)
			Value = entire Book struct

		Example:
			map[string]model.Book {
				"id1": {ID:"id1", Title:"Go 101", Author:"Alice", Year:2024},
				"id2": {ID:"id2", Title:"System Design", Author:"Bob", Year:2023},
			}


	// Encapsulation
		-> Instead of writing global CRUD functions, we define a "BookStore" struct:

			type BookStore struct {
			    books map[string]model.Book
			}

		-> And attach methods (Add, Get, Update, Delete).
		-> This encapsulates the storage and makes the design loosely coupled.



	// Why this design?
		-> Loose coupling: The service layer interacts with `BookStore` methods, not the map directly. Later, we can swap the internal map with a SQL database without changing service/business logic.
		-> Reusability: "model.Book" is reusable across services without duplicating CRUD logic.
		-> Scalability: This design scales conceptually—today it's an in-memory map, tomorrow it could be Postgres, Redis, or a distributed cache, with the same API.
		-> Maintainability: By returning errors like "ErrNotFound", higher layers (API) can make informed decisions.


*/

package store

import (
	"errors"

	"github.com/sahilwep/learning-go/Go-Projects/Book-Management-api/model"
)

var ErrNotFound = errors.New("book not found") // error message when any book is not found inside the storage

type BookStore struct {
	books map[string]model.Book // create a map: {key: string(id) & val:{struct of Book{}}}
}

// Function to create bookStore, it's constructor for above struct:
func NewBookStore() *BookStore {
	book := BookStore{
		books: make(map[string]model.Book),
	}

	return &book

	// Or directly return like this:
	// return &BookStore{
	// 	books: make(map[string]model.Book),
	// }
}

// Create a method for BookStore -> ADD(): This will create a new space into map & add book to it..
func (b *BookStore) Add(book model.Book) {
	b.books[book.Id] = book
}

// Function to fetch book from map by their ID & return book property if found else error...
func (b *BookStore) Get(id string) (model.Book, error) {
	book, err := b.books[id]

	// If book not found:
	if !err {
		return model.Book{}, ErrNotFound // if book is not found return (zero value struct & error)
	}

	return book, nil // else return book & error as nil.
}

// When this function will return all the available books in map of type: slice of mode.Book
func (b *BookStore) List() []model.Book {
	all := []model.Book{}

	for _, val := range b.books {
		all = append(all, val)
	}

	return all
}

// Function to update Book values (If only book ID was there...):
func (b *BookStore) Update(book model.Book) error {
	_, ok := b.books[book.Id]

	if !ok {
		return ErrNotFound
	}

	b.books[book.Id] = book // update book values into their specific book ID
	return nil
}

// Delete book from bookStore:
func (b *BookStore) Delete(id string) error {
	_, ok := b.books[id] // fetch the book from map

	if !ok {
		return ErrNotFound // return book not found
	}

	// If it was there => remove book from there...
	delete(b.books, id)
	return nil // as operations is done, return nil
}
