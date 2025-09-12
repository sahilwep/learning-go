/*

// Closing a channel:
	-> We can use "close(channelName)" built-in function.


*/

package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	// Producer:
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch) // closing after sending all values.
	}()

	// Consumer:
	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("Channel Closed, consumer finished")
}
