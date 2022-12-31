package sp

type Variable struct {
	memory Memory
}

func NewVariable(address Address, words Words) Variable {
	return Variable{
		memory: NewMemory(address, words),
	}
}

func (v Variable) Address() Address {
	return v.memory.Address()
}

func (v Variable) Words() Words {
	return v.memory.Words()
}

func (v Variable) Memory() Memory {
	return v.memory
}

func (v Variable) Area() Area {
	return NewArea(v.Address(), v.Words())
}

type ConverterFloat64 func(int, Scales) float64
type ConverterString func(int, Scales) string

type VariableFloat64 struct {
	memory    Memory
	converter ConverterFloat64
}

/*
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
}*/

type VariableString struct {
	memory    Memory
	converter ConverterString
}

/*
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
}*/

// TODO: make sure we have the following working:
//    uint           -> string
//    ushort         -> string
//    unsigned int   -> float
//    signed int     -> float
//    unsigned short -> float
//    signed short   -> float
