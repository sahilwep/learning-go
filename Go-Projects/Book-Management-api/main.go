/*

// Main Entry for our application Executions:



	// Explanation:
		-> Create a Gin engine with default middlewares:
			r := gin.Default() creates a gin engine pre-wired with two middleware:
				Logger - logs request to stdout (method, path, status, latency).
				Recovery - recover from panics and returns HTTP 500 instead of crashing the process.

		-> We then create the pieces of the application:
			bookStore := store.NewBookStore() - instantiate our store
				Initializing storage. For now this returns an in-memory store.
				In future we can replace this with a Postgres-backend implementations
				By implementing the same interface used by the service.


			bookService := service.NewBookService(bookStore) - create a service and inject the store dependency.
				Create the business-logic layer and inject the store dependency.
				Dependency injection makes testing easier and future swaps(e.g. Postgres)..
				easy to keeps components decoupled.


			bookHandler := api.NewBookHandler(bookService) - create the HTTP handler with injected service...
				Create the HTTP Handlers, injection the service layer.
				HTTP handlers should be thin and delegate to service for logic.

		-> Register routes: map HTTP method + path -> handler function
			NOTE: ":id" is a path parameter accessible in handlers via c.Param("id")
				r.POST("/book", bookHandler.CreateBook)
				...
				r.GET("/book/:id", bookHandler.GetBook)
				...

		-> Run the HTTP server on local port 8080 (blocks until process is killed)


*/

package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sahilwep/learning-go/Go-Projects/Book-Management-api/api"
	"github.com/sahilwep/learning-go/Go-Projects/Book-Management-api/service"
	"github.com/sahilwep/learning-go/Go-Projects/Book-Management-api/store"
)

func main() {

	r := gin.Default() // Initialize router

	// Create pieces of application:
	bookStore := store.NewBookStore()
	bookService := service.NewBookService(bookStore)
	bookHandler := api.NewBookHandler(bookService)

	// Register routes:
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books", bookHandler.ListBooks)
	r.GET("/books/:id", bookHandler.GetBook)
	r.PUT("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	// serve:
	if err := r.Run("0.0.0.0:8080"); err != nil { // Basic server run call
		log.Fatal(err)
	}

}
