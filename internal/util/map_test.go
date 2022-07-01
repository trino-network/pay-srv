package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trino-network/pay-srv/internal/util"
)

func TestMergeStringMap(t *testing.T) {
	baseMap := map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
	}

	toMerge := map[string]string{
		"C": "1",
		"D": "2",
	}

	expected := map[string]string{
		"A": "a",
		"B": "b",
		"C": "c",
		"D": "2",
	}

	res := util.MergeStringMap(baseMap, toMerge)
	assert.Equal(t, expected, res)

	expected = map[string]string{
		"C": "1",
		"D": "2",
		"A": "a",
		"B": "b",
	}

	res = util.MergeStringMap(toMerge, baseMap)
	assert.Equal(t, expected, res)
}
