package sp

import (
	"testing"
)

func TestMessageString(t *testing.T) {
	for _, tt := range []struct {
		message  Message
		expected string
	}{
		{Message("\x00\x01\xff"), "0x0001ff"},
		{Message(""), "0x"},
	} {
		if tt.expected != tt.message.String() {
			t.Errorf("Expected %v, but got %v", tt.expected, tt.message.String())
		}
	}
}
