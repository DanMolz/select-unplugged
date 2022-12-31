package sp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractMemory(t *testing.T) {
	for _, tt := range []struct {
		memories []Memory
		area     Area
		expected Memory
	}{
		{
			[]Memory{
				Memory{NewArea(10, 10), Data("--somedata----")},
			},
			NewArea(11, 4),
			Memory{NewArea(11, 4), Data("somedata")},
		},
		{
			[]Memory{
				Memory{NewArea(10, 10), Data("somedata------")},
			},
			NewArea(10, 4),
			Memory{NewArea(10, 4), Data("somedata")},
		},
		{
			[]Memory{
				Memory{NewArea(10, 10), Data("----somedata--")},
			},
			NewArea(12, 4),
			Memory{NewArea(12, 4), Data("somedata")},
		},
		{
			[]Memory{
				Memory{NewArea(0x0000, 10), Data("aaaaaaaaaaaaaaaaaaaa")},
				Memory{NewArea(0x1230, 10), Data("mmmmmmmm78901234mmmm")},
				Memory{NewArea(0x9999, 10), Data("zzzzzzzzzzzzzzzzzzzz")},
			},
			NewArea(0x1234, 4),
			Memory{NewArea(0x1234, 4), Data("78901234")},
		},
	} {
		assert.Equal(t, tt.expected, ExtractMemory(tt.area, tt.memories))
	}
}
