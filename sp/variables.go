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


// 40981, 182, parSelectronicConnectionInfo, "TechnicalDataTab"
var VarDCVolts = NewVariable(NewArea(41006, 1))
var VarDCBatteryPower1 = NewVariable(NewArea(41007, 1))
var VarDCBatteryPower2 = NewVariable(NewArea(41008, 1))
var VarVersionNumber = NewVariable(NewArea(41010, 1))
var VarGridSoftwareVersion = NewVariable(NewArea(41014, 1))
// var VarBuildDate = NewVariable(NewArea(41011, 2))

var VarChargeStatus = NewVariable(NewArea(41029, 2))

var VarInverterRunHrsTotalAcc1 = NewVariable(NewArea(41023, 1))
var VarInverterRunHrsTotalAcc2 = NewVariable(NewArea(41024, 1))


/*
TYPES = {
    "ushort": {
        FORMAT: "<H",
        WORDS: 1,
    },
    "short": {
        FORMAT: "<h",
        WORDS: 1,
    },
    "uint": {
        FORMAT: "<I",
        WORDS: 2,
    },
    "int": {
        FORMAT: "<i",
        WORDS: 2,
    },
}
*/