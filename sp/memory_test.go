package sp

import (
	"bytes"
	"testing"
)

func TestMemorySetDataRejectsIncorrectSize(t *testing.T) {
	for _, tt := range []struct {
		memory   Memory
		data     []byte
		expected string
	}{
		{Memory{area: Area{words: 1}}, []byte(""), "Got 0 bytes, expecting 2"},
		{Memory{area: Area{words: 1}}, []byte("a"), "Got 1 bytes, expecting 2"},
		{Memory{area: Area{words: 1}}, []byte("aaa"), "Got 3 bytes, expecting 2"},
		{Memory{area: Area{words: 2}}, []byte("aa"), "Got 2 bytes, expecting 4"},
	} {
		actual := tt.memory.SetData(tt.data)
		if actual.Error() == tt.expected {
			continue
		}
		t.Errorf("%v.SetData(%v) => %v, actual %v", tt.memory, tt.data, tt.expected, actual)
	}
}

func TestMemorySetDataGetData(t *testing.T) {
	for _, tt := range []struct {
		memory Memory
		data   []byte
	}{
		{Memory{area: Area{words: 1}}, []byte("aa")},
		{Memory{area: Area{words: 1}}, []byte("bb")},
		{Memory{area: Area{words: 2}}, []byte("cccc")},
	} {
		error := tt.memory.SetData(tt.data)
		if error != nil {
			t.Errorf(error.Error())
		}
		actual := tt.memory.data
		if bytes.Equal(actual, tt.data) {
			continue
		}
		t.Errorf("expected %v, actual %v", tt.data, actual)
	}
}
