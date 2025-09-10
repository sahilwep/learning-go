/*

// Mutex & Race Condition:
	// Race Conditions:
		-> A race condition happens when:
			-> Two or more goroutines access the same shared variable at the same time.
			-> At least one of them write to it.
			-> There's no proper synchronization.
		-> Because golang runs concurrently, the final result can change depending on scheduling, making your program more unpredictable.
		-> Explore "Race-Condition" directory to know more about race condition..

	-> Mutex(Mutual Exclusion) ensure that only one goroutine can access a piece of code(critical section) at a time.
	-> In Go:
		-> "sync.Mutex" is the type.
		-> "Lock()" -> acquired the lock
		-> "Unlock()" -> release the lock


	// Important Note on Mutex:
		-> Always unlock after lock(use "defer mu.Unlock()" right after blocking)
		-> Never try to lock the same mutex twice from the same goroutines -> Will deadlock.
		-> Keep critical section small to avoid performance bottlenecks.
		-> Prefer channels for Communication if possible (Go philosophy: "Don't communicate by sharing memory, share memory by communicating").

		// Example Defer Unlocks:
			mu.Lock()
			defer mu.Unlock()
			counter++	// safe & cleaner

			-> This ensure the lock always released, even function return early.


*/

package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex // Mutex which locks/unlocks the shared resource, until our job is done
var counter int = 0

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mu.Lock()   // lock
		counter++   // Critical Section
		mu.Unlock() // unlock
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&wg)
	go increment(&wg)
	wg.Wait()

	fmt.Println("Final Counter:", counter) // Result always = 1000

}
