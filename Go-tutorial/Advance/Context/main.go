/*
// context.Context:
	-> context.Context is a built-in Go type used to carry deadlines, cancellation signals, and request-scoped values across API boundaries and goroutines.
	-> It's immutable - you derive new context from existing ones.
	-> Commonly used in HTTP servers, database queries, and long-running goroutines to propagate cancellations.

	// Key Functions:
		context.Background()
			-> Root context, usually for main or top-level goroutines.

		context.TODO()
			-> placeholder with you are not yet sure which context to use.

		context.WithCancel(parent)
			-> Returns a derived context + a cancel function.
			-> Calling cancel() signals all goroutines using that context to stop.

		context.WithTimeout(parent, duration)
			-> Automatically cancels the context after the timeout.

		context.WithDeadline(parent, time)
			-> Cancels at a specific deadline.

		context.WithValue(parent, key, value)
			-> Carries request-scoped values (e.g., user ID, tokens).


	// Backend Use case:
		Imagine a backend service that calls multiple microservices concurrently:
		You create a context with timeout for the request.
		Pass it to all downstream calls.
		If any service is slow, the context cancels all the remaining goroutines automatically.
		Avoids goroutine leaks and long-running stuck operations.


	// Example:
		-> Let's understand context with backend scenario, where we have 3 microservices:
			-> User services, Payment, Notification services.
				-> User Service will take 1 sec(fast)
				-> Payment Service will take 4 sec(moderate)
				-> Notification service will take 2 sec(fast),
			-> But never executes fully,because the payment service trigger timeout first.
		-> This Demonstrate how a one slow service can cause cancellation of all other works.
		-> The example down here, We have is timeout at 3 sec if any microservice will take more than 3 second, then that services will not be executed...


*/

package main

import (
	"context"
	"fmt"
	"time"
)

// ---------------- Microservices ----------------

// Simulate User service (1 sec - fast)
func userService(ctx context.Context, result chan<- string) {
	select {
	// This branch simulate the service completing its work in 1s
	case <-time.After(1 * time.Second):
		result <- "User Data fetch"
	// This branch will run if the context gets canceled
	case <-ctx.Done():
		// ctx.Err() gives the reason (deadline executed or canceled)
		fmt.Println("User service canceled - ", ctx.Err())
		return
	}
}

// Simulate Payment Service (4 sec - very slow)
func paymentService(ctx context.Context, result chan<- string) {
	select {
	// This branch simulate the service completing its work in 4s
	case <-time.After(4 * time.Second):
		result <- "Payment Done"
	// This branch will run if the context gets canceled
	case <-ctx.Done():
		// ctx.Err() gives the reason (deadline executed or canceled)
		fmt.Println("Payment Service canceled - ", ctx.Err())
		return
	}
}

// Simulate Notifications Service (2 sec - moderate)
func notificationService(ctx context.Context, result chan<- string) {
	select {
	// This branch simulate the service completing its work in 2s
	case <-time.After(2 * time.Second):
		result <- "Notification Sent"
	// This branch will run if the context gets canceled
	case <-ctx.Done():
		// ctx.Err() gives the reason (deadline executed or canceled)
		fmt.Println("Notifications Service canceled - ", ctx.Err())
		return
	}
}

func main() {
	// API request with 3s timeout: This means the whole API request must finish in 3s, otherwise cancel everything.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // Always call cancel to free resources

	// Channel to collect results from services:
	requestChan := make(chan string, 3) // buffered (3) because we expect 3 results at most

	// Start all services parallel:
	go userService(ctx, requestChan)
	go paymentService(ctx, requestChan)
	go notificationService(ctx, requestChan)

	// Collect results from 3 services:
	for i := 0; i < 3; i++ {
		select {
		// If any service writes a result, receive and print it
		case res := <-requestChan:
			fmt.Println("API Gateway Received:", res)
		case <-ctx.Done():
			fmt.Println("API Gateway timeout:", ctx.Err())
			return
		}
	}

	// If all services finishes in time (impossible here), API Gateway responds successfully:
	fmt.Println("API Gateway: Response send to client")
}
