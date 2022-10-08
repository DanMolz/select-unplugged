package sp

/*
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
}*/

func shortSliceToArray(data []byte) [2]byte {
	var ret [2]byte
	copy(ret[:], data[:4])
	return ret
}

/*
func TestInt16DcWConversion(t *testing.T) {
	for _, tt := range []struct {
		data     []byte
		expected float64
	}{
		{[]byte("\x46\xfe"), -1699.5849609375},
		{[]byte("\x3f\xfe"), -1726.50146484375},
		{[]byte("\x7a\xfe"), -1499.6337890625},
		{[]byte("\xfd\xff"), -11.53564453125},
	} {
		var variable = VariableFloat64{
			memory:    Memory{area: Area{words: 1}},
			converter: ConvertDcW,
		}
		variable.memory.SetData(tt.data)
		actual := variable.Convert()
		if actual != tt.expected {
			t.Errorf("%v.Convert() => %v, actual %v", variable, tt.expected, actual)
		}
	}
}*/

/*
int32 -> dc_w
('DCBatteryPower', b'\xd4\xff\xff\xff', -169.189453125),
('DCBatteryPower', b'\xf4\xff\xff\xff', -46.142578125),

ushort -> noop
('CommonScaleForDcVolts', b'\x1a\x04', 1050),
('CommonScaleForTemperature', b'\x12\x02', 530),

ushort -> dc_v
('BatteryVolts', b'\x41\x44', 55.989532470703125),
('BatteryVolts', b'\x3a\x44', 55.96710205078125),

ushort -> DegC
('BatteryTemperature', b'\x44\x05', 21.802978515625),

uint -> ac_w
('LoadAcPower', b'\xff\x02\x00\x00', 341.1567687988281),
('LoadAcPower', b'\x1d\x03\x00\x00', 354.5005798339844),

uint -> ac_wh
('ACLoadkWhTotalAcc', b'\x19\xeb\x00\x00', 5139822.509765625),

uint -> dc_wh
('BattOutkWhPreviousAcc', b'\x29\x00\x00\x00', 3783.69140625),

ushort -> ratio
('BattSocPercent', b'\xdb\x63', 99.85546875),

(u?)short -> shunt name
('Shunt1Name', b'\x01\x00', 'Solar'),
('Shunt2Name', b'\x00\x00', 'None'),
('Shunt1Name', b'\x3a\x00', 'Error'),
*/
