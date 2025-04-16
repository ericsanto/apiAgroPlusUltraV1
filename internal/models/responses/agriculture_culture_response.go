package responses

import "github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"

type AgricultureCultureResponse struct {
	Id                         uint
	Name                       string
	Variety                    string
	Region                     enums.Region
	UseType                    enums.UseType
	SoilTypeId                 uint
	PhIdealSoil                float32
	MaxTemperature             float32
	MinTemperature             float32
	ExcellentTemperature       float32
	WeeklyWaterRequirememntMax float32
	WeeklyWaterRequirememntMin float32
	SunlightRequirement        uint
}
