package sp

type SpUint16 = uint16

// uint, ac_w, 41107
// var LoadAcPower variable = variableUint32{variable: {Address: AreaAddress{41107}}}
var LoadAcPower = NewVariableFloat64(41107, 2, ConvertUnsignedAcW)

// ushort, percent, 41089
// var BatterySoc variable = new(41089)
var BatterySoc = NewVariableFloat64(41089, 1, ConvertRatio)

// short, Shunt1Name, 49417
var Shunt1Name = newVariableString(49417, 1, ConvertShunt)
