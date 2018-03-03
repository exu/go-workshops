package astiio_test

import (
	"bytes"
	"sync"
	"testing"

	"github.com/asticode/go-astitools/io"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

// MockedReader is a mocked io.Reader
type MockedReader struct {
	buf      *bytes.Buffer
	infinite bool
}

// NewMockedReader creates a new mocked reader
func NewMockedReader(i string, infinite bool) MockedReader {
	return MockedReader{buf: bytes.NewBuffer([]byte(i)), infinite: infinite}
}

// Read allows MockedReader to implement the io.Reader interface
func (r MockedReader) Read(p []byte) (n int, err error) {
	if r.infinite {
		return
	}
	n, err = r.buf.Read(p)
	return
}

func TestCopy(t *testing.T) {
	// Init
	var w = &bytes.Buffer{}
	var r1, r2 = NewMockedReader("testiocopy", true), NewMockedReader("testiocopy", false)

	// Test cancel
	var nw int64
	var err error
	var ctx, cancel = context.WithCancel(context.Background())
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		nw, err = astiio.Copy(ctx, r1, w)
	}()
	cancel()
	wg.Wait()
	assert.EqualError(t, err, "context canceled")

	// Test success
	w.Reset()
	ctx = context.Background()
	nw, err = astiio.Copy(ctx, r2, w)
	assert.NoError(t, err)
	assert.Equal(t, "testiocopy", w.String())
}
