/*

// Timeout Middleware:

	// What is timeout middleware:
		- Think of your server as help desk with limited staff.
		- A visiter arrives with complicated request, If one visitor occupies a staff member forever, no one else can be served.
		- A Timeout middleware is like the desk's timer: "You have 10 min to finish, If you don't we'll politely stop you and help next visitor."
		- In HTTP terms: a timeout ensures a request doesn't hang forever and tie up server resources.

	// Why use a timeout middleware (purpose)
		- Protect resource - prevent slow or stuck request from exhausting worker threads / goroutines.
		- Maintain responsiveness - ensure SLA / Latency targets: request that would take too long should fails test.
		- Avoid cascading failures - slow external service -> many request wait -> System degrades. Timeout limits the damage.
		- Signal to client - Client receive 504 Gateway Timeout (or similar) instead of waiting indefinitely.

	// Code Explanation:
		- "SimpleTimeoutMiddleware" accept timeout parameter which is the durations like: 2 sec, 4 sec..
		- Function return Gin middleware(gin.handlerFunc)
		- "return func(c *gin.Context)" we define the middleware logic:


		- Create a Channel name "done" with buffer size 1
			- Why Buffered:
				- If goroutines finishes execution quickly and tries to send to the channel before the select is ready, it won't block(because of buffer).
				- So, this ensure smooth communication between goroutines and main middleware.
			- The Type struct{} is used because it it will take zero memory, We only need a signal not actual data.
			- We start a new goroutine to execute the rest of the Gin handler chain:
				- c.Next() -> execute the next middleware/handler in the pipeline.
				- Once it finishes, we send an empty struct signal (done <- struct{}{}) into channel.
				- This means "hey handler finished"
				- Why Goroutines?
					- Because we want the middleware itself to keep watching the time (main thread), while the handler runs independently.
					- If the handler takes too long, the middleware can abort the request.

			- The Select waits for either:
				- A signal from done channel (handler finish normally).
				- If <- done happens first -> it means handler finishes before timeout -> request proceeds normally.
				- time.After(timeout) -> creates a channel that automatically send a signal after timeout durations.
					- If timeout happens  first:
						- We abort the request first:
						- We abort the request with  c.AbortWithStatusJSON(504, {"error":"request timeout"}).
						- Status code 504 Gateway timeout is used when server didn't respond in time.
				- c.Abort() stops further middleware/handlers from running immediately send the response.


*/

package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SimpleTimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		done := make(chan struct{}, 1)

		go func() {
			c.Next() // run handlers
			done <- struct{}{}
		}()

		select {
		case <-done:
			// handler finish in time
		case <-time.After(timeout):
			// Context canceled (timeout reached)
			c.AbortWithStatusJSON(http.StatusGatewayTimeout, gin.H{
				"error": "request timeout",
			})
		}
	}
}
