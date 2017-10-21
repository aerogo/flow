package flow_test

import (
	"testing"

	"github.com/aerogo/flow"

	"github.com/stretchr/testify/assert"
)

func TestParallel(t *testing.T) {
	a := 0
	b := 0

	flow.Parallel(func() {
		a = 13
	}, func() {
		b = 10
	})

	assert.Equal(t, 13, a)
	assert.Equal(t, 10, b)
}
