/*

// Service Layer:
	-> This is the place where we place all the business logic of our application..
	-> This is our brain of our application:
		-> Right now, we have a BookStore (data access/storage) in /store/book_store.go
		-> This layer only know how to add, fetch... books in memory (or later in DB).
		-> But our app isn't just talking about to the DB. It has rule like:
			"A book's title must not be empty"
			"Each book must have a unique ID"
			"Year must be grater than 0"
			"Don't allow deletion of book that are 'archived'"
		-> These rules are not DB concerns, they are business rules -> and this is where the service layer comes in..


	-> If we directly put all logic inside store or controllers/handler, things get messy very fast.
	-> So we separate layers:
		Controllers/Handlers layer (API endpoints)
			Talks to the outside world (HTTP request/response)
			Example: /books GET, /books POST
			It shouldn't know business rules, just pass input to service.
		Service layer(business logic)
			contains your real application rules
			Example: Validate book details, check duplicates, enforce rules.
			calls the store layer to actually persist/fetch data.
		Store/Repository layer(data access):
			Responsible for storage operations like (DB, cache, in-memory).
			It doesn't care about rules.

	// Code Explanation:
		NewBookService() -> Constructor, enforces dependencies
		CreateBook() -> adds new book with unique ID.
		GetBook() -> fetches book by ID.
		ListBooks() -> returns all books
		UpdateBook() -> changes existing book
		DeleteBook() -> removes book.


*/

package service

import (
	"github.com/google/uuid"
	"github.com/sahilwep/learning-go/Go-Projects/Book-Management-api/model"
	"github.com/sahilwep/learning-go/Go-Projects/Book-Management-api/store"
)

// We're using dependency injection, The service doesn't "create" a store itself, instead we inject one from outside(in main.go)
type BookService struct {
	store *store.BookStore // we use pointer so the service works with the same instance everywhere(shared state).
}

// This is Constructor for BookService
func NewBookService(store *store.BookStore) *BookService {
	return &BookService{store: store}
}

// function to create new Book Record:
func (s *BookService) CreateBook(title, author string, year int) model.Book {
	book := model.Book{
		Id:     uuid.New().String(), // uuid.New() is used to generate a new unique id for book
		Title:  title,
		Author: author,
		Year:   year,
	}

	/*
		You might be wondering, we have Add() method in store why we are rewriting everything?
		NOTE: If any Creation or Modification needed, we can do it here, not in API or DB
		eg: "Year must be > 0"
	*/

	s.store.Add(book) // Add book into BookStore
	return book       // return book
}

// Function to retrieve book if available:
func (s *BookService) GetBook(id string) (model.Book, error) {
	return s.store.Get(id) // Get() method fetches from store package of BookService struct & return directly..
}

// Function to retrieve all the books from storage:
func (s *BookService) ListBooks() []model.Book {
	return s.store.List() // directly return the list, as if there's nothing, it will return empty list.
}

// Function used to update the book values:
func (s *BookService) UpdateBook(book model.Book) error {
	return s.store.Update(book)
}

// Function to delete book from database:
func (s *BookService) Delete(id string) error {
	return s.store.Delete(id)
}
