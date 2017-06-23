package flow

import "sync"

// Parallel starts all functions asynchronously as goroutines and waits until they are completed.
func Parallel(funcs ...func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(funcs))

	for _, fun := range funcs {
		// Bind iterator to a local variable so we can capture it in the closure.
		task := fun

		go func() {
			task()
			wg.Done()
		}()
	}

	wg.Wait()
}
