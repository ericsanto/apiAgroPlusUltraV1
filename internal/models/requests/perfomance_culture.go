package requests

import (
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"
)

type PerformancePlantingRequest struct {
	ProductionObtained     float64    `json:"production_obtained" validate:"required"`
	UnitProductionObtained enums.Unit `json:"unit_production_obtained" validate:"required"`
	HarvestedArea          float64    `json:"harvested_area" validate:"required"`
	UnitHarvestedArea      enums.Unit `json:"unit_harvested_area" validate:"required"`
	HarvestedDate          time.Time  `json:"harvested_date" validate:"required"`
}
