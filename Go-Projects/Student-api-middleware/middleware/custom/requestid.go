/*

// Custom Middleware:
	- Add a unique ID to each request for logging and tracing.
	- If we want to add a custom ID to every incoming request for better logging, we can inject this middleware

*/

package custom

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := uuid.New().String()
		c.Set("RequestID", reqID)                    // store in context for handler or logging
		c.Writer.Header().Set("X-Request-ID", reqID) // Optional: Send to client also
		c.Next()
	}
}
