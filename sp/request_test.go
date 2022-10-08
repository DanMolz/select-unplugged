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

func TestCreateWriteRequest(t *testing.T) {
	for _, tt := range []struct {
		memory   Memory
		expected Request
	}{
		{
			Memory{
				address: 0x001f0000,
				data:    Data("\xb6\xd1\x36\x04\x08\x0c\x87\xce\x81\xc1\x82\xc6\x6f\xa5\xfb\x35"),
			},
			Request("W\x07\x00\x00\x1f\x00\x35\x7a\xb6\xd1\x36\x04\x08\x0c\x87\xce\x81\xc1\x82\xc6\x6f\xa5\xfb\x35w\xaa"),
		},
	} {
		actual := CreateWriteRequest(tt.memory)
		if bytes.Equal(actual, tt.expected) {
			continue
		}
		t.Errorf("CreateWriteRequest(%v) => %v, actual %v", tt.memory, tt.expected, actual)
	}
}
