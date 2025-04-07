package requests

type IrrigationRecomendedRequest struct {
	PhenologicalPhase string  `json:"phenological_phase" validate:"required"`
	PhaseDurationDays int     `json:"phase_duration_days" validate:"required"`
	IrrigationMax     float32 `json:"irrigation_max" validate:"required"`
	IrrigationMin     float32 `json:"irrigation_min" validate:"required"`
	Description       string  `json:"description" validate:"required"`
	Unit              string  `json:"unit" validate:"required"`
}
