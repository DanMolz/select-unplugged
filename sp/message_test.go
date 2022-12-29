package sp

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestMessageDescribe(t *testing.T) {
	for _, tt := range []struct {
		message  Message
		expected string
	}{
		// TODO: represent invalid CRCs
		// TODO: represent write messages
		// TODO: represent messages with known variables
		// TODO: represent messages with multiple variables
		{
			Message("\x51\x00\x00\xa0\x00\x00\x9d\x4b\x01\x00\xd8\x19"),
			"Q@40960=0x0100",
		},
		{
			Message("\x51\x00\x00\xa0\x00\x00\x9d\x4b"),
			"Q@40960",
		},
	} {
		assert.Equal(t, tt.expected, tt.message.Describe())
	}

}
