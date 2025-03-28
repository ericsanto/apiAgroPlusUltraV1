package entities


type SoilTypeEntity struct {
  Id          uint  `gorm:"primaryKey;autoIncrement"` 
  Name        string `gorm:"size:100;not null;unique"`
  Description string `gorm:"type:text;not null"`
  AgricultureCultures []AgricultureCultureEntity `gorm:"foreignKey:SoilTypeId"`
}
