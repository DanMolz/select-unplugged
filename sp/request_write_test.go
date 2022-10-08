package sp

import (
	"bytes"
	"testing"
)

func TestNewRequestWrite(t *testing.T) {
	for _, tt := range []struct {
		memory   Memory
		expected RequestWrite
	}{
		{
			Memory{
				address: 0x001f0000,
				data:    Data("\xb6\xd1\x36\x04\x08\x0c\x87\xce\x81\xc1\x82\xc6\x6f\xa5\xfb\x35"),
			},
			RequestWrite("W\x07\x00\x00\x1f\x00\x35\x7a\xb6\xd1\x36\x04\x08\x0c\x87\xce\x81\xc1\x82\xc6\x6f\xa5\xfb\x35w\xaa"),
		},
	} {
		actual := NewRequestWrite(tt.memory)
		if bytes.Equal(actual, tt.expected) {
			continue
		}
		t.Errorf("NewRequestWrite(%v) => %v, actual %v", tt.memory, tt.expected, actual)
	}
}
