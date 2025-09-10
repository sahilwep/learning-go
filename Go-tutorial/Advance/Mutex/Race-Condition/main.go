/*

// Example of race conditions:
	-> We have used a variable counter = 0
	-> We designed function increment() which will increment the counter value by 1000 whenever it calls.
	-> We have fire up two concurrent goroutines, which will increment value by 2000, but due the modification in shared variable, value will not be as intended.
	-> Expected output: 2000
	-> Actual output: 1391, 1630, 1093...


	// Go Race Detector:
		-> Go provides race conditions detector.
		-> run "go run -race main.go" to check.
		-> If there's a race, this will report.

		// Race condition checker Output:
			==================
			WARNING: DATA RACE
			Read at 0x000104bf70d0 by goroutine 8:
			main.increment()
			....
			...
			==================
			Counter: 2000
			Found 2 data race(s)
			exit status 66


*/

package main

import (
	"fmt"
	"sync"
)

var counter = 0

func increment(wg *sync.WaitGroup) {
	defer wg.Done() // execute at the end of this function

	for i := 0; i < 1000; i++ {
		counter++
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&wg)
	go increment(&wg)
	wg.Wait()

	fmt.Println("Counter:", counter)
}
