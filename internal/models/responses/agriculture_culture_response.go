package responses



type AgricultureCultureResponse struct {

  Id                         uint            
  Name                       string         
  NameCientific              string         
  SoilTypeId                 uint
  PhIdealSoil                float32        
  MaxTemperature             float32        
  MinTemperature             float32        
  ExcellentTemperature       float32        
  WeeklyWaterRequirememntMax float32        
  WeeklyWaterRequirememntMin float32        
  SunlightRequirement        uint           
}

