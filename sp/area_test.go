package sp

import (
	"bytes"
	"testing"
)

func TestAreaString(t *testing.T) {
	var cases = []struct {
		address  AreaAddress
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

func TestAreaMessage(t *testing.T) {
	var cases = []struct {
		address  AreaAddress
		words    Words
		expected Message
	}{
		{0, 0, []byte("\x00\x00\x00\x00\x00")},
		{0x12345678, 59, []byte("\x3b\x78\x56\x34\x12")},
		{0x0000a000, 0, Message("\x00\x00\xa0\x00\x00")},
		{0x0000a093, 3, Message("\x03\x93\xa0\x00\x00")},
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
