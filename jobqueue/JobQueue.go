package jobqueue

import (
	"runtime"
	"sync"
	"sync/atomic"
)

// Results maps jobs to their corresponding results.
type Results map[interface{}]interface{}

// JobQueue is a queue with inputs and outputs that are saved in a hash map.
type JobQueue struct {
	jobCount    uint64
	jobs        chan interface{}
	done        chan bool
	results     Results
	resultsLock sync.RWMutex
}

// New creates a new job queue.
func New(work func(interface{}) interface{}, bufferSize int) *JobQueue {
	pool := &JobQueue{}
	pool.jobs = make(chan interface{}, bufferSize)
	pool.done = make(chan bool, bufferSize)
	pool.results = make(Results)

	for w := 1; w <= runtime.NumCPU(); w++ {
		go func() {
			for job := range pool.jobs {
				result := work(job)

				pool.resultsLock.Lock()
				pool.results[job] = result
				pool.resultsLock.Unlock()

				pool.done <- true
			}
		}()
	}

	return pool
}

// Queue ...
func (pool *JobQueue) Queue(job interface{}) {
	pool.jobs <- job
	atomic.AddUint64(&pool.jobCount, 1)
}

// Wait ...
func (pool *JobQueue) Wait() Results {
	jobCount := atomic.LoadUint64(&pool.jobCount)

	for i := uint64(0); i < jobCount; i++ {
		<-pool.done
	}

	return pool.results
}
