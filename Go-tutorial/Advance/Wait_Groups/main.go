/*

// Wait Group:
	-> WaitGroup (from the sync package) is counter that lets one goroutine wait until a collection of goroutines finish their executions.
	-> You can think of it like a task counter:
		-> You "add" a tasks `before` starting goroutine.
		-> Each goroutine `marks done` when it finishes
		-> The main goroutine "waits" until the counter goes back to zero.

	//  Function of WaitGroup:
		-> The type is "sync.WaitGroup", and it has 3 main methods:
			1. "Add(n int)" increment counter by 'n'.
				-> You usually call "wg.Add(1)" before launching a goroutines.
			2. "Done()" Decrement the counter by '1'.
				-> Usually deferred inside the goroutine.
			3. "Wait()" Block until counter reaches '0'.
				-> Called in main goroutine or on the one coordinating work.

	// Important Rule:
		-> "Add()" must be called before the goroutine start, otherwise race condition happen.
		-> Counter must never go negative (panic if Done() is called too many times).
		-> WaitGroup is not reusable like a rest - once counter reaches zero, you can reuse it by adding task again, but you cannot reset "mid-flight".
		-> Always pass &wg(pointer), not a copy.


	// Example of Race Condition:
			for i := 1; i <= 5; i++ {
				go func() {
					wg.Add(1)
					defer wg.Done()
					fmt.Println("Working")
				}()
			}
			wg.Wait()

		This may panic because:
			Goroutine may call Add(1) after wait() has already started.
			Solution -> Always call Add() before starting goroutines.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Any Random Function which uses WaitGroup:
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrease counter when this done, as we know defer called at the end of the function
	fmt.Printf("Worker %d started \n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d finished \n", id)
}

func main() {
	var wg sync.WaitGroup // create a variable of sync.WaitGroup type

	for i := 1; i <= 3; i++ {
		wg.Add(1) // tell WaitGroup there's a new Goroutine, & Make user to add this before the goroutine.
		go worker(i, &wg)
	}

	wg.Wait() // wait for all workers
	fmt.Println("All workers finished")

}
