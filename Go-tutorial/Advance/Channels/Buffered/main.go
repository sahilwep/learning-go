/*

// Buffered Channel:
	-> Buffered channel in go is a channel that can store a fixed number of values before blocking.
		ch := make(chan int, 5)	// channel size = 5
	-> The channel (3, in this example) determines how many values the channel can hold without receiver.
	-> Send blocks only if the buffer is full.
	-> Receiver blocks only if the buffer is empty.
	-> Buffers channels decouples sender and receiver to some context.

	// How Buffer channel works:
		-> Imagine the channel as a queue with fixed size.
		-> Producer can send values into the queue until it's full.
		-> Consumer can read values from the queue until it's empty.
		-> This is difference from unbuffered channels, where every send must wait for a receiver.

	// Example:

		func example() {
			ch := make(chan int, 3) // buffered channel size = 3
			ch <- 1
			ch <- 2
			ch <- 3

			fmt.Println("All values send without a receiver")

			fmt.Println(<-ch) // 1
			fmt.Println(<-ch) // 2
			fmt.Println(<-ch) // 3
		}


	-> Producer does not block until the buffer is full
	-> After sending 3rd value, the buffer is full; any future send would block.

	// How It works:
		-> Buffered channel size = 3 -> producer can send 3 items before blocking.
		-> Consumer is slower(500ms per items), so the buffer absorbs some values.
		-> After buffer is full, producer blocks until consumer consume a value.
		-> When channel is closed, consumer exits the range loop gracefully.

	// Key Points:
		-> Buffer size matters:
			-> Too small behaves almost like unbuffered
			-> Too large memory overhead.
		-> Synchronization:
			-> Buffer: Strong (send & receive always sync)
			-> Unbuffer: Loose (send can continue until buffer fills)
		-> Closing Behavior:
			-> Even buffered must be closed by the sender to signal no more values.
		-> Deadlocks still possible:
			-> Example: buffered channel is full + no consumer -> send block forever.
		-> Uncase:
			-> Unbuffer: tight coordination, signaling
			-> Buffered: Queuing, decoupling producer/consumer.
		-> Think of like:
			-> Unbuffered: Hand to hand transfer
			-> Buffered: drop into mailbox.


*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Producer will send value to channel:
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		fmt.Println("Producing:", i)
		ch <- i                            // only block if channel is full.
		time.Sleep(200 * time.Millisecond) // taking some time so that our channel will send some value
	}
	close(ch) // closing channel
	fmt.Println("Producer Finished!!")
}

// Consumer will receive values from channel:
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range ch {
		fmt.Println("Consuming:", val)     // printing the values from channel
		time.Sleep(500 * time.Millisecond) // Simulate slow consumer
	}

	fmt.Println("Consumer Finished!!")
}

func bufferExample() {
	ch := make(chan int, 3) // Buffer channel of capacity 3
	var wg sync.WaitGroup

	wg.Add(2)
	go consumer(ch, &wg)
	go producer(ch, &wg)

	wg.Wait()
	fmt.Println("All Goroutines works done!!")
}

// Extra: Deadlock demonstration
func deadlockExample() {
	ch := make(chan int, 3)

	// Buffer capacity = 3
	ch <- 1
	ch <- 2
	ch <- 3

	// At this point buffer is full
	ch <- 4 // blocks forever
	fmt.Println("Never Execute, Deadlock")
}

func main() {
	bufferExample()

	// deadlockExample()	// Deadlock demonstration
}
