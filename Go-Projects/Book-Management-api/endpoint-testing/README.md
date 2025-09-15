# Endpoint Testing:

- CRUD Operations:
```plain
POST    /books          -> Create: books by passing JSON{"title":"xyz","author":"SSS","year":123 } object
GET     /books          -> Read: List all available books 
GET     /books:id       -> Read: Specific book by passing id param
PUT     /books:id       -> Update: Update books Details by id param
DELETE  /books:id       -> Delete: delete books by id param
```

## Create:

- `POST` req on `http://localhost:8080/books/`
```JSON
{
  "title": "XYZ",
  "author": "Sahil",
  "year": 1039
}
```

## Read:
- `GET` req on `http://localhost:8080/books/` to list all books.
- `GET` req on `http://localhost:8080/books/{id}` to list specific book.


## Update:
- `PUT` req on `http://localhost:8080/books/{id}` to update values by passing json object on specific id.
```JSON
{
    "title": "xyz",
    "author": "sahil",
    "year": 234
}
```

## Delete:
- `DELETE` req on `http://localhost:8080/books/{id}` to delete delete specific book
