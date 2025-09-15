# Simple Book Management API

## Features:

- Add a book (POST /books)
- List all books (GET /books)
- Get book by ID (GET /books/:id)
- Update a book (PUT /books/:id)
- Delete a book (DELETE /books/:id)

## Stack:

- Go (latest stable)
- Gin (lightweight HTTP framework)
- In-memory store (map) for now — no database, we’ll add Postgres later
- Unit tests

## Directory Structure

```plain
book-api/
├─ main.go                # entry point
├─ go.mod                 # module file
├─ /api                   # HTTP handlers
│   └─ book_handler.go
├─ /model                 # data models
│   └─ book.go
├─ /service               # business logic
│   └─ book_service.go
├─ /store                 # storage (in-memory for now)
│   └─ book_store.go
└─ /tests                 # unit tests
    └─ book_test.go

```

## How to use:

```sh
go run /book-api/main.go
```