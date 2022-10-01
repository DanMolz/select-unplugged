package sp

import "testing"

var scales Scales = Scales{
	DcVolts:     1050,
	DcCurrent:   12000,
	Temperature: 530,
	AcVolts:     5300,
	AcCurrent:   2200,
}

func TestConvertFloats(t *testing.T) {
	assertion := func(t *testing.T, expected float64, actual float64) {
		if actual == expected {
			return
		}
		t.Fatalf("Expected converted value to be %v, but it was %v", expected, actual)
	}
	assertion(t, 354.5005798339844, ConvertUnsignedAcW(797, scales))
	assertion(t, 341.1567687988281, ConvertUnsignedAcW(767, scales))
	assertion(t, 55.96710205078125, ConvertDcV(17466, scales))
	assertion(t, 55.989532470703125, ConvertDcV(17473, scales))
	assertion(t, -1699.5849609375, ConvertDcW(-442, scales))
	assertion(t, -169.189453125, ConvertDcW(-44, scales))
	assertion(t, -46.142578125, ConvertDcW(-12, scales))
	assertion(t, -11.53564453125, ConvertDcW(-3, scales))
	assertion(t, 5139822.509765625, ConvertAcWh(60185, scales))
	assertion(t, 3783.69140625, ConvertDcWh(41, scales))
	assertion(t, 0.9985546875, ConvertRatio(25563, scales))
	assertion(t, 21.802978515625, ConvertTemperature(1348, scales))
}

func TestConvertShunts(t *testing.T) {
	assertion := func(t *testing.T, expected string, actual string) {
		if actual == expected {
			return
		}
		t.Fatalf("Expected converted value to be %v, but it was %v", expected, actual)
	}
	assertion(t, "None", ConvertShunt(0, scales))
	assertion(t, "Solar", ConvertShunt(1, scales))
	assertion(t, "Error", ConvertShunt(58, scales))
}
