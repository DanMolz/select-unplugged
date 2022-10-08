package sp

import (
	"testing"
)

func TestResponseExpectedLength(t *testing.T) {
	for _, tt := range []struct {
		request  Request
		expected int
	}{
		{Request(NewRequestQuery(Area{0xa000, 0})), 12},
		{Request(NewRequestQuery(Area{0xa000, 0xff})), 522},
		{
			Request(NewRequestWrite(Memory{
				address: 0x01f,
				data:    []byte("0000000000000000"),
			})),
			26,
		},
	} {
		actual := NewResponse(tt.request).ExpectedLength()
		if actual == tt.expected {
			continue
		}
		t.Errorf("Expected %v, actual %v", tt.expected, actual)
	}
}
