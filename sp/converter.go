package sp

// Magic constants used for shifting integers to floating point numbers
const magic = 32768.0
const magicAcWUnsignedDivisor = magic * 800.0
const magicAcWSignedDivisor = magic * 100.0
const magicDcWDivisor = magic * 100.0
const magicDcVDivisor = magic * 10.0
const magicWhMultiplier = 24.0
const magicWhDivisor = magic * 100.0
const magicTemperatureDivisor = magic
const magicRatioDivisor = 25600.0

func ConvertUnsignedAcW(raw int, scales Scales) float64 {
	return float64(raw) * scales.AcVolts * scales.AcCurrent / magicAcWUnsignedDivisor
}

func ConvertAcWh(raw int, scales Scales) float64 {
	return float64(raw) * magicWhMultiplier * scales.AcVolts * scales.AcCurrent / magicWhDivisor
}

func ConvertDcV(raw int, scales Scales) float64 {
	return float64(raw) * scales.DcVolts / magicDcVDivisor
}

func ConvertDcW(raw int, scales Scales) float64 {
	return float64(raw) * scales.DcVolts * scales.DcCurrent / magicDcWDivisor
}

func ConvertTemperature(raw int, scales Scales) float64 {
	return float64(raw) * scales.Temperature / magicTemperatureDivisor
}

func ConvertRatio(raw int, scales Scales) float64 {
	return float64(raw) / magicRatioDivisor
}

func ConvertDcWh(raw int, scales Scales) float64 {
	return float64(raw) * magicWhMultiplier * scales.DcVolts * scales.DcCurrent / magicWhDivisor
}

func ConvertShunt(raw int, scales Scales) string {
	switch raw {
	case 0:
		return "None"
	case 1:
		return "Solar"
	case 2:
		return "Wind"
	case 3:
		return "Hydro"
	case 4:
		return "Charger"
	case 5:
		return "Load"
	case 6:
		return "Dual"
	case 7:
		return "Multiple SP PROs"
	case 8:
		return "Log Only"
	case 9:
		return "System SoC"
	case 10:
		return "Direct SoC Input"
	}
	return "Error"
}
