package sp

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMemory(t *testing.T) {
	for _, tt := range []struct {
		words Words
	}{
		{1},
		{2},
		{512},
	} {
		assert.PanicsWithError(t, "Read from uninitialized Memory", func() {
			memory := NewMemory(NewArea(0, tt.words))
			memory.Data()
		})
	}
}

func TestMemorySetDataRejectsIncorrectSize(t *testing.T) {
	for _, tt := range []struct {
		memory   Memory
		data     Data
		expected string
	}{
		{Memory{data: Data("aa")}, Data(""), "Got 0 bytes, expecting 2"},
		{Memory{data: Data("aa")}, Data("a"), "Got 1 bytes, expecting 2"},
		{Memory{data: Data("aa")}, Data("aaa"), "Got 3 bytes, expecting 2"},
		{Memory{data: Data("aaaa")}, Data("aa"), "Got 2 bytes, expecting 4"},
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
		data   Data
	}{
		{Memory{data: Data("aa")}, Data("aa")},
		{Memory{data: Data("aa")}, Data("bb")},
		{Memory{data: Data("aaaa")}, Data("cccc")},
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
