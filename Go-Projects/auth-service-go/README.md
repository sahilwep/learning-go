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

## User + Password:
- At this phase we will build:
  - User Table (proper schema)
  - Password hashing with bcrypt
  - signup API
  - Login API (no JWT yet)
  - Zero Plaintext password ever.

## Database Schema:
- We are not adding everything we will add only the things which matters now.
- Fields:
  - id  -> Primary key
  - email -> unique, login Identifier
  - role -> authorizations later
  - created_at
  - Updated_at

- We are not using migrations tool yet, first will do MySQL manual tool.

### Connect to Postgres Container:

```sh
docker exec -it auth_postgres psql -U auth_user -d auth_db

auth_db=#
```
- Now paste this SQL Carefully:

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'user',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);
```

- If `gen_random_uuid()` errors, run this first:

```sql
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
```

- Then we can check:

```sql
\d                    -- to list database tables
\d users              -- to list user table
SELECT * FROM users;  -- to query into table
\q                    -- to exit.
```

### Password Hashing:

- First install bcrypt into the local system using: `go get golang.org/x/crypto/bcrypt`
- Then we can write the logic inside the `internal/user/password.go`
- While Writing logic make sure that Never decrypt password.
- We only compare hash, bcrypt is slower by design making sure burteforce ineffective.

### Database Layer:

- Writing the logic for user repository (DB access layer)
- Create file inside `internal/users/repository.go`
- This will separate DB logic from HTTP logic, & makes testing possible.

### Writing Singup API (First endpoint):
- we will create a file insdie `internal/router/auth.go`.
- We will Create one function `RegisterAuthRoutes()` which will bind incoming data with JSON & then hash that using bcrypt & then store into the DB.
- Lastly we will wiring this logic into  `cmd/server/main.go` so that we can have our routes calling..

### Testing:
- After running a server & making sure that all the Docker images running.
- We can test `/signup` endpoint with curl:

```sh
curl -X POST http://localhost:8080/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","password":"secret123"}'

{"message":"user created"}
```

- Confirm this with DB:

```sh
docker exec -it auth_postgres psql -U auth_user -d auth_db
SELECT email, password_hash FROM users;
```

- As of now, our signup is working were user can signup thyself using {email, password} and {email, password} is stored into the DB after hashing.
- NOTE: We haven't handel the input validations as of now.


## Login + JWT (access & refresh)