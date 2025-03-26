package requests

type AgricultureCultureRequest struct {
 
  Name                       string                 `json:"name" binding:"required,min=5,max=255"`
  NameCientific              string                 `json:"name_cientific" binding:"required,min=5,max=255"`
  SoilTypeId                 uint                    `json:"soil_type_id" binding:"required"`
  PhIdealSoil                float32                 `json:"ph_ideal_soil" binding:"required"`
  MaxTemperature             float32                 `json:"max_temperature" binding:"required"`
  MinTemperature             float32                 `json:"min_temperature" binding:"required"`
  ExcellentTemperature       float32                 `json:"excellent_temperature" binding:"required"`
  WeeklyWaterRequirememntMax float32                 `json:"weekly_water_requirement_max" binding:"required"`
  WeeklyWaterRequirememntMin float32                 `json:"weekly_water_requirement_min" binding:"required"`
  SunlightRequirement        uint                    `json:"sunlight_requirement" binding:"required"`

}

