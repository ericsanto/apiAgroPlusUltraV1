package entities

import "github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"

type AgricultureCultureEntity struct {
	Id                        uint             `gorm:"primaryKey;autoIncrement"`
	Name                      string           `gorm:"size:255;not null"`
	Variety                   string           `gorm:"size:255;not null;unique"`
	UseType                   enums.UseType    `gorm,:"size:100;not null"`
	Region                    enums.Region     `gorm,:"size:100;not null"`
	Planting                  []PlantingEntity `gorm:"foreignKey:AgricultureCultureID"`
	SoilTypeId                uint
	SoilTypeEntity            SoilTypeEntity                 `gorm:"foreignKey:SoilTypeId;references:Id"`
	PhIdealSoil               float32                        `gorm:"not null"`
	MaxTemperature            float32                        `gorm:"not null"`
	MinTemperature            float32                        `gorm:"not null"`
	ExcellentTemperature      float32                        `gorm:"not null"`
	WeeklyWaterRequirementMax float32                        `gorm:"not null"`
	WeeklyWaterRequirementMin float32                        `gorm:"not null"`
	SunlightRequirement       uint                           `gorm:"not null"`
	Pests                     []PestAgricultureCulture       `gorm:"foreignKey:AgricultureCultureId"`
	IrrigationRecomended      []AgricultureCultureIrrigation `gorm:"foreignKey:AgricultureCultureId"`
}
