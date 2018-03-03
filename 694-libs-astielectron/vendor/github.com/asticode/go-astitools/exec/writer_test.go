package astiexec_test

import (
	"testing"

	"github.com/asticode/go-astitools/exec"
	"github.com/stretchr/testify/assert"
)

func TestStdWriter(t *testing.T) {
	// Init
	var o []string
	var w = astiexec.NewStdWriter(func(i []byte) {
		o = append(o, string(i))
	})

	// No EOL
	w.Write([]byte("bla bla "))
	assert.Empty(t, o)

	// Multi EOL
	w.Write([]byte("bla \nbla bla\nbla"))
	assert.Equal(t, []string{"bla bla bla ", "bla bla"}, o)

	// Close
	w.Close()
	assert.Equal(t, []string{"bla bla bla ", "bla bla", "bla"}, o)
}
