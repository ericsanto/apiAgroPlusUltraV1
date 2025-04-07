package responses

type AgricultureCultureIrrigationResponse struct {
	Name              string  `json:"name"`
	PhenologicalPhase string  `json:"pheneological_phase"`
	PhaseDurationDays int     `json:"phase_duration_days"`
	IrrigationMax     float32 `json:"irrigation_max"`
	IrrigationMin     float32 `json:"irrigation_min"`
	Unit              string  `json:"unit"`

	// WeeklyWaterRequirementMax float32 `json:"weekly_water_requirement_max"`
	// WeeklyWaterRequirementMin float32 `json:"weekly_water_requirement_min"`
}
