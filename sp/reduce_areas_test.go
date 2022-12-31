package sp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReduceAreas(t *testing.T) {
	for _, tt := range []struct {
		areas    []Area
		expected []Area
	}{
		{
			[]Area{NewArea(0, 1)},
			[]Area{NewArea(0, 1)},
		},
		{
			[]Area{NewArea(0x0010, 1), NewArea(0x0014, 2)},
			[]Area{NewArea(0x0010, 6)},
		},
		{
			[]Area{NewArea(0x0004, 2), NewArea(0x0000, 1)},
			[]Area{NewArea(0x0000, 6)},
		},
		{
			[]Area{NewArea(0x0001, 2), NewArea(0x0002, 2), NewArea(0x0003, 2)},
			[]Area{NewArea(0x0001, 4)},
		},
		// split across max words (256)
		{
			[]Area{
				NewArea(0x0000, 2),
				NewArea(0x00ff, 1),
				NewArea(0x0103, 1),
				NewArea(0x0203, 1),
				NewArea(0x0301, 2),
			},
			[]Area{
				NewArea(0x0000, 256), NewArea(0x0103, 1), NewArea(0x0203, 256)},
		},
	} {
		assert.Equal(t, tt.expected, ReduceAreas(tt.areas))
	}
}
