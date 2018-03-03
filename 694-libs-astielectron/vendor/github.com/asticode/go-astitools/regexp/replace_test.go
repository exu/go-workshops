package astiregexp_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/asticode/go-astitools/regexp"
	"github.com/stretchr/testify/assert"
)

func TestReplaceAll(t *testing.T) {
	// Initialize
	rgx := regexp.MustCompile("{[A-Za-z0-9_]+}")
	src := []byte("/test/{m1}/test/{ma2}/test/{match3}")
	rpl := []byte("valuevaluevaluevaluevaluevaluevalue")

	// Replace all
	astiregexp.ReplaceAll(rgx, &src, rpl)

	// Assert
	assert.Equal(t, fmt.Sprintf("/test/%s/test/%s/test/%s", string(rpl), string(rpl), string(rpl)), string(src))
}
