package astilimiter

import (
	"time"
)

// Bucket represents a bucket
type Bucket struct {
	cap         int
	channelQuit chan bool
	count       int
	period      time.Duration
}

// newBucket creates a new bucket
func newBucket(cap int, period time.Duration) (b *Bucket) {
	b = &Bucket{
		cap:         cap,
		channelQuit: make(chan bool),
		count:       0,
		period:      period,
	}
	go b.tick()
	return
}

// Inc increments the bucket count
func (b *Bucket) Inc() bool {
	if b.count >= b.cap {
		return false
	}
	b.count++
	return true
}

// tick runs a ticker to purge the bucket
func (b *Bucket) tick() {
	var t = time.NewTicker(b.period)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			b.count = 0
		case <-b.channelQuit:
			return
		}
	}
}

// close closes the bucket properly
func (b *Bucket) close() {
	if b.channelQuit != nil {
		close(b.channelQuit)
		b.channelQuit = nil
	}
}
