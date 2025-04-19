package entities

type SalePlantingEntity struct {
	ID         uint    `gorm:"primaryKey;autoIncrement"`
	PlantingID uint    `gorm:"unique;not null"`
	ValueSale  float32 `gorm:"not null"`
}
