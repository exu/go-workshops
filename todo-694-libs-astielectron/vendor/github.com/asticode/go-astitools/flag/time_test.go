package astiflag_test

import (
	"testing"
	"time"

	"github.com/asticode/go-astitools/flag"
	"github.com/stretchr/testify/assert"
)

func TestTime(t *testing.T) {
	const tm = "2017-12-13T12:34:56+02:00"
	ft := &astiflag.Time{}
	err := ft.Set(tm)
	assert.NoError(t, err)
	at, _ := time.Parse(time.RFC3339, tm)
	assert.Equal(t, at, ft.Time)
}
