package sp

import (
	"encoding/binary"
	"encoding/json"
	"math"
	"os"

	log "github.com/sirupsen/logrus"
)

type StatsPayload struct {
	DcVolts                 float64 `json:"dcVolts"`
	DcBatteryPower          float64 `json:"dcBatteryPower"`
	InverterRunHrsTotalAcc  float64 `json:"inverterRunHrsTotalAcc"`
	ChargeStatus            float64 `json:"chargeStatus"`
	DaysSinceEqualise       float64 `json:"daysSinceEqualise"`
	NowTotalAcPowerSolar    float64 `json:"nowTotalAcPowerSolar"`
	NowAcPowerSolar1        float64 `json:"nowAcPowerSolar1"`
	NowAcPowerSolar2        float64 `json:"nowAcPowerSolar2"`
	NowAcPowerSolar3        float64 `json:"nowAcPowerSolar3"`
	NowAcPowerSolar4        float64 `json:"nowAcPowerSolar4"`
	NowAcPowerSolar5        float64 `json:"nowAcPowerSolar5"`
	BattSocPercent          float64 `json:"battSocPercent"`
	DcCurrent               float64 `json:"dcCurrent"`
	DcBatteryCurrent        float64 `json:"dcBatteryCurrent"`
	BatteryVolts            float64 `json:"batteryVolts"`
	AcLoadVoltage           float64 `json:"acLoadVoltage"`
	LoadAcPower             float64 `json:"loadAcPower"`
	PowerUsed               float64 `json:"powerUsed"`
	AcInverterRmsAmps       float64 `json:"acInverterRmsAmps"`
	AcCurrent               float64 `json:"acCurrent"`
	AcGeneratorPower        float64 `json:"acGeneratorPower"`
	AcGeneratorPower5minAvg float64 `json:"acGeneratorPower5minAvg"`
	AcGeneratorRmsVolts     float64 `json:"acGeneratorRmsVolts"`
	AcGeneratorRmsAmps      float64 `json:"acGeneratorRmsAmps"`
	NowPercentSolar         float64 `json:"nowPercentSolar"`
	TodayTotalAcEnergySolar float64 `json:"todayTotalAcEnergySolar"`
	EnergyGenerated         float64 `json:"energyGenerated"`
	TodayAcEnergySolar1     float64 `json:"todayAcEnergySolar1"`
	TodayAcEnergySolar2     float64 `json:"todayAcEnergySolar2"`
	TodayAcEnergySolar3     float64 `json:"todayAcEnergySolar3"`
	TodayAcEnergySolar4     float64 `json:"todayAcEnergySolar4"`
	TodayAcEnergySolar5     float64 `json:"todayAcEnergySolar5"`
	AcLoadkWhAcc            float64 `json:"acLoadkWhAcc"`
	EnergyUsed              float64 `json:"energyUsed"`
	BattInkWhAcc            float64 `json:"battInkWhAcc"`
	BattOutkWhAcc           float64 `json:"battOutkWhAcc"`
	BatteryEnergyInToday    float64 `json:"batteryEnergyInToday"`
}

