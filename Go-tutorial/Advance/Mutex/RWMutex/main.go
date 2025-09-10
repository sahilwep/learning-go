/*

// Mutex vs RWMutex:
	-> Mutex: Only one goroutine can access the critical section at a time(wether reading or writing).
	-> RWMutex:
		-> Multiple goroutine can read concurrently.
		-> But only one goroutine can write, and it blocks everyone else(readers + writers).


		// Example use case:
			-> Imagine you have a cache(like a map) that many goroutines need to read from often, but occasionally Write/update.
			// Using a Mutex:
				-> Every read will block other readers, even though they're only reading, that's inefficient.
			// Using a RWMutex:
				-> Many reader can access the cache at the same time -> faster.
				-> Only when updating the cache(writing), exclusive lock is taken.



// RWMutex:
	-> RWMutex stands for Read-Write Mutex.
	-> It's like a normal Mutex, but with an extra capability.
	-> It allows multiple readers to access a resource at the same time, as long as no one is writing to it.
	-> If a writer lock's in:
		-> No-other writer or reader can access it until the writer unlocks it.
	-> If any writer locks in, they can all read simultaneously (concurrent safe reads).

	// Methods of RWMutex:
		var rw sync.RWMutex

		rw.RLock()		// Acquire a read lock
		rw.RUnlock()	// Acquire a read unlock

		rw.Lock()		// Acquire a write lock
		rw.Unlock		// Acquire a write unlock


	// Summary:
		-> Race Condition = multiple goroutines access shared memory without sync.
		-> Mutex = ensure only one goroutine enters critical section.
		-> Use sync.Mutex for general Protection, sync.RWMutex for read-heavy cases.
		-> Always use go run -race to check for unsafe code.


	// Real World Use Cases:
		-> Protection shared counters(metrics, ID's)
		-> Managing in-memory caches.
		-> Writing to shared log files.
		-> Synchronizing access to maps (Go mas are not safe for concurrent use without sync).


*/

package main

import (
	"fmt"
	"sync"
)

var (
	data = make(map[string]string)
	mu   sync.RWMutex
)

func read(key string, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.RLock()
	fmt.Println("Read:", key, "=", data[key])
	mu.RUnlock()
}

func write(key, value string, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	data[key] = value // Critical Section
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go write("name", "Sahil", &wg)
	go read("name", &wg)
	go read("name", &wg)

	wg.Wait()

}
