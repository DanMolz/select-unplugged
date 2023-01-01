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
		{Memory{area: Area{words: 1}}, Data(""), "Got 0 bytes, expecting 2"},
		{Memory{area: Area{words: 1}}, Data("a"), "Got 1 bytes, expecting 2"},
		{Memory{area: Area{words: 1}}, Data("aaa"), "Got 3 bytes, expecting 2"},
		{Memory{area: Area{words: 2}}, Data("aa"), "Got 2 bytes, expecting 4"},
	} {
		assert.PanicsWithError(t, tt.expected, func() {
			tt.memory.SetData(tt.data)
		})
	}
}

func TestMemorySetDataGetData(t *testing.T) {
	for _, tt := range []struct {
		memory Memory
		data   Data
	}{
		{Memory{area: Area{words: 1}}, Data("aa")},
		{Memory{area: Area{words: 1}}, Data("bb")},
		{Memory{area: Area{words: 2}}, Data("cccc")},
	} {
		tt.memory.SetData(tt.data)
		actual := tt.memory.data
		if bytes.Equal(actual, tt.data) {
			continue
		}
		t.Errorf("expected %v, actual %v", tt.data, actual)
	}
}
