/*

// Custom Middleware:
	- Allow requests with a valid API key. Prevent unauthorized access.
	- This middleware checks the valid API key is in request or not, if it's not it will return back immediately.
	- If we have Valid API key, it will process for next middleware / handler.


	// TESTING:
		- we will pass validKey as "sahilwep"
		- If request header has {X-API-Key: sahilwep}, then only we will process to authorized.


*/

package custom

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// APIKeyMiddleware ensure requests have a valid API key in the header
func APIKeyMiddleware(validKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key != validKey {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid API key",
			})
			return
		}

		c.Next() // Key is valid continue
	}
}
