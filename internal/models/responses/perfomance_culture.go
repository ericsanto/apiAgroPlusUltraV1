package responses

import (
	"time"
)

type PerformanceCultureResponse struct {
	Planting                   BatchPlantiesResponse `json:"planting"`
	ID                         uint                  `json:"id"`
	ProductionObtained         float64               `json:"production_obtained"`
	ProductionObtainedFormated string                `json:"production_obtained_formated"`
	HarvestedArea              float64               `json:"harvested_area"`
	HarvestedAreaFormated      string                `json:"harvested_area_formated"`
	HarvestedDate              time.Time             `json:"harvested_date"`
}

type DbResultPerformancePlanting struct {
	BatchName                  string
	AgricultureCultureName     string
	StartDatePlanting          time.Time
	IsPlanting                 bool
	ID                         uint
	ProductionObtained         float64
	ProductionObtainedFormated string
	HarvestedArea              float64
	HarvestedAreaFormated      string
	HarvestedDate              time.Time
}
