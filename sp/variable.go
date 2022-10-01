package sp

type Memory struct {
	area Area
	data []byte
}

type ConverterFloat64 func(int, Scales) float64
type ConverterString func(int, Scales) string

type VariableFloat64 struct {
	memory    Memory
	converter ConverterFloat64
}

func newVariableFloat64(
	address AreaAddress,
	words AreaWords,
	converter ConverterFloat64,
) VariableFloat64 {
	return VariableFloat64{
		memory: Memory{
			area: Area{
				address: address,
				words:   words,
			},
		},
		converter: converter,
	}
}

type VariableString struct {
	memory    Memory
	converter ConverterString
}

func newVariableString(
	address AreaAddress,
	words AreaWords,
	converter ConverterString,
) VariableString {
	return VariableString{
		memory: Memory{
			area: Area{
				address: address,
				words:   words,
			},
		},
		converter: converter,
	}
}

// uint, ac_w, 41107
// var LoadAcPower variable = variableUint32{variable: {Address: AreaAddress{41107}}}
var LoadAcPower = newVariableFloat64(41107, 2, ConvertUnsignedAcW)

// ushort, percent, 41089
// var BatterySoc variable = new(41089)
var BatterySoc = newVariableFloat64(41089, 1, ConvertRatio)

// short, Shunt1Name, 49417
var Shunt1Name = newVariableString(49417, 1, ConvertShunt)

// TODO: make sure we have the following working:
//    uint           -> string
//    ushort         -> string
//    unsigned int   -> float
//    signed int     -> float
//    unsigned short -> float
//    signed short   -> float
