package astitime_test

import (
	"sync"
	"testing"
	"time"

	"github.com/asticode/go-astitools/time"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestSleep(t *testing.T) {
	var ctx, cancel = context.WithCancel(context.Background())
	var err error
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = astitime.Sleep(ctx, time.Minute)
	}()
	cancel()
	wg.Wait()
	assert.EqualError(t, err, "context canceled")
}
