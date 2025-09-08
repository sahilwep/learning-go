/*

// Backend Style Example:
	-> Imagine a backend API fetching data from multiple microservices.
	-> If we fetch them sequentially, it takes longer
	-> With Goroutines + WaitGroup -> All run parallels & Super fast.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchFromService(name string, wg *sync.WaitGroup) {
	defer wg.Done() // Execute at the end of this fetchFromService function
	fmt.Println("Fetching", name)
	time.Sleep(500 * time.Millisecond) // simulate request
	fmt.Println("Done:", name)
}

func main() {
	services := []string{"Users", "Orders", "Payments", "Notifications"}
	var wg sync.WaitGroup

	for _, s := range services {
		wg.Add(1)
		go fetchFromService(s, &wg)
	}

	wg.Wait()
	fmt.Println("All Services responded")
}
