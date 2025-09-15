# Student-API

- Basic Student API to Perform ALL CRUD applications..


## Features:
- Add a Student                 (POST   /student)
- List all Student              (GET    /student)
- Get Student by id             (GET    /student/:id)
- Update Student details        (PUT    /student/:id)
- Delete a Students             (DELETE /student/:id)

## Stack:
- Go (latest stable)
- Gin (lightweight HTTP framework)
- In-memory store (map) for now — no database, we’ll add any database later-on


## Directory Structure

```plain
Student-api/
├─ main.go                # entry point
├─ go.mod                 # module file
├─ /api                   # HTTP handlers
│   └─ student_handler.go
├─ /model                 # data models
│   └─ student.go
├─ /service               # business logic
│   └─ student_service.go
└─ /store                 # storage (in-memory for now)
    └─ student_store.go
```

## How to use:

```sh
go run main.go
```