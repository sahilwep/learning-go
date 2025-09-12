/*

// Fan-out:
	-> Fan-out is a concurrency pattern where one channel feeds multiple goroutines, often called workers, so they can process tasks in parallel.
	-> One Producer: many consumers.
	-> Helps scale processing & reduce latency.
	-> Typically used in worker pools.

	// NOTE:
		-> Fan-in & Fan-out they are not opposite, Instead they are design pattern & very different from each other, "{{{it's a design pattern}}}"


	// Quick Comparison: Fan-in | Fan-out
		Fan-in:
			multiple channels are merged into single channels.
			purpose: aggregate result from multiple source.
			many producer: one consumer.
			Message may interleave; order not guarantee.
			Often used to collect result from parallel workers.
			Example:
				Worker1 channel -> merged channel
				Worker2 channel -> merged channel
				Worker3 channel -> merged channel


		Fan-out:
			One channel feeds multiple goroutines(worker) to process task concurrently.
			Purpose: parallelize work and reduce latency.
			one producer: multiple consumer
			Each one pulls from same input channel.
			Often used in worker pools.
			Example:
				jobs channel -> Worker1
							 -> Worker2
							 -> Worker3


	// Simple Backend Example: Worker pool
		-> Imagine a backend scenario: you have a list of jobs(API request, tasks, or database queries), and you want to process them concurrently using multiple workers.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function: process jobs from the jobs channel:
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}

}

func main() {
	const numWorkers = 3
	const numJobs = 6

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers (fan-out)
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs into jobs channel:
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Important: Close so workers know no more jobs are coming.

	// Wait for all worker to finish
	wg.Wait()
	fmt.Println("All jobs Processed.")

}
