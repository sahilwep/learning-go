/*

// Unbuffered Channel:
	-> A unbuffered channel in go is a channel without capacity, Create like:
		ch := make(chan int)	// unbuffered
	-> No internal storage - The channel cannot hold any value.
	-> Send block until a receiver is ready, and receive block until sender is ready.
	-> Guarantee synchronization between goroutines - data transferred only when both side are ready.


	// Why use unbuffered channels:
		-> Unbuffered channels are perfect when you want strict synchronization between goroutines
		-> The sender waits until the receiver is ready.
		-> Guaranteeing ordering
		-> No values are stored, so you know exactly when a value is passed.
		-> Coordination of events useful for singling between goroutines.

*/

package main

import (
	"fmt"
	"sync"
)

// Producer send number into channel then closed it.
func producer(ch chan<- int, count int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= count; i++ {
		ch <- i
		fmt.Println("Sending:", i)
	}

	close(ch)
	fmt.Println("Producer Finished!")
}

// Consumer receive value from channel until it's closed.
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("received:", val)
	}

	fmt.Println("Consumer done (Channel Closed)")
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	// Run 2 concurrent goroutines:
	wg.Add(2)
	go producer(ch, 3, &wg)
	go consumer(ch, &wg)

	wg.Wait()
	fmt.Println("All goroutines done!!")
}
