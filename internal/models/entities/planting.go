package entities

import "time"

type PlantingEntity struct {
	ID                   uint `gorm:"primaryKey;autoIncrement"`
	BatchID              uint
	AgricultureCultureID uint
	IsPlanting           bool                   `gorm:"default:false"`
	StartDatePlanting    time.Time              `gorm:"type:timestamp"`
	ProductionCost       []ProductionCostEntity `gorm:"foreignKey:PlantingID"`
}
