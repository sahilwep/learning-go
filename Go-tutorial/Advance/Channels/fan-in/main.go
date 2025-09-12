/*

// Fan-in:
	-> Fan-in is a concurrency pattern where multiple channels are combined into single channel so single goroutine can receive all messages from multiple producers.
	-> Many Producer: 1 Receiver
	-> Useful for margining results or aggregating outputs.
	-> Messages from different channels may interleave; no guaranteed order.

	// Example:
		-> Imagine you are building a service that fetches data from multiple microservices concurrently,
		-> and you want to aggregate all result into a single response.

			ch1 ----\
			         \
			          --> out --> consumer
			ch2 ----/


*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulate microservice calls
func service1() <-chan string { // let say this will fetch User details
	ch := make(chan string)

	go func() {
		time.Sleep(500 * time.Millisecond) // simulate network delay
		ch <- "User: Sahil"
		close(ch)
	}()

	return ch
}

func service2() <-chan string { // let say this will fetch User Bank balance
	ch := make(chan string)

	go func() {
		time.Sleep(700 * time.Millisecond) // simulate network delay
		ch <- "User has: 999.99B $"
		close(ch)
	}()

	return ch
}

// Fan-in: merge multiple channel into one.
func merge(ch1, ch2 <-chan string) <-chan string {
	out := make(chan string) // this will collect all message form ch1 and ch2

	// Use WaitGroup to sync all the incoming channels into single channel.
	var wg sync.WaitGroup
	wg.Add(2)

	// Forwarded message from channel 1
	go func() {
		defer wg.Done()
		for v := range ch1 {
			out <- v
		}
	}()

	// Forwarded message from channel 1
	go func() {
		defer wg.Done()
		for v := range ch2 {
			out <- v
		}
	}()

	// Close out after all forwarders are done
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {

	// Call service concurrently:
	ch1 := service1()
	ch2 := service2()

	// Merge their outputs:
	merged := merge(ch1, ch2)

	fmt.Println("Aggregating data from multiple services: ")
	for msg := range merged {
		fmt.Println(msg)
	}

	fmt.Println("All services Processed..")
}
