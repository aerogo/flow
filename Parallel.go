package flow

import "sync"

// Parallel starts all functions asynchronously as goroutines and waits until they are completed.
func Parallel(funcs ...func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(funcs))

	for _, fun := range funcs {
		go func(task func()) {
			task()
			wg.Done()
		}(fun)
	}

	wg.Wait()
}

// ParallelRepeat starts the function asynchronously as goroutines n times and waits until they are completed.
func ParallelRepeat(times int, funcs ...func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(funcs) * times)

	for i := 0; i < times; i++ {
		for _, fun := range funcs {
			go func(task func()) {
				task()
				wg.Done()
			}(fun)
		}
	}

	wg.Wait()
}
