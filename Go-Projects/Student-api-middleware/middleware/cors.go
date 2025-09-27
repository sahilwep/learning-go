/*

// CORS:
	// What is CORS:
		- CORS stand for Cross-Origin Resource Sharing.
		- Origin = protocol + domain + port
		- Example:
			http://localhost:3000 → React frontend
			http://localhost:8080 → Go Gin backend

		- Even though both are on localhost, the port differ, so they are different origins.
		- By default, browser blocks request made from one origin to another (for security reason).

	// Why do we need CORS?
		- Without CORS:
			Your frontend (say React at http://localhost:3000) cannot call your API (http://localhost:8080/api/students) because the browser says: "Blocked by CORS policy".

		- With CORS:
			The backend explicitly tells the browser:
				"It's okay, I trust this frontend, You can make requests."

	// Real Life Example:
		- Frontend + Backend Dev Setup
			React frontend (localhost:3000)
			Go Gin backend (localhost:8080)
			- Without CORS -> every fetch request fails.
			- With CORS -> frontend can consume backend APIs.
		- Public APIs
			- Suppose your API server data multiple clients(e.g., weather API)
			- you may want to allow only specific trusted domains (e.g., myweatherapp.com) and block others.
		- Security
			- You can restrict methods: allow only GET and POST, block DELETE and PUT.
			- Prevents malicious sites from abusing your API.


	// Purpose of a CORS middleware:
		- A CORS middleware sites between the request and your routes, and:
			- Checks if the incoming request's Origin is allowed.
			- Sets proper CORS headers (Access-Control-Allow-Origin, etc.).
			- Responds to preflight requests (OPTIONS) before they even reach your routers.

	// Preflight Request (OPTIONS)
		- When you send a request like POST, PUT, or DELETE from JS, the browser first sends a preflight request:

				OPTIONS /api/students HTTP/1.1
				Origin: http://localhost:3000
				Access-Control-Request-Method: POST
				Access-Control-Request-Headers: Content-Type

		- The backend must reply with:

				Access-Control-Allow-Origin: http://localhost:3000
				Access-Control-Allow-Methods: GET, POST, PUT, DELETE
				Access-Control-Allow-Headers: Content-Type


*/

package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {

	// Reads list of allowed origin from .env file
	allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Check if request origin is in allowed list:
		allowed := false
		for _, o := range allowedOrigins {
			if origin == o {
				allowed = true
				break
			}
		}

		// Debugging: Debug the origin & allowedOrigins
		// log.Printf("Origin from request: '%s'", origin)
		// log.Printf("Allowed origins: %#v", allowedOrigins)

		if !allowed && origin != "" {
			// Origin provided is not in allowed list -> block
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "CORS policy: Origin not allowed",
			})
			return
		}

		if allowed {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Preflight:
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
