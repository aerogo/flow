package jobqueue_test

import (
	"testing"

	"github.com/aerogo/flow/jobqueue"
	"github.com/akyoto/assert"
)

func TestJobQueue(t *testing.T) {
	work := func(input interface{}) interface{} {
		text := input.(string)
		return len(text)
	}

	jobs := jobqueue.New(work)

	jobs.Queue("Hello World")
	jobs.Queue("Test")

	results := jobs.Wait()

	assert.Equal(t, len("Hello World"), results["Hello World"])
	assert.Equal(t, len("Test"), results["Test"])
}
