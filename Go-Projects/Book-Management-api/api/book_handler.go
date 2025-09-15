/*

// Handler:
	-> This file is the HTTP layer (also called handler or controller layer) of you application.
	-> It's job is to:
		-> Translate incoming HTTP request into typed Go values (parse JSON, path params, query params).
		-> Validate / enforce basic request-level invariants.
		-> Call business logic (the service package) to perform operations.
		-> Translate the results / errors from the service into HTTP responses (status code + JSON).

	// Why Separate this file / package ?
		-> Separations  of concerns:
			-> The handler should only care about HTTP concerns(parsing, status code, handlers),
			-> while the service handles business rules & store handles persistence. This makes the code easier to test, maintain, and replace parts later (e.g, swap in Postgres for the in-memory store).


// Gin Context quick Notes:
	-> *gin.Context encapsulates request & response objects:
		-> c.request.Context() gives you the context.Context for cancellation, deadlines, and request-scoped values (useful to pass down).
		-> ShouldBindJSON reads and binds the request body to a struct; it also respects Content-Type.
		-> c.Param("id") reads a path params, like ("/read", "/write",...)
		-> c.JSON(status, obj) sets Content-Type: application/json and writes the body.



*/

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sahilwep/learning-go/Go-Projects/Book-Management-api/model"
	"github.com/sahilwep/learning-go/Go-Projects/Book-Management-api/service"
)

type BookHandler struct {
	service *service.BookService
}

// Constructor for above struct helps to create the handler with it's dependency
func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// Function to parse the JSON body into a small DTO (Data Transfer Object), Using DTO prevents clients from setting field they shouldn't
func (h *BookHandler) CreateBook(c *gin.Context) {
	// create a temporary struct:
	var req struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Year   int    `json:"year"`
	}

	// This is like taking input from user & validating from any malformed input.
	err := c.ShouldBindJSON(&req) // Validate JSON structure, This will return 404 for any malformed input.

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // if we get any error
		return
	}

	// Calls h.service.CreateBook(...) which encapsulates ID generations and persistence
	book := h.service.CreateBook(req.Title, req.Author, req.Year)

	c.JSON(http.StatusCreated, book) // return 201 Created with the created book in the response body.
}

// Function to fetch specific book from the storage using their id:
func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")

	book, err := h.service.GetBook(id) // fetch book from the database

	if err != nil { // if book not is found:
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book) // return successful response for request, return book object
}

// Function to list al list all the available books from database:
func (h *BookHandler) ListBooks(c *gin.Context) {
	book := h.service.ListBooks() // fetch book from database
	c.JSON(http.StatusOK, book)   // return successful response for request, return book object
}

// Function to Update Books details:
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book model.Book // create a empty instance of struct

	// bind incoming JSON data value with above empty book struct & handel error:
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // if we got any error during JSON binding with above struct instance, return error
		return
	}

	book.Id = id // else update the book id to the incoming id

	// Now update book details into database & if not found return error
	if err := h.service.UpdateBook(book); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book) // return successful response for request, return book object
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nil) // else return nil as delete request successful
}
