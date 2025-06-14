package entities

import (
	"time"
)

type PlantingEntity struct {
	ID                   uint `gorm:"primaryKey;autoIncrement"`
	BatchID              uint
	AgricultureCultureID uint
	IsPlanting           bool      `gorm:"default:false"`
	StartDatePlanting    time.Time `gorm:"type:timestamp"`
	ExpectedProduction   float64   `gorm:"not null"`
	SpaceBetweenPlants   float64   `gorm:"not null"`
	SpaceBetweenRows     float64   `gorm:"not null"`
	IrrigationTypeID     uint
	ProductionCost       []ProductionCostEntity      `gorm:"foreignKey:PlantingID"`
	SalePlanting         []SalePlantingEntity        `gorm:"foreignKey:PlantingID"`
	PerformancePlanting  []PerformancePlantingEntity `gorm:"foreignKey:PlantingID"`
}
