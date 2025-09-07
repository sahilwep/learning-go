/*

// Goroutines:
	-> Goroutines are lightweight thread managed by the Go runtimes.
	-> They ley you run function concurrently (at the same time) without blocking each other.
	-> Much cheaper than OS threads -> You can spawn thousand of goroutines at the same times.
	-> "go" keyword used to fire goroutines.

	// Key Points:
		-> Start with go funcName() => launches a goroutines
		-> Goroutines run independently of each other.
		-> If main() exits, all goroutines stop immediately (That's why we used time.Sleep to wait).
		-> To properly sync, WaitGroups and Channels used.

	// Real Backend Use Case:
		-> Imagine handling a web request
		-> Goroutines fetches data from DB
		-> Another Goroutines call for an external API
		-> Both runs at the same time, making the response much faster.


	// In Sort:
		-> Goroutines are super cheap threads
		-> They allows massive Concurrency.
		-> Foundation for Go's scalability in backend & cloud apps.

	// In layman's' Term:
		-> We can run multiple tasks at the same time using Goroutines, which makes overall application super fast.


*/

package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond) // taking 500ms time for every iterations.
	}
}

func main() {

	// Normal function call (runs synchronous)
	printMessage("Synchronous")

	// Goroutines (runs Concurrently)
	go printMessage("1st Concurrent")
	go printMessage("2nd Concurrent")

	// Give Goroutines time to finish:
	time.Sleep(3 * time.Second)
	fmt.Println("Main Finished")
}
