package util_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trino-network/pay-srv/internal/util"
)

func TestMinAndMapInt(t *testing.T) {
	max := math.MaxInt32
	min := math.MinInt32
	assert.Equal(t, max, util.MaxInt(max, min))
	assert.Equal(t, max, util.MaxInt(min, max))
	assert.Equal(t, min, util.MinInt(max, min))
	assert.Equal(t, min, util.MinInt(min, max))
	assert.Equal(t, min, util.MaxInt(min, min))
	assert.Equal(t, max, util.MinInt(max, max))
}
