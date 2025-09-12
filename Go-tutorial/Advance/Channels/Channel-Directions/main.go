/*

// Channel Directions:
	-> In go channels bidirectional (default), or directional(restrict data send or receive).

	// Bidirectional Channel:
		var ch chan int
		ch := make(chan int)	// or create with make()

		ch <- 10   // send
		val := <-ch // receive


	// Directional Channel:

		// Send Only:
			var ch chan<- int
			ch <- 10   		// allowed
			val := <-ch 	// Not allowed

		// Receive Only:
			var ch <-chan int
			val := <-ch 	// allowed
			ch <- 10    	// not allowed


	// Visual Summary:
		Channel_Type		Send Allowed			Received Allowed
		chan T					Yes						Yes
		chan<-					Yes						No
		<-chan					No						Yes

	-> Extra:
		-> When we pass channel into function, we can restrict their actions, by changing their parameter ordering..
		-> Okey, let's see how to define function with channel restrictions...


*/

package main

import "fmt"

// Bidirectional Channel:
func process(ch chan int) {
	ch <- 10    // send
	val := <-ch // receive
	fmt.Println(val)
	close(ch)
}

// Send-only channel:
func producer(ch chan<- int) {
	ch <- 10 // send
	close(ch)
}

// Receive-only channel:
func consumer(ch <-chan int) {
	val := <-ch // receive
	fmt.Println(val)
}

func main() {

	// Bidirectional:
	ch1 := make(chan int)
	go process(ch1)

	// Send-only:
	ch2 := make(chan int)
	go producer(ch2)

	// Receive only:
	ch3 := make(chan int)
	go consumer(ch3)
	ch3 <- 99
	close(ch3)

}
