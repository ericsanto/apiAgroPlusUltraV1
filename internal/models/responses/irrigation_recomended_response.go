package responses

type IrrigationRecomendedResponse struct {
	Id                uint    `json:"id"`
	PhenologicalPhase string  `json:"phenological_phase"`
	PhaseDurationDays int     `json:"phase_duration_days"`
	IrrigationMax     float32 `json:"irrigation_max"`
	IrrigationMin     float32 `json:"irrigation_min"`
	Description       string  `json:"description"`
	Unit              string  `json:"unit"`
}
