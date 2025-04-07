package entities

type AgricultureCultureIrrigation struct {
	AgricultureCultureId   uint `gorm:"primaryKey"`
	IrrigationRecomendedId uint `gorm:"primaryKey"`

	AgricultureCulture   AgricultureCultureEntity   `gorm:"foreignKey:AgricultureCultureId;references:Id"`
	IrrigationRecomended IrrigationRecomendedEntity `gorm:"foreignKey:IrrigationRecomendedId;references:Id"`
}