func StatsSelectLiveRenderV2(protocol *Protocol) {
	variables := []*Variable{
		&VarCommonScaleForDcVolts,
		&VarCommonScaleForDcCurrent,
		&VarCommonScaleForAcVolts,
		&VarCommonScaleForAcCurrent,
		&VarCommonScaleForTemperature,
		&VarCommonScaleForInternalVoltages,
		&VarDCVolts,
		&VarDCBatteryPower1,
		&VarDCBatteryPower2,
		&VarInverterRunHrsTotalAcc1,
		&VarInverterRunHrsTotalAcc2,
		&VarNowTotalAcPowerSolar,
		&VarNowAcPowerSolar1,
		&VarNowAcPowerSolar2,
		&VarNowAcPowerSolar3,
		&VarNowAcPowerSolar4,
		&VarNowAcPowerSolar5,
		&VarChargeStatus,
		&VarDaysSinceEqualise,
		&VarBattSocPercent,
		&VarDCCurrent,
		&VarDCBatteryCurrent1,
		&VarDCBatteryCurrent2,
		&VarBatteryVolts,
		&VarACLoadVoltage,
		&VarLoadAcPower1,
		&VarLoadAcPower2,
		&VarACInverterRmsAmps,
		&VarAcCurrent,
		&VarACGeneratorPower,
		&VarACGeneratorPower5minAvg,
		&VarACGeneratorRmsVolts,
		&VarACGeneratorRmsAmps,
		&VarNowPercentSolar,
		&VarTodayTotalAcEnergySolar,
		&VarEnergyGenerated,
		&VarTodayAcEnergySolar1,
		&VarTodayAcEnergySolar2,
		&VarTodayAcEnergySolar3,
		&VarTodayAcEnergySolar4,
		&VarTodayAcEnergySolar5,
		&VarEnergyUsed,
		&VarBattInkWhAcc1,
		&VarBattInkWhAcc2,
		&VarBattOutkWhAcc1,
		&VarBattOutkWhAcc2,
		&VarBatteryEnergyInToday,
	}

	protocol.Query(variables)

	CommonScaleForAcVolts := float64(binary.LittleEndian.Uint16(VarCommonScaleForAcVolts.memory.Data()))
	CommonScaleForAcCurrent := float64(binary.LittleEndian.Uint16(VarCommonScaleForAcCurrent.memory.Data()))
	CommonScaleForDcVolts := float64(binary.LittleEndian.Uint16(VarCommonScaleForDcVolts.memory.Data()))
	CommonScaleForDcCurrent := float64(binary.LittleEndian.Uint16(VarCommonScaleForDcCurrent.memory.Data()))
	// CommonScaleForTemperature := float64(binary.LittleEndian.Uint16(VarCommonScaleForTemperature.memory.Data()))
	// CommonScaleForInternalVoltages := float64(binary.LittleEndian.Uint16(VarCommonScaleForInternalVoltages.memory.Data()))

	// (arrayOfUInt16[25].shortValue() * this.j.CommonScaleForDcVolts / 327680.0F)
	dcVolts := float64(binary.LittleEndian.Uint16(VarDCVolts.memory.Data())) * CommonScaleForDcVolts / MAGIC_DC_V_DIVISOR
	log.Debugf("dcVolts: %f", dcVolts)

	// (this.g.convert2UShortsInto1Uint(arrayOfUInt16[26], arrayOfUInt16[27]) * -1.0F * this.j.CommonScaleForDcVolts * this.j.CommonScaleForDcCurrent / 3276800.0F)
	dcBatteryPower := float64(convert2UShortsInto1Uint(binary.LittleEndian.Uint16(VarDCBatteryPower1.memory.Data()), binary.LittleEndian.Uint16(VarDCBatteryPower2.memory.Data()))) * -1.0 * CommonScaleForDcVolts * CommonScaleForDcCurrent / MAGIC_DC_W_DIVISOR
	log.Debugf("dcBatteryPower: %f", dcBatteryPower)

	// (this.g.convert2UShortsInto1Uint(arrayOfUInt16[42], arrayOfUInt16[43]) / 60.0F)
	inverterRunHrsTotalAcc := float64(convert2UShortsInto1Uint(binary.LittleEndian.Uint16(VarInverterRunHrsTotalAcc1.memory.Data()), binary.LittleEndian.Uint16(VarInverterRunHrsTotalAcc2.memory.Data()))) / 60
	log.Debugf("inverterRunHrsTotalAcc: %f", inverterRunHrsTotalAcc)

	// arrayOfUInt16[48].intValue()
	chargeStatus := float64(binary.LittleEndian.Uint16(VarChargeStatus.memory.Data()))
	log.Debugf("chargeStatus: %f", chargeStatus)

	// arrayOfUInt16[50].shortValue()
	daysSinceEqualise := float64(binary.LittleEndian.Uint16(VarDaysSinceEqualise.memory.Data()))
	log.Debugf("daysSinceEqualise: %f", daysSinceEqualise)

	// arrayOfUInt16[134].intValue() * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 3276800.0F
	nowTotalAcPowerSolar := float64(binary.LittleEndian.Uint16(VarNowTotalAcPowerSolar.memory.Data())) * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	log.Debugf("powerGenerated: %f", nowTotalAcPowerSolar)

	// (arrayOfUInt16[139].intValue() * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 3276800.0F)
	nowAcPowerSolar1 := float64(binary.LittleEndian.Uint16(VarNowAcPowerSolar1.memory.Data())) * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	nowAcPowerSolar2 := float64(binary.LittleEndian.Uint16(VarNowAcPowerSolar2.memory.Data())) * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	nowAcPowerSolar3 := float64(binary.LittleEndian.Uint16(VarNowAcPowerSolar3.memory.Data())) * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	nowAcPowerSolar4 := float64(binary.LittleEndian.Uint16(VarNowAcPowerSolar4.memory.Data())) * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	nowAcPowerSolar5 := float64(binary.LittleEndian.Uint16(VarNowAcPowerSolar5.memory.Data())) * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	log.Debugf("nowAcPowerSolar1: %f", nowAcPowerSolar1)
	log.Debugf("nowAcPowerSolar2: %f", nowAcPowerSolar2)
	log.Debugf("nowAcPowerSolar3: %f", nowAcPowerSolar3)
	log.Debugf("nowAcPowerSolar4: %f", nowAcPowerSolar4)
	log.Debugf("nowAcPowerSolar5: %f", nowAcPowerSolar5)

	// arrayOfUInt16[108].intValue() / 256.0F
	battSocPercent := float64(binary.LittleEndian.Uint16(VarBattSocPercent.memory.Data())) / 256
	log.Debugf("battSocPercent: %f", battSocPercent)

	// (arrayOfUInt16[110].shortValue() * this.j.CommonScaleForDcCurrent / 327680.0F)
	dcCurrent := float64(binary.LittleEndian.Uint16(VarDCCurrent.memory.Data())) * CommonScaleForDcCurrent / MAGIC_DC_V_DIVISOR
	log.Debugf("dcCurrent: %f", dcCurrent)

	// (this.g.convert2UShortsInto1Uint(arrayOfUInt16[111], arrayOfUInt16[112]) * this.j.CommonScaleForDcCurrent / 327680.0F)
	dcBatteryCurrent := float64(convert2UShortsInto1Uint(binary.LittleEndian.Uint16(VarDCBatteryCurrent1.memory.Data()), binary.LittleEndian.Uint16(VarDCBatteryCurrent2.memory.Data()))) * CommonScaleForDcCurrent / MAGIC_DC_V_DIVISOR
	log.Debugf("dcBatteryCurrent: %f", dcBatteryCurrent)

	// arrayOfUInt16[71].shortValue() * this.j.CommonScaleForDcVolts / 327680.0F
	batteryVolts := float64(binary.LittleEndian.Uint16(VarBatteryVolts.memory.Data())) * CommonScaleForDcVolts / MAGIC_DC_V_DIVISOR
	log.Debugf("batteryVolts: %f", batteryVolts)

	// (arrayOfUInt16[123].intValue() * this.j.CommonScaleForAcVolts / 327680.0F)
	acLoadVoltage := float64(binary.LittleEndian.Uint16(VarACLoadVoltage.memory.Data())) * CommonScaleForAcVolts / MAGIC_DC_V_DIVISOR
	log.Debugf("acLoadVoltage: %f", acLoadVoltage)

	// this.g.convert2UShortsInto1Uint(arrayOfUInt16[126], arrayOfUInt16[127]) * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 2.62144E7F
	loadAcPower := float64(convert2UShortsInto1Uint(binary.LittleEndian.Uint16(VarLoadAcPower1.memory.Data()), binary.LittleEndian.Uint16(VarLoadAcPower2.memory.Data()))) * CommonScaleForAcVolts * CommonScaleForAcCurrent / 2.62144e7
	log.Debugf("loadAcPower: %f", loadAcPower)

	// this.g.convert2UShortsInto1Uint(arrayOfUInt16[126], arrayOfUInt16[127]) * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 2.62144E7F
	powerUsed := float64(convert2UShortsInto1Uint(binary.LittleEndian.Uint16(VarLoadAcPower1.memory.Data()), binary.LittleEndian.Uint16(VarLoadAcPower2.memory.Data()))) * CommonScaleForAcVolts * CommonScaleForAcCurrent / 2.62144e7
	log.Debugf("powerUsed: %f", powerUsed)

	// (arrayOfUInt16[124].shortValue() * this.j.CommonScaleForAcCurrent / 327680.0F)
	acInverterRmsAmps := float64(binary.LittleEndian.Uint16(VarACInverterRmsAmps.memory.Data())) * CommonScaleForAcCurrent / MAGIC_DC_V_DIVISOR
	log.Debugf("acInverterRmsAmps: %f", acInverterRmsAmps)

	// Math.abs(arrayOfUInt16[119].shortValue() * this.j.CommonScaleForAcCurrent / 327680.0F)
	acCurrent := math.Abs(float64(binary.LittleEndian.Uint16(VarAcCurrent.memory.Data())) * CommonScaleForAcCurrent / MAGIC_DC_V_DIVISOR)
	log.Debugf("acCurrent: %f", acCurrent)

	// (-1.0F * arrayOfUInt16[117].shortValue() * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 3276800.0F)
	acGeneratorPower := (-1.0 * float64(binary.LittleEndian.Uint16(VarACGeneratorPower.memory.Data())) * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR)
	log.Debugf("acGeneratorPower: %f", acGeneratorPower)

	// (-1.0F * arrayOfUInt16[118].shortValue() * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 3276800.0F)
	acGeneratorPower5minAvg := (-1.0 * float64(binary.LittleEndian.Uint16(VarACGeneratorPower5minAvg.memory.Data())) * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR)
	log.Debugf("acGeneratorPower5minAvg: %f", acGeneratorPower5minAvg)

	// (arrayOfUInt16[128].shortValue() * this.j.CommonScaleForAcVolts / 327680.0F)
	acGeneratorRmsVolts := float64(binary.LittleEndian.Uint16(VarACGeneratorRmsVolts.memory.Data())) * CommonScaleForAcVolts / MAGIC_DC_V_DIVISOR
	log.Debugf("acGeneratorRmsVolts: %f", acGeneratorRmsVolts)

	// Math.abs(arrayOfUInt16[119].shortValue() * this.j.CommonScaleForAcCurrent / 327680.0F)
	acGeneratorRmsAmps := math.Abs(float64(binary.LittleEndian.Uint16(VarACGeneratorRmsAmps.memory.Data())) * CommonScaleForAcCurrent / MAGIC_DC_V_DIVISOR)
	log.Debugf("acGeneratorRmsAmps: %f", acGeneratorRmsAmps)

	// arrayOfUInt16[140].intValue()
	nowPercentSolar := float64(binary.LittleEndian.Uint16(VarNowPercentSolar.memory.Data()))
	log.Debugf("nowPercentSolar: %f", nowPercentSolar)

	// arrayOfUInt16[176].intValue() * 24.0F * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 3276800.0F
	todayTotalAcEnergySolar := float64(binary.LittleEndian.Uint16(VarTodayTotalAcEnergySolar.memory.Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	log.Debugf("todayTotalAcEnergySolar: %f", todayTotalAcEnergySolar)

	// ((arrayOfUInt16[176].intValue() * 24.0F * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 3276800.0F)) / 1000.0F)
	energyGenerated := float64(float64(binary.LittleEndian.Uint16(VarEnergyGenerated.memory.Data()))*MAGIC_WH_MULTIPLIER*CommonScaleForAcVolts*CommonScaleForAcCurrent/MAGIC_DC_W_DIVISOR) / 1000
	log.Debugf("energyGenerated: %f", energyGenerated)

	// (arrayOfUInt16[178].intValue() * 24.0F * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 3276800.0F)
	todayAcEnergySolar1 := float64(binary.LittleEndian.Uint16(VarTodayAcEnergySolar1.memory.Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	todayAcEnergySolar2 := float64(binary.LittleEndian.Uint16(VarTodayAcEnergySolar2.memory.Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	todayAcEnergySolar3 := float64(binary.LittleEndian.Uint16(VarTodayAcEnergySolar3.memory.Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	todayAcEnergySolar4 := float64(binary.LittleEndian.Uint16(VarTodayAcEnergySolar4.memory.Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	todayAcEnergySolar5 := float64(binary.LittleEndian.Uint16(VarTodayAcEnergySolar5.memory.Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR
	log.Debugf("todayAcEnergySolar1: %f", todayAcEnergySolar1)
	log.Debugf("todayAcEnergySolar2: %f", todayAcEnergySolar2)
	log.Debugf("todayAcEnergySolar3: %f", todayAcEnergySolar3)
	log.Debugf("todayAcEnergySolar4: %f", todayAcEnergySolar4)
	log.Debugf("todayAcEnergySolar5: %f", todayAcEnergySolar5)

	// arrayOfUInt16[169].intValue() * 24.0F * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 3276800.0F / 1000.0F
	acLoadkWhAcc := float64(binary.LittleEndian.Uint16(VarEnergyUsed.memory.Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR / 1000
	log.Debugf("acLoadkWhAcc: %f", acLoadkWhAcc)

	// arrayOfUInt16[169].intValue() * 24.0F * this.j.CommonScaleForAcVolts * this.j.CommonScaleForAcCurrent / 3276800.0F / 1000.0F
	energyUsed := float64(binary.LittleEndian.Uint16(VarEnergyUsed.memory.Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForAcVolts * CommonScaleForAcCurrent / MAGIC_DC_W_DIVISOR / 1000
	log.Debugf("energyUsed: %f", energyUsed)

	// (this.g.convert2UShortsInto1Uint(arrayOfUInt16[160], arrayOfUInt16[161]) * 24.0F * this.j.CommonScaleForDcVolts * this.j.CommonScaleForDcCurrent / 3276800.0F / 1000.0F)
	battInkWhAcc := float64(convert2UShortsInto1Uint(binary.LittleEndian.Uint16(VarBattInkWhAcc1.memory.Data()), binary.LittleEndian.Uint16(VarBattInkWhAcc2.memory.Data()))) * MAGIC_WH_MULTIPLIER * CommonScaleForDcVolts * CommonScaleForDcCurrent / MAGIC_DC_V_DIVISOR / 1000
	log.Debugf("battInkWhAcc: %f", battInkWhAcc)

	// (this.g.convert2UShortsInto1Uint(arrayOfUInt16[162], arrayOfUInt16[163]) * 24.0F * this.j.CommonScaleForDcVolts * this.j.CommonScaleForDcCurrent / 3276800.0F / 1000.0F)
	battOutkWhAcc := float64(convert2UShortsInto1Uint(binary.LittleEndian.Uint16(VarBattOutkWhAcc1.memory.Data()), binary.LittleEndian.Uint16(VarBattOutkWhAcc2.memory.Data()))) * MAGIC_WH_MULTIPLIER * CommonScaleForDcVolts * CommonScaleForDcCurrent / MAGIC_DC_V_DIVISOR / 1000
	log.Debugf("battOutkWhAcc: %f", battOutkWhAcc)

	batteryEnergyInToday := float64(binary.LittleEndian.Uint32(VarBatteryEnergyInToday.Memory().Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForDcVolts * CommonScaleForDcCurrent / MAGIC_WH_DIVISOR
	log.Debugf("batteryEnergyInToday: %f", batteryEnergyInToday/1000)

	payload := StatsPayload{
		DcVolts:                 dcVolts,
		DcBatteryPower:          dcBatteryPower,
		InverterRunHrsTotalAcc:  inverterRunHrsTotalAcc,
		ChargeStatus:            chargeStatus,
		DaysSinceEqualise:       daysSinceEqualise,
		NowTotalAcPowerSolar:    nowTotalAcPowerSolar,
		NowAcPowerSolar1:        nowAcPowerSolar1,
		NowAcPowerSolar2:        nowAcPowerSolar2,
		NowAcPowerSolar3:        nowAcPowerSolar3,
		NowAcPowerSolar4:        nowAcPowerSolar4,
		NowAcPowerSolar5:        nowAcPowerSolar5,
		BattSocPercent:          battSocPercent,
		DcCurrent:               dcCurrent,
		DcBatteryCurrent:        dcBatteryCurrent,
		BatteryVolts:            batteryVolts,
		AcLoadVoltage:           acLoadVoltage,
		LoadAcPower:             loadAcPower,
		PowerUsed:               powerUsed,
		AcInverterRmsAmps:       acInverterRmsAmps,
		AcCurrent:               acCurrent,
		AcGeneratorPower:        acGeneratorPower,
		AcGeneratorPower5minAvg: acGeneratorPower,
		AcGeneratorRmsVolts:     acGeneratorRmsVolts,
		AcGeneratorRmsAmps:      acGeneratorRmsAmps,
		NowPercentSolar:         nowPercentSolar,
		TodayTotalAcEnergySolar: todayTotalAcEnergySolar,
		EnergyGenerated:         energyGenerated,
		TodayAcEnergySolar1:     todayAcEnergySolar1,
		TodayAcEnergySolar2:     todayAcEnergySolar2,
		TodayAcEnergySolar3:     todayAcEnergySolar3,
		TodayAcEnergySolar4:     todayAcEnergySolar4,
		TodayAcEnergySolar5:     todayAcEnergySolar5,
		AcLoadkWhAcc:            acLoadkWhAcc,
		EnergyUsed:              energyUsed,
		BattInkWhAcc:            battInkWhAcc,
		BattOutkWhAcc:           battOutkWhAcc,
		BatteryEnergyInToday:    batteryEnergyInToday,
	}

	// Marshaling payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("Error marshaling JSON: %v", err)
		return
	}

	// Logging JSON payload
	log.Infof("Stats payload: %s", jsonPayload)

	// Optionally, write JSON payload to a file
	file, err := os.Create("stats_payload.json")
	if err != nil {
		log.Errorf("Error creating file: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonPayload)
	if err != nil {
		log.Errorf("Error writing JSON to file: %v", err)
	}
}
