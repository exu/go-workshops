package red

import (
	"errors"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPing(t *testing.T) {
	sampleErr := errors.New("sample error")
	tests := map[string]struct {
		storageErr error
		response   string
		err        error
	}{
		"successful": {
			storageErr: nil,
			response:   "pong",
			err:        nil,
		},
		"with db error": {
			storageErr: sampleErr,
			response:   "",
			err:        sampleErr,
		},
	}
	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		storage := &mockStorager{}
		storage.
			On("Cmd", "INCR", []interface{}{"ping:count"}).
			Return(&redis.Resp{
				Err: test.storageErr,
			}).
			Once()
		h := &Handler{
			db: storage,
		}
		response, err := h.Ping()
		assert.Equal(t, test.err, err)
		assert.Equal(t, test.response, response)
		storage.AssertExpectations(t)
	}
}
