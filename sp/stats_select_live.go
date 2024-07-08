package sp

import (
	"encoding/binary"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const MAGIC = 32768.0
const MAGIC_AC_W_DIVISOR = MAGIC * 800.0
const MAGIC_AC_W_SIGNED_DIVISOR = MAGIC * 100.0
const MAGIC_DC_W_DIVISOR = MAGIC * 100.0
const MAGIC_DC_V_DIVISOR = MAGIC * 10.0
const MAGIC_WH_MULTIPLIER = 24.0
const MAGIC_WH_DIVISOR = MAGIC * 100.0
const MAGIC_TEMPERATURE_DIVISOR = MAGIC
const MAGIC_RATIO_DIVISOR = 25600.0

func StatsSelectLiveRender(protocol *Protocol) string {
	variables := []*Variable{
		&VarBatteryEnergyInToday,
		&VarBatteryEnergyInTotal,
		&VarBatterySoc,
		&VarCommonScaleForDcVolts,
		&VarCommonScaleForDcCurrent,
		&VarACLoadkWhTotalAcc,
		&VarLoadAcPower,
	}

	log.Debugf("Querying variables: %v", variables)
	protocol.Query(variables)

	// "battery_in_wh_today": vars["DCkWhInToday"].get_value(self.scales) / 1000,
	// "battery_in_wh_total": vars["BattInkWhTotalAcc"].get_value(self.scales) / 1000,
	// "battery_out_wh_today": vars["DCkWhOutToday"].get_value(self.scales) / 1000,
	// "battery_out_wh_total": vars["BattOutkWhTotalAcc"].get_value(self.scales) / 1000,
	// "battery_soc": vars["BattSocPercent"].get_value(self.scales),
	// "battery_w": vars["DCBatteryPower"].get_value(self.scales),
	// "grid_in_wh_today": vars["ACInputWhTodayAcc"].get_value(self.scales) / 1000,
	// "grid_in_wh_total": vars["ACInputWhTotalAcc"].get_value(self.scales) / 1000,
	// "grid_out_wh_today": vars["ACExportWhTodayAcc"].get_value(self.scales) / 1000, # unverified guess
	// "grid_out_wh_total": vars["ACExportWhTotalAcc"].get_value(self.scales) / 1000, # unverified guess
	// "grid_w": vars["ACGeneratorPower"].get_value(self.scales),
	// "load_w": vars["LoadAcPower"].get_value(self.scales),
	// "load_wh_today": vars["ACLoadWhAcc"].get_value(self.scales) / 1000,
	// "load_wh_total": vars["ACLoadkWhTotalAcc"].get_value(self.scales) / 1000,
	// "shunt_w": 0 - vars["Shunt1Power"].get_value(self.scales),
	// # TODO: assumes shunt 1 is always a solar shunt
	// "solar_wh_today": ( vars["ACSolarWhTodayAcc"].get_value(self.scales) + 0 - vars["Shunt1WhTodayAcc"].get_value(self.scales)) / 1000,
	// "solar_wh_total": (vars["ACSolarWhTotalAcc"].get_value(self.scales) + (0 -vars["Shunt1WhTotalAcc"].get_value(self.scales))) / 1000,
	// "solarinverter_w": vars["CombinedKacoAcPowerHiRes"].get_value(self.scales),

	bytes := VarBatterySoc.Memory().Data()
	batterySoc := (float64(bytes[0]) + float64(bytes[1])*256) / MAGIC_RATIO_DIVISOR

	// TODO: shift at least the integer conversion up to variable somehow
	CommonScaleForDcVolts := float64(binary.LittleEndian.Uint16(VarCommonScaleForDcVolts.memory.Data()))
	CommonScaleForDcCurrent := float64(binary.LittleEndian.Uint16(VarCommonScaleForDcCurrent.memory.Data()))
	batteryEnergyInToday := float64(binary.LittleEndian.Uint32(VarBatteryEnergyInToday.Memory().Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForDcVolts * CommonScaleForDcCurrent / MAGIC_WH_DIVISOR

	log.Debugf("CommonScaleForDcVolts: %f", CommonScaleForDcVolts)
	log.Debugf("CommonScaleForDcCurrent: %f", CommonScaleForDcCurrent)

	// Testing
	acLoadkWhTotalAcc := float64(binary.LittleEndian.Uint32(VarACLoadkWhTotalAcc.Memory().Data()))
	log.Debugf("AC Lifetime Solar Energy: %f", acLoadkWhTotalAcc)
	loadAcPower := float64(binary.LittleEndian.Uint32(VarLoadAcPower.Memory().Data()))
	log.Debugf("AC Load Power: %f", loadAcPower)

	return fmt.Sprintf(
		"Battery in kWh today: %f\nBattery SoC %%: %f\n",
		batteryEnergyInToday/1000,
		batterySoc*100,
	)
}

func StatsSelectLiveRenderV2(protocol *Protocol) {
	variables := []*Variable{
		&VarBatteryEnergyInToday,
		&VarBatteryEnergyInTotal,
		&VarBatterySoc,
		&VarCommonScaleForDcVolts,
		&VarCommonScaleForDcCurrent,
		&VarACLoadkWhTotalAcc,
		&VarLoadAcPower,
		&TechnicalData,
		&TechnicalDataTab,
	}

	log.Debugf("Querying variables: %v", variables)
	protocol.Query(variables)

	// Testing outputs

	CommonScaleForDcVolts := float64(binary.LittleEndian.Uint16(VarCommonScaleForDcVolts.memory.Data()))
	log.Debugf("CommonScaleForDcVolts: %f", CommonScaleForDcVolts)

	CommonScaleForDcCurrent := float64(binary.LittleEndian.Uint16(VarCommonScaleForDcCurrent.memory.Data()))
	log.Debugf("CommonScaleForDcCurrent: %f", CommonScaleForDcCurrent)

	bytes := VarBatterySoc.Memory().Data()
	batterySoc := (float64(bytes[0]) + float64(bytes[1])*256) / MAGIC_RATIO_DIVISOR
	log.Debugf("batterySoc: %f", batterySoc)

	batteryEnergyInToday := float64(binary.LittleEndian.Uint32(VarBatteryEnergyInToday.Memory().Data())) * MAGIC_WH_MULTIPLIER * CommonScaleForDcVolts * CommonScaleForDcCurrent / MAGIC_WH_DIVISOR
	log.Debugf("batteryEnergyInToday: %f", batteryEnergyInToday)

	acLoadkWhTotalAcc := float64(binary.LittleEndian.Uint32(VarACLoadkWhTotalAcc.Memory().Data()))
	log.Debugf("acLoadkWhTotalAcc: %f", acLoadkWhTotalAcc)

	loadAcPower := float64(binary.LittleEndian.Uint32(VarLoadAcPower.Memory().Data()))
	log.Debugf("loadAcPower: %f", loadAcPower)

	// Testing TechnicalDataTab
	TechnicalDataTabBytes := TechnicalDataTab.Memory().Data()
	log.Debugf("TechnicalDataTabBytes: %v", TechnicalDataTabBytes)
	TechnicalDataTab := string(TechnicalDataTabBytes)
	log.Debugf("TechnicalDataTab: %v", TechnicalDataTab)

}
