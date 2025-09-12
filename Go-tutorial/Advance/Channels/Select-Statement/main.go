/*

// Select Statement:
	-> Select is like in go is like switch, but it works only with channel.
	-> It allows a goroutines to wait on "multiple channel operations" simultaneously.
	-> Think of it as "whichever channel is ready first that case will run".

	// Syntax:
		switch {
			case val := <-ch1:
				fmt.Println("Received from ch1:", val)
			case ch2 <- 42:
				fmt.Println("Sent 42 to ch2")
			default:
				fmt.Println("No channel is ready")
		}

		case: must be a send or receive operations on a channel.
		default: runs if no channel is ready immediately (non-blocking).


	// Behaviors of select:
		-> If one channel is ready, that case runs.
		-> If multiple channels are ready, one is picked at random (to avoid starvation).
		-> If no channel is ready:
			-> If there's a default, it runs immediately.
			-> Otherwise, the goroutine blocks until one channel becomes ready.


	//  Real-world use case of select:
		-> Listening to multiple channels (fan-in).
		-> Timeouts (with time.After).
		-> Cancelling signals (with done channel).
		-> Non-blocking send/receive (using default).
		-> Worker pools (waiting on jobs + stop signal).

	// Let's Understand it with Example
		First loop:
			Both channels are empty -> default executes.
			After 1 second, ch1 becomes ready -> that case runs.
			After 2 second, ch2 becomes ready -> that case runs.

	// Core Concept:
		-> default makes select non-blocking.
			-> Avoid deadlock/keep program moving.
		-> If no channel is ready: program keeps moving instead of freezing.
		-> This is how Go avoid deadlock when waiting on uncertain channels.


*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Creating Unbuffer channel:
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: send after 1 second:
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello Sahil"
	}()

	// Goroutine 2: send after 2 second:
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello Prince"
	}()

	// Main listens both channels until -> it receives both the channel:
	receivedCnt := 0
	for receivedCnt < 2 {
		select {
		case msg1 := <-ch1:
			fmt.Println("received:", msg1)
			receivedCnt++
		case msg2 := <-ch2:
			fmt.Println("received:", msg2)
			receivedCnt++
		default:
			fmt.Println("No channel is ready, doing something else...")
			time.Sleep(500 * time.Millisecond) // simulate other task
		}
	}

	fmt.Println("Main Finish!!")
}
