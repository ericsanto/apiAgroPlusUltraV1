package entities

type BatchEntity struct {
	ID       uint             `gorm:"primaryKey;autoIncrement"`
	Name     string           `gorm:"size:50;not null;unique"`
	Area     float32          `gorm:"not null"`
	Unit     string           `gorm:"size:10;not null"`
	Planting []PlantingEntity `gorm:"foreignKey:BatchID"`
	FarmID   uint
}
