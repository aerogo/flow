package flow_test

import (
	"sync/atomic"
	"testing"

	"github.com/aerogo/flow"

	"github.com/stretchr/testify/assert"
)

func TestParallel(t *testing.T) {
	a := int64(0)
	b := int64(0)

	flow.Parallel(func() {
		atomic.AddInt64(&a, 1)
	}, func() {
		atomic.AddInt64(&b, 1)
	})

	assert.Equal(t, int64(1), a)
	assert.Equal(t, int64(1), b)
}

func TestParallelRepeat(t *testing.T) {
	a := int64(0)
	b := int64(0)
	n := 10

	flow.ParallelRepeat(n, func() {
		atomic.AddInt64(&a, 1)
	}, func() {
		atomic.AddInt64(&b, 1)
	})

	assert.Equal(t, int64(n), a)
	assert.Equal(t, int64(n), b)
}
