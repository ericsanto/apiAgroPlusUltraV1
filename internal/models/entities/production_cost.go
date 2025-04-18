package entities

import "time"

type ProductionCostEntity struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	PlantingID  uint
	Item        string    `gorm:"size:100;not null"`
	Unit        string    `gorm:"size:50;not null"`
	Quantity    float32   `gorm:"not null"`
	CostPerUnit float32   `gorm:"not null"`
	CostDate    time.Time `gorm:"type:timestamp;not null"`
}
