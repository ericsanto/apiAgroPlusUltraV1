package entities

type PestEntity struct {
	Id                  uint   `gorm:"primaryKey;autoIncrement"`
	Name                string `gorm:"size:100;not null"`
	TypePestId          uint
	TypePestEntity      TypePestEntity           `gorm:"foreignKey:TypePestId;references:Id"`
	AgricultureCultures []PestAgricultureCulture `gorm:"foreignKey:PestId"`
}
