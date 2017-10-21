package jobqueue

import (
	"runtime"
	"sync"
)

// Results maps jobs to their corresponding results.
type Results map[interface{}]interface{}

// JobQueue is a queue with inputs and outputs that are saved in a hash map.
type JobQueue struct {
	jobs        chan interface{}
	wg          sync.WaitGroup
	results     Results
	resultsLock sync.RWMutex
}

// New creates a new job queue.
func New(work func(interface{}) interface{}) *JobQueue {
	pool := &JobQueue{}
	pool.jobs = make(chan interface{})
	pool.results = make(Results)

	for w := 1; w <= runtime.NumCPU(); w++ {
		go func() {
			for job := range pool.jobs {
				result := work(job)

				pool.resultsLock.Lock()
				pool.results[job] = result
				pool.resultsLock.Unlock()

				pool.wg.Done()

				// pool.done <- true
			}
		}()
	}

	return pool
}

// Queue ...
func (pool *JobQueue) Queue(job interface{}) {
	pool.wg.Add(1)
	pool.jobs <- job
}

// Wait ...
func (pool *JobQueue) Wait() Results {
	pool.wg.Wait()
	return pool.results
}
