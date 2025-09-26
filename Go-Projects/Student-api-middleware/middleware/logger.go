/*

// Logger middleware:
	// What is logger middleware:
		- A logger middleware is like a "CCTV camera" for your API.
		- Every request that enters your server -> it records what happened:
		- What endpoint was hit (/student/123)
		- Which method (GET, POST, etc..)
		- When it was called (timestamp)
		- How long it took (latency)
		- What status code you responded with (200, 404, 500)

	// Why do we need it?
		- Imagine you're running a student API in production:
		- someone hit POST /student and gets a 500 Internal Server Error.
		- Without logs, you have no idea what went wrong.
		- With logs, you can see:

			[2025-09-16 20:12] POST /students → 500 (latency: 12ms, IP: 103.42.x.x)
			Error: student not found

	// Where do logs go?
		- In Console - Dev mode
			- In local developments, logs go straight to your console.
			- It's simple, fast, no extra setup.
			- Example:
				POST /students → 201 (12ms)
				GET /students → 200 (3ms)

		- Logs files - Server mode
			- In production, you usually configure your logger to write logs to a file.
			- Example:
				/var/log/student-api/access.log
				/var/log/student-api/error.log


	// Centralized logging System:
		- When app scale to multiple server/containers, logs can't just sit on one machine.
		- That's where log aggregator comes in:
			- ELK stack (Elasticsearch + logstash + Kibana)
			- Grafana loki
			- Datadog
			- Splunk
			- AWS CloudWatch / GCP Stackdriver

	// Let's write a logger() function to write logs for development & productions
		- Directory Structure:

				student-api/
				├─ middleware/
				│   └─ logger.go       # our custom logger
				│
				├─ logs/
				│   ├─ access.log      # will be created automatically
				│   └─ error.log       # can be used later for error logs

	// Program Flow:
		- At app startup (when you call Logger(true) and register it as middleware):
			- It only sets up the logger -> create the file, sets up io.MultiWriter, and return the middleware function.
			- No logs is written yet.
			- At request time (when client make an HTTP request):
				- Gin calls your returned middleware function.
				- that function runs logger.Printf(...).
				- At that exact movement, the logs entry is written into:
					/logs/access.log (file)
					stdout (console)


*/

package middleware

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger(toFile bool) gin.HandlerFunc {
	// Setup logger:
	var logger *log.Logger

	if toFile {
		// Create log directory if not exist:
		if _, err := os.Stat("logs"); os.IsNotExist(err) {
			_ = os.Mkdir("logs", 0755)
		}

		// Open File in append mode:
		file, err := os.OpenFile("logs/access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic("Failed to open log file: " + err.Error())
		}

		// Logger write to both file & console:
		logger = log.New(io.MultiWriter(file, os.Stdout), "LOGGER: ", log.LstdFlags)

	} else {
		// just to log on console:
		logger = log.New(os.Stdout, "Logger: ", log.LstdFlags)
	}

	// Return Gin middleware:
	return func(c *gin.Context) {
		// Before Request:
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Pass to next handler:
		c.Next()

		// After request:
		latency := time.Since(start)
		status := c.Writer.Status()
		clientIP := c.ClientIP()

		logger.Printf("%s %s -> %d (%v) from %s", method, path, status, latency, clientIP) // When this function executed, it will print the logs according to setup above..

	}
}
