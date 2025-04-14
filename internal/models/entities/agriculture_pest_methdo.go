package entities

type AgricultureCulturePestMethodEntity struct {
	AgricultureCultureId     uint   `gorm:"primaryKey"`
	PestId                   uint   `gorm:"primaryKey"`
	SustainablePestControlId uint   `gorm:"primaryKey"`
	Description              string `gorm:"text;not null"`

	AgricultureCulture     AgricultureCultureEntity     `gorm:"foreignKey:AgricultureCultureId;references:Id"`
	Pest                   PestEntity                   `gorm:"foreingKey:PestId;references:Id"`
	SustainablePestControl SustainablePestControlEntity `gorm:"foreingKey:SustainablePestControlId;references:Id"`
}
