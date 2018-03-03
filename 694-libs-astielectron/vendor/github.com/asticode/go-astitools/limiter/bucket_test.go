package astilimiter_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBucket(t *testing.T) {
	var l = astilimiter.New()
	var b = l.Add("test", 2, time.Second)
	assert.True(t, b.Inc())
	assert.True(t, b.Inc())
	assert.False(t, b.Inc())
}
