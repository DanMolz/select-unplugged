package sp

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaString(t *testing.T) {
	var cases = []struct {
		address  Address
		words    Words
		expected string
	}{
		{0, 0, "Area(0x00000000, 0)"},
		{0x12345678, 59, "Area(0x12345678, 59)"},
	}
	for _, tt := range cases {
		actual := Area{tt.address, tt.words}.String()
		if actual != tt.expected {
			t.Errorf("Area(%v, %v).String() => %v, actual %v", tt.address, tt.words, tt.expected, actual)
		}
	}
}

func TestNewAreaErrors(t *testing.T) {
	assert.PanicsWithError(t, "Non-zero word length required", func() {
		NewArea(0, 0)
	})
}

func TestAreaMessage(t *testing.T) {
	var cases = []struct {
		address  Address
		words    Words
		expected Message
	}{
		{0, 1, []byte("\x00\x00\x00\x00\x00")},
		{0x12345678, 60, []byte("\x3b\x78\x56\x34\x12")},
		{0x0000a000, 1, Message("\x00\x00\xa0\x00\x00")},
		{0x0000a093, 4, Message("\x03\x93\xa0\x00\x00")},
	}

	for _, tt := range cases {
		actual := Area{tt.address, tt.words}.Message()
		if bytes.Equal(actual, tt.expected) {
			continue
		}
		t.Errorf(
			"Area(%v, %v).Message() => %v, actual %v",
			tt.address,
			tt.words,
			tt.expected,
			actual,
		)
	}
}
