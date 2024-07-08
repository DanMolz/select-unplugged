package sp

type SpUint16 = uint16

// TODO: construct variables in a function and inject them

// uint, ac_w, 41107
// var LoadAcPower variable = variableUint32{variable: {Address: AreaAddress{41107}}}
//var LoadAcPower = NewVariableFloat64(41107, 2, ConvertUnsignedAcW)

// ushort, percent, 41089
// var BatterySoc variable = new(41089)
//var BatterySoc = NewVariableFloat64(41089, 1, ConvertRatio)

// short, Shunt1Name, 49417
//var Shunt1Name = newVariableString(49417, 1, ConvertShunt)

// Read to commence login challenge, written with md5 of read challenge + password
var VarLoginHash = NewVariable(NewArea(2031616, 8))

// ushort, 1 if we're logged in, CRC fails if we're not
var VarLoginStatus = NewVariable(NewArea(2031632, 1))

// ushort, disconnect from comm port 1
var VarSpLinkDisconnectingComms1 = NewVariable(NewArea(40973, 1))

// uint, how much energy has gone in to the battery today
var VarBatteryEnergyInToday = NewVariable(NewArea(41135, 2))

// uint, how much energy has gone in to the battery total
var VarBatteryEnergyInTotal = NewVariable(NewArea(41354, 2))

// ushort, percent, battery state of charge
var VarBatterySoc = NewVariable(NewArea(41089, 1))

// ushort
var VarCommonScaleForAcVolts = NewVariable(NewArea(41000, 1))
var VarCommonScaleForAcCurrent = NewVariable(NewArea(41001, 1))
var VarCommonScaleForDcVolts = NewVariable(NewArea(41002, 1))
var VarCommonScaleForDcCurrent = NewVariable(NewArea(41003, 1))
var VarCommonScaleForTemperature = NewVariable(NewArea(41004, 1))
var VarCommonScaleForInternalVoltages = NewVariable(NewArea(41005, 1))

// TESTING
var VarACLoadkWhTotalAcc = NewVariable(NewArea(41519, 2))
var VarLoadAcPower = NewVariable(NewArea(41107, 2))

//uint
var VarInverterTime = NewVariable(NewArea(1900544, 7))

// getStatus
var LedStatus = NewVariable(NewArea(41175, 69))
var CommonParameters = NewVariable(NewArea(49152, 196))
var ServiceRequired = NewVariable(NewArea(41533, 75))
var TechnicalData = NewVariable(NewArea(41912, 50))
var SystemSchedulerConfigSettings = NewVariable(NewArea(51200, 184))
var TechnicalDataTab = NewVariable(NewArea(40981, 182))
