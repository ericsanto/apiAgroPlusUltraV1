package requests

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

type AgricultureCultureRequest struct {
	Name                      string `json:"name" validate:"required,min=3,max=100"`
	Variety                   string `json:"variety" validate:"required"`
	SoilTypeId                uint   `json:"soil_type_id" validate:"required"`
	SoilType                  entities.SoilTypeEntity
	Region                    enums.Region  `json:"region" validate:"required"`
	UseType                   enums.UseType `json:"use_type" validate:"required"`
	PhIdealSoil               float32       `json:"ph_ideal_soil" validate:"required,min=0,max=14"`
	MaxTemperature            float32       `json:"max_temperature" validate:"required"`
	MinTemperature            float32       `json:"min_temperature" validate:"required"`
	ExcellentTemperature      float32       `json:"excellent_temperature" validate:"required"`
	WeeklyWaterRequirementMax float32       `json:"weekly_water_requirement_max" validate:"required,gtfield=WeeklyWaterRequirementMin"`
	WeeklyWaterRequirementMin float32       `json:"weekly_water_requirement_min" validate:"required"`
	SunlightRequirement       uint          `json:"sunlight_requirement" validate:"required"`
}
