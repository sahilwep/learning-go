# Middleware Implementation on Student-api

- We have used project student-api, & Implementing middleware:
  - middleware/recovery.go - catch panics
  - middleware/logger.go - log request + response time + correlations IDs
  - middleware/cors.go - configuration cross-organ resource sharing
  - middleware/timeout.go - set timeouts on requests so they don't hang forever
  - middleware/ratelimit.go - limit client request to prevent abuse

```plain
student-api/
├─ api/
│   └─ student_handler.go
├─ model/
│   └─ student.go
├─ service/
│   └─ student_service.go
├─ store/
│   └─ student_store.go
├─ middleware/
│   ├─ recovery.go
│   ├─ logger.go
│   ├─ cors.go
│   ├─ timeout.go
│   └─ ratelimit.go
├─ main.go
└─ go.mod
```

## initial Configurations:

- Make sure to install all the required modules..

```sh
cd path/to/student-api
go get github.com/google/uuid
go get golang.org/x/time/rate
go mod tidy
go run main.go
```

## middleware/recovery.go

- 