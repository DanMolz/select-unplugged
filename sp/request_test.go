package sp

import (
	"bytes"
	"testing"
)

func TestCreateQueryRequest(t *testing.T) {
	var cases = []struct {
		area     Area
		expected Request
	}{
		{Area{0x0000a000, 0}, Request("Q\x00\x00\xa0\x00\x00\x9d\x4b")},
		{Area{0x0000a093, 3}, Request("Q\x03\x93\xa0\x00\x00\x53\x9d")},
	}

	for _, tt := range cases {
		actual := CreateQueryRequest(tt.area)
		if bytes.Equal(actual, tt.expected) {
			continue
		}
		t.Errorf("CreateQueryRequest(%v) => %v, actual %v", tt.area, tt.expected, actual)
	}
}
