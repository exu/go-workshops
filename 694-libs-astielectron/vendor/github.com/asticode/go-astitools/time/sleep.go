package astitime

import (
	"context"
	"time"
)

// Sleep is a cancellable sleep
func Sleep(ctx context.Context, d time.Duration) (err error) {
	for {
		select {
		case <-time.After(d):
			return
		case <-ctx.Done():
			err = ctx.Err()
			return
		}
	}
	return
}
