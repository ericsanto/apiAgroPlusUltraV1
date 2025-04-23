package entities

import (
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"
)

type PerformancePlantingEntity struct {
	ID                     uint       `gorm:"primaryKey;autoIncrement"`
	PlantingID             uint       `gorm:"unique"`
	ProductionObtained     float64    `gorm:"not null"`
	UnitProductionObtained enums.Unit `gorm:"size:10"`
	HarvestedArea          float64    `gorm:"not null"`
	UnitHarvestedArea      enums.Unit `gorm:"size:10"`
	HarvestedDate          time.Time  `gorm:"type:timestamp"`
}
