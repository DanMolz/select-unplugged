package sp

import (
	"testing"
)

func TestCRC(t *testing.T) {
	var cases = []struct {
		data     string
		expected uint16
	}{
		{"Q\x0a\x12\xa0\x00\x00", 0xfde2},
		{"Q\x01\xff\xaf\x00\x00", 0xcfcc},
	}

	for _, tt := range cases {
		actual := Crc([]byte(tt.data))
		if actual != tt.expected {
			t.Errorf("Crc(%q) => %v, actual %v", tt.data, tt.expected, actual)
		}
	}
}
