/*

// Channels:
	-> A Channels is type conduit that allows goroutines to communicate & synchronize by sending and receiving values.
	-> It's Go implementation of CSP(Communicating sequential processes).
	-> They allows goroutines to synchronize without explicit locks (like mutexes).
	-> Think of it as a pipe:
		-> One goroutines sends value in it.
		-> Another goroutines receives value from it.
	-> Channels ensure safe data sharing without needing explicit lock like mutexes.

	// Declaring Channels:
			ch := make(chan int)	// unbuffered channel of int
			chanBuf := make(chan int, 5)	// buffered channel of int capacity of 5

		-> Channels are typed: 'chan int', 'chan string', etc.
		-> Zero value of channel is nil(cannot be used until initialization with make)

	// Sending & Receiving:
			ch := make(chan int)

			// Sending
			ch <- 10

			// Receiving:
			val := <-ch
			fmt.println(val)

		-> Send & receive are block until the other side is ready (unless the channel is buffered)
		-> This is why channels are synchronize.

		func channelDemonstration() {
			ch := make(chan int)  // Creating unbuffered channel
			var wg sync.WaitGroup // Creating waitGroup to synchronize concurrent goroutines.

			wg.Add(2)
			go func() {
				defer wg.Done()
				ch <- 5 // send data to channel
			}()

			go func() {
				defer wg.Done()
				val := <-ch // receive data from channel
				fmt.Println(val)
			}()

			wg.Wait() // wait until all concurrent services finish
			fmt.Println("Two concurrent services Interacting")
		}


	// Let's Understand with better producer & consumer example:
		-> Producer will send data to channel
		-> Consumer will receive data from channel
		-> NOTE: You will notice the 'chan' position will change <- with arrow, on sending & receiving case
		-> Then we will close the channel when our operations is done.


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

	close(ch) // Important: close the channel when it's done.
	fmt.Println("Producer Finished")
}

// Consumer receive value from channel until it's closed.
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch { // automatically stops when channel is closed
		fmt.Println("Received:", val)
	}

	fmt.Println("Consumer Finished (channel closed)")
}

func main() {
	ch := make(chan int) // unbuffered channel
	var wg sync.WaitGroup

	// Run 2 concurrent goroutine.
	wg.Add(2)
	go producer(ch, 3, &wg)
	go consumer(ch, &wg)

	wg.Wait()
	fmt.Println("All goroutines done!!")
}
