package sp

import (
	"testing"
)

func TestVariableWords(t *testing.T) {
	for _, tt := range []struct {
		variable Variable
		expected AreaWords
	}{
		{LoadAcPower, 2},
		{BatterySoc, 1},
		{Shunt1Name, 1},
	} {
		actual := tt.variable.Memory().area.words
		if actual != tt.expected {
			t.Errorf("%v.memory.area.words => %v, actual %v", tt.variable, tt.expected, actual)
		}
	}
}
