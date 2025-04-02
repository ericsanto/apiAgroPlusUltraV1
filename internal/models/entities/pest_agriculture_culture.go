package entities

type PestAgricultureCulture struct {
	PestId               uint   `gorm:"primaryKey"`
	AgricultureCultureId uint   `gorm:"primaryKey"`
	Description          string `gorm:"type:text;not null"`
	ImageUrl             string `gorm:"type:text;not null"`

	Pest               PestEntity               `gorm:"foreignKey:PestId;references:Id"`
	AgricultureCulture AgricultureCultureEntity `gorm:"foreignKey:AgricultureCultureId;references:Id"`
}
