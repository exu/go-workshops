package astitime_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/asticode/go-astitools/time"
	"github.com/stretchr/testify/assert"
)

func TestTimestamp(t *testing.T) {
	var tsp = astitime.Timestamp{}
	err := json.Unmarshal([]byte("1495290215"), &tsp)
	assert.NoError(t, err)
	assert.Equal(t, time.Unix(1495290215, 0), tsp.Time)
	b, err := json.Marshal(tsp)
	assert.NoError(t, err)
	assert.Equal(t, "1495290215", string(b))
}
