package sp

type SpUint16 = uint16

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

// ArrayTechnicalDataTab
var ArrayTechnicalDataTab Address = 40981

var VarCommonScaleForAcVolts = NewVariable(NewArea(ArrayTechnicalDataTab+19, 1))
var VarCommonScaleForAcCurrent = NewVariable(NewArea(ArrayTechnicalDataTab+20, 1))
var VarCommonScaleForDcVolts = NewVariable(NewArea(ArrayTechnicalDataTab+21, 1))
var VarCommonScaleForDcCurrent = NewVariable(NewArea(ArrayTechnicalDataTab+22, 1))
var VarCommonScaleForTemperature = NewVariable(NewArea(ArrayTechnicalDataTab+23, 1))
var VarCommonScaleForInternalVoltages = NewVariable(NewArea(ArrayTechnicalDataTab+24, 1))

var VarDCVolts = NewVariable(NewArea(ArrayTechnicalDataTab+25, 1)) //shortValue

var VarDCBatteryPower1 = NewVariable(NewArea(ArrayTechnicalDataTab+26, 1)) //shortValue
var VarDCBatteryPower2 = NewVariable(NewArea(ArrayTechnicalDataTab+27, 1)) //shortValue

var VarVersionNumber = NewVariable(NewArea(ArrayTechnicalDataTab+29, 1)) //shortValue
var VarGridSoftwareVersion = NewVariable(NewArea(ArrayTechnicalDataTab+33, 1)) //shortValue

var VarInverterRunHrsTotalAcc1 = NewVariable(NewArea(ArrayTechnicalDataTab+42, 1)) //shortValue
var VarInverterRunHrsTotalAcc2 = NewVariable(NewArea(ArrayTechnicalDataTab+43, 1)) //shortValue

var VarChargeStatus = NewVariable(NewArea(ArrayTechnicalDataTab+48, 2)) //intValue

var VarDaysSinceEqualise = NewVariable(NewArea(ArrayTechnicalDataTab+50, 1)) //shortValue

var VarUnitSerialNumber1 = NewVariable(NewArea(ArrayTechnicalDataTab+77, 1)) //shortValue
var VarUnitSerialNumber2 = NewVariable(NewArea(ArrayTechnicalDataTab+78, 1)) //shortValue

var VarNowTotalAcPowerSolar = NewVariable(NewArea(ArrayTechnicalDataTab+134, 2)) //intValue

var VarNowAcPowerSolar1 = NewVariable(NewArea(ArrayTechnicalDataTab+135, 2)) //intValue
var VarNowAcPowerSolar2 = NewVariable(NewArea(ArrayTechnicalDataTab+136, 2)) //intValue
var VarNowAcPowerSolar3 = NewVariable(NewArea(ArrayTechnicalDataTab+137, 2)) //intValue
var VarNowAcPowerSolar4 = NewVariable(NewArea(ArrayTechnicalDataTab+138, 2)) //intValue
var VarNowAcPowerSolar5 = NewVariable(NewArea(ArrayTechnicalDataTab+139, 2)) //intValue

var VarBattSocPercent = NewVariable(NewArea(ArrayTechnicalDataTab+108, 2)) //intValue

var VarDCCurrent = NewVariable(NewArea(ArrayTechnicalDataTab+110, 1)) //shortValue

var VarDCBatteryCurrent1 = NewVariable(NewArea(ArrayTechnicalDataTab+111, 1)) //shortValue
var VarDCBatteryCurrent2 = NewVariable(NewArea(ArrayTechnicalDataTab+112, 1)) //shortValue

var VarBatteryVolts = NewVariable(NewArea(ArrayTechnicalDataTab+71, 1)) //shortValue

var VarACLoadVoltage = NewVariable(NewArea(ArrayTechnicalDataTab+123, 2)) //intValue

var VarLoadAcPower1 = NewVariable(NewArea(ArrayTechnicalDataTab+126, 1)) //shortValue
var VarLoadAcPower2 = NewVariable(NewArea(ArrayTechnicalDataTab+127, 1)) //shortValue

var VarACInverterRmsAmps = NewVariable(NewArea(ArrayTechnicalDataTab+124, 1)) //shortValue
var VarAcCurrent = NewVariable(NewArea(ArrayTechnicalDataTab+119, 1)) //shortValue
var VarACGeneratorPower = NewVariable(NewArea(ArrayTechnicalDataTab+117, 1)) //shortValue
var VarACGeneratorPower5minAvg = NewVariable(NewArea(ArrayTechnicalDataTab+118, 1)) //shortValue
var VarACGeneratorRmsVolts = NewVariable(NewArea(ArrayTechnicalDataTab+128, 1)) //shortValue
var VarACGeneratorRmsAmps = NewVariable(NewArea(ArrayTechnicalDataTab+119, 1)) //shortValue

var VarNowPercentSolar = NewVariable(NewArea(ArrayTechnicalDataTab+140, 2)) //intValue

var VarShunt1Power = NewVariable(NewArea(ArrayTechnicalDataTab+115, 1)) //shortValue
var VarShunt2Power = NewVariable(NewArea(ArrayTechnicalDataTab+116, 1)) //shortValue

var VarTodayTotalAcEnergySolar = NewVariable(NewArea(ArrayTechnicalDataTab+176, 2)) //intValue

var VarEnergyGenerated = NewVariable(NewArea(ArrayTechnicalDataTab+176, 2)) //intValue

var VarTodayAcEnergySolar1 = NewVariable(NewArea(ArrayTechnicalDataTab+178, 2)) //intValue
var VarTodayAcEnergySolar2 = NewVariable(NewArea(ArrayTechnicalDataTab+179, 2)) //intValue
var VarTodayAcEnergySolar3 = NewVariable(NewArea(ArrayTechnicalDataTab+180, 2)) //intValue
var VarTodayAcEnergySolar4 = NewVariable(NewArea(ArrayTechnicalDataTab+181, 2)) //intValue
var VarTodayAcEnergySolar5 = NewVariable(NewArea(ArrayTechnicalDataTab+182, 2)) //intValue

var VarShunt1kWhAcc = NewVariable(NewArea(ArrayTechnicalDataTab+165, 1)) //shortValue
var VarShunt2kWhAcc = NewVariable(NewArea(ArrayTechnicalDataTab+166, 1)) //shortValue

var VarEnergyUsed = NewVariable(NewArea(ArrayTechnicalDataTab+169, 2)) //intValue

var VarBattInkWhAcc1 = NewVariable(NewArea(ArrayTechnicalDataTab+160, 1)) //shortValue
var VarBattInkWhAcc2 = NewVariable(NewArea(ArrayTechnicalDataTab+161, 1)) //shortValue

var VarBattOutkWhAcc1 = NewVariable(NewArea(ArrayTechnicalDataTab+162, 1)) //shortValue
var VarBattOutkWhAcc2 = NewVariable(NewArea(ArrayTechnicalDataTab+163, 1)) //shortValue

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
