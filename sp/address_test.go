package sp

import (
	"bytes"
	"testing"
)

func TestAddressMessage(t *testing.T) {
	for _, tt := range []struct {
		address  Address
		expected Message
	}{
		{0x001f0000, Message("\x00\x00\x1f\x00")},
		{0x01234567, Message("\x67\x45\x23\x01")},
	} {
		actual := tt.address.Message()
		if bytes.Equal(actual, tt.expected) {
			continue
		}
		t.Errorf("Expected %v, actual %v", tt.expected, actual)
	}
}
