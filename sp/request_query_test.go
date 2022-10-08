package sp

import (
	"bytes"
	"testing"
)

func TestNewRequestQuery(t *testing.T) {
	var cases = []struct {
		area     Area
		expected RequestQuery
	}{
		{Area{0x0000a000, 0}, RequestQuery("Q\x00\x00\xa0\x00\x00\x9d\x4b")},
		{Area{0x0000a093, 3}, RequestQuery("Q\x03\x93\xa0\x00\x00\x53\x9d")},
	}

	for _, tt := range cases {
		actual := NewRequestQuery(tt.area)
		if bytes.Equal(actual, tt.expected) {
			continue
		}
		t.Errorf("CreateQueryRequest(%v) => %v, actual %v", tt.area, tt.expected, actual)
	}
}
