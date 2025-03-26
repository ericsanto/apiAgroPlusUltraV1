package entities


type SoilTypeEntity struct {
  Id          uint  `gorm:"primaryKey"` 
  Name        string `gorm:"size:100;not null"`
  Description string `gorm:"type:text;not null"`
  AgricultureCultures []AgricultureCultureEntity `gorm:"foreignKey:SoilTypeId"`
}
