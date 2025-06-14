package entities

import "github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"

type IrrigationTypeEntity struct {
	ID       uint                 `gorm:"primaryKey;autoIncrement"`
	Name     enums.IrrigationType `gorm:"size:50;not null"`
	Planting []PlantingEntity     `gorm:"foreignKey:IrrigationTypeID"`
}
