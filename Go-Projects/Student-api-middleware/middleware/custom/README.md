# Custom Middleware:

- In backend development, a middleware is a piece of code that runs before or after the main request handler.
- Custom middleware is middleware that you write yourself, tailored to your specific applications needs.
- Unlike built in middleware (like logging, CORS, or recovery), custom middleware is your own logic to handle requests or responses.
  
## Key Points:
- Runs on every request (or specific routes)
- Can inspect, modify, block, or log requests.
- Can add data to the Context for handler to use
- Can stop the request early if needed (c.Abort() in Gin)


## Real-world example / use cases
- API key Validation
- Only allow requests with a valid API key.
- Example: /student API is private, only internal app can access.
- Middleware checks header X-API-Key -> if invalid, return 401 Unauthorized.
- Request Timing / Performance metrics
- Track how long each request takes and store matrices in prometheus or logs.
- Custom Header injection
- Add a standard header to all responses, e.g., X-App-Version: 1.0.
- IP whitelisting / Blacklisting
- Only allow a certain IPs to access sensitive endpoints.
- Feature Flags / Maintenance Mode.
- Return 503 Service Unavailable if API is under maintenance.


## How it works:
- A middleware functions wraps around the request lifecycle.
- In Gin, it looks like this:

```go
func MyMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Code before handler
        c.Next() // call the main handler
        // Code after handler
    }
}
```
- You can stop the request early by using c.Abort().
- You can also add values to the context using c.Set("key", value).


## APIKey:

- Make a `GET` request on `http://localhost:8080/student` with header

```plain
Accept: */*
User-Agent: Sahil Sharma
Origin: http://localhost:3000
Content-Type: application/json
X-API-Key: sahilwep
```
- We will get authorized access of every resource

- If we temper with `X-API-Key`, we will get unauthorized access.


## Request ID:

- For every request we will get an additional details about the "X-Request-ID"
- Response Headers
```plain
access-control-allow-credentials: true
access-control-allow-headers: Content-Type, Authorization
access-control-allow-methods: GET, POST, PUT, DELETE, OPTIONS
access-control-allow-origin: http://localhost:3000
content-type: application/json; charset=utf-8
x-request-id: 1f3f8a42-8114-476b-bab0-e5ef0e097bce
date: Fri, 03 Oct 2025 05:20:09 GMT
content-length: 2
connection: close
```

## Custom Maintenance Mode:

- If our server is on maintenance mode, we can used this custom middleware.
- Normal mode:
```plain
[]
```

- Maintenance Mode:
```plain
{
  "error": "service is under maintenance, try again later"
}
```
