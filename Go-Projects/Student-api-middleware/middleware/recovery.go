/*

// Recovery Middleware:

	// What is recovery middleware?
		- In Go, if your code panics(like a runtime crash: nil pointer dereference, divide by zero, out of range index, etc.), your entire program will crash and stop serving requests.
		- All current request fails.
		- Your server goes down until restarted
		- Bad for productions.
		- Example a function will panic because of nil pointer dereference.

				func (s *StudentHandler) CrashEndpoint(c *gin.Context) {
					var stu *model.Student
					// OOPS: stu is nil, so this will panic (nil pointer dereference)
					c.JSON(http.StatusOK, stu.Name)
				}


	- The Recovery middleware acts like an airbag in a car - you hope you never need it, but if a crash happens, it saves you.
	- It catches the panic, logs the error and responds gracefully to the client instead of killing the whole server.

	// How it works:
		- `defer` ensure the function runs at the end of the request handling.
		- recover() capture the panic if one occurred.
		- instead of server crash:
			- logs the panic in backend logs.
			- Returns a 500 Internal server error to the client.

	//  UseCase:
		- For recovery commit, i have demonstrated how recovery function takes place...
		- Start server & hit: http://localhost:8080/student/crash

*/

package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the panic for debugging
				log.Printf("Panic Caught: %v", err)

				// Response gracefully to client:
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error, please try again later",
				})
			}
		}()

		c.Next() // continue request
	}
}
