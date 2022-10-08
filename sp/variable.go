package sp

type Variable interface {
}

type ConverterFloat64 func(int, Scales) float64
type ConverterString func(int, Scales) string

type VariableFloat64 struct {
	memory    Memory
	converter ConverterFloat64
}

func NewVariableFloat64(
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

// TODO: make sure we have the following working:
//    uint           -> string
//    ushort         -> string
//    unsigned int   -> float
//    signed int     -> float
//    unsigned short -> float
//    signed short   -> float
