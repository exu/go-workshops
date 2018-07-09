package astibinary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryWriter(t *testing.T) {
	w := New()
	w.Write("1100")
	assert.Empty(t, w.Bytes())
	assert.Equal(t, "1100", w.remainingBits)
	w.Write("11001")
	assert.Equal(t, []byte{0xcc}, w.Bytes())
	assert.Equal(t, "1", w.remainingBits)
	w.Reset()
	assert.Empty(t, w.Bytes())
	w.Write("11")
	w.Write(uint8(170))
	assert.Equal(t, []byte{0xea}, w.Bytes())
	assert.Equal(t, "10", w.remainingBits)
	w.Write([]byte{0xcc, 0xcd})
	assert.Equal(t, []byte{0xea, 0xb3, 0x33}, w.Bytes())
	assert.Equal(t, "01", w.remainingBits)
}
