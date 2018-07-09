package astilimiter

import (
	"sync"
	"time"
)

// Limiter represents a limiter
type Limiter struct {
	buckets map[string]*Bucket
	m       *sync.Mutex // Locks buckets
}

// New creates a new limiter
func New() *Limiter {
	return &Limiter{
		buckets: make(map[string]*Bucket),
		m:       &sync.Mutex{},
	}
}

// Add adds a new bucket
func (l *Limiter) Add(name string, cap int, period time.Duration) *Bucket {
	l.m.Lock()
	defer l.m.Unlock()
	if _, ok := l.buckets[name]; !ok {
		l.buckets[name] = newBucket(cap, period)
	}
	return l.buckets[name]
}

// Bucket retrieves a bucket from the limiter
func (l *Limiter) Bucket(name string) (b *Bucket, ok bool) {
	l.m.Lock()
	defer l.m.Unlock()
	b, ok = l.buckets[name]
	return
}
