/*
// Custom Maintenance mode middleware:
	- Temporarily block all requests during maintenance or deployment.


*/

package custom

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MaintenanceMiddleware blocks requests if the server is in maintenance mode
func MaintenanceMiddleware(active bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if active {
			c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
				"error": "service is under maintenance, try again later",
			})
			return
		}
		c.Next()
	}
}
