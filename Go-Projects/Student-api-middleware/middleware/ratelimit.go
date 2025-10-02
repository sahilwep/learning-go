/*

// Rate Limiting:
	// What is Rate Limiting:
		- Rate limiting is a technique to control the number of request client can make to your API in a certain period.
		- Think of it as a traffic signal for your backend.
		- It prevent abuse, overload, or accidental flooding of your server.

	// Real-World Example:
		// Twitter API:
			- Limit you to 300 tweets per 3 hours.
			- Without it, a bot could spam millions of tweets and crash the system.
		// Login Endpoints:
			- Limit Login attempts to 5 per minute per IP.
			- Prevents bute-force attacks.

		// Student API(your project):
			- Imagine hundreds of clients calling / student endpoints repeatedly.
			- Rate limiting can prevent your server from begin overwhelmed.

	// Why we use it:
		// Prevent server overload:
			- Even a simple in-memory API can crash if hundreds of requests hit it simlationously.
		// Security:
			- Stops brute-force, spam or DDoS attempts.
		// Fair usage:
			- Ensure no single client hogs all resources.
		// Predictable performance:
			- Server can maintain stable response times.

	// How it works:
		- Each client (usually tracked by IP of API key) is allowed N request per duration.
		- If they exceed the limit, the server returns HTTP 429 Too Many Requests.
		// Popular Algorithms:
			// Fixed Window Counter:
				- Example: 10 requests per 1 minutes windows.
				- Easy but "burst" at window edges.
			// Sliding window log:
				- Tracks exact timestamp of requests.
				- More precise, prevents bursts.
			// Token Bucket / Leaky Buckets:
				- Requests are allowed if "tokens" are available.
				- Smooths out bursts and allows some flexibility.

	// Real backend Scenario:
		// Imagine your student API is publicly available:
			// Without rate limiting:
				- A single script can call POST / student 1000 times per second.
				- Memory usage spikes -> Server slows -> Potential Crash.
			// With rate limiting:
				- Allow 5 requests per second per client.
				- Extra request get 429 Too many Requests, server remains stable.

	// How the middleware works internally (conceptually).
		- Identify the client (IP or API key).
		- Track how many requests they made in the time window.
		- If under the limit -> call c.Next() (allow the request).
		- If over the limit -> c.AbortWithStatusJSON(429, ...) (block request).
		// Key Points:
			- Rate limiting is server-side, it does not rely on the client to behave.
			- Works for all routers or specific routes (like login, signup).
			- Can be memory-based (simple map) or Redis-based for distribution app.


	// Why it's important even for small apps:
		- Even a small app can be accidentally abused by client or testing scripts.
		- Prevents wasting CPU, memory and bandwidth.
		- Sets the foundations for scaling the app in the future.


*/

package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Clients holds information about requests from a single IP
type Client struct {
	Request  int
	LastSeen time.Time
}

func RateLimiterMiddleware(maxRequests int, window time.Duration) gin.HandlerFunc {
	clients := make(map[string]*Client)
	var mu sync.Mutex

	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		mu.Lock()
		client, exists := clients[ip]

		if !exists || now.Sub(client.LastSeen) > window {
			// New client or window expired
			clients[ip] = &Client{
				Request:  1,
				LastSeen: now,
			}
			mu.Unlock()
			c.Next()
			return
		}

		// 	Existing Client Within window:
		if client.Request >= maxRequests {
			mu.Unlock()
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "rate limit exceeds, try again later",
			})
			return
		}

		client.Request++
		client.LastSeen = now
		mu.Unlock()

		c.Next()
	}

}
