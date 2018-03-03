package asticontext_test

import (
	"sync"
	"testing"

	"github.com/asticode/go-astitools/context"
	"github.com/stretchr/testify/assert"
)

func TestCanceller_Cancel(t *testing.T) {
	var c = asticontext.NewCanceller()
	var ctx1, cancel1 = c.NewContext()
	var ctx2, cancel2 = c.NewContext()
	defer cancel1()
	defer cancel2()
	var wg = &sync.WaitGroup{}
	wg.Add(2)
	var count int
	go func() {
		for {
			select {
			case <-ctx1.Done():
				count += 1
				wg.Done()
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx2.Done():
				count += 2
				wg.Done()
				return
			}
		}
	}()
	c.Cancel()
	wg.Wait()
	assert.Equal(t, 3, count)
	assert.True(t, c.Cancelled())
	c.Reset()
	assert.False(t, c.Cancelled())
}
