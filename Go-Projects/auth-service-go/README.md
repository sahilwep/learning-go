# Build Walkthrough:

## Docker-Compose for (Postgres & redis)

- Running compose:
```sh
$ docker compose up -d
$ docker ps
```

- Additionally, we can check the running status using `nc -vv localhost PORT` to get the name of the running services...

- To stop all the running services, we can use this command.
```sh
$ docker compose stop
```

## Config + Database Connections:

- Created `internal/config/config.go` & add .env file & wrote the logic of env fetching from the `.env` file..
- This helps us to separate our config from the code.
- For database, we will use `pgx` fast & Go-native.
- Written PostgreSQL connection logic into `internal/store/postgres.go` and successfully added into `cmd/main.go`.
- Make sure that our container images are running before connecting the database..


```sh
# ---------- Terminal 1 ----------
 go run cmd/server/main.go
2025/12/24 15:59:37 Connected to PostgreSQL
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /health                   --> main.main.func1 (3 handlers)
2025/12/24 15:59:37 Starting server on:  8080
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://github.com/gin-gonic/gin/blob/master/docs/doc.md#dont-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2025/12/24 - 16:00:28 | 404 |         458ns |       127.0.0.1 | GET      "/"
[GIN] 2025/12/24 - 16:00:28 | 404 |         708ns |       127.0.0.1 | GET      "/favicon.ico"



# ---------- Terminal 2 ----------
sahilwep~$ curl http://localhost:8080/health
{"db":"connected","status":"ok"}%                                                   
```

- Extra:
  - In production we use Connection pool, not every time we write a logic to {connect -> query -> close}, instead we uses the concept of connection pool..

