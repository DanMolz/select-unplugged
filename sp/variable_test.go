package sp

import (
	"testing"
)

func TestVariableWords(t *testing.T) {
	for _, tt := range []struct {
		variable VariableFloat64
		expected AreaWords
	}{
		{LoadAcPower, 2},
		{BatterySoc, 1},
		// TODO: figure out how to mix this type - {Shunt1Name, 1},
	} {
		actual := tt.variable.memory.area.words
		if actual != tt.expected {
			t.Errorf("%v.memory.area.words => %v, actual %v", tt.variable, tt.expected, actual)
		}
	}
}
