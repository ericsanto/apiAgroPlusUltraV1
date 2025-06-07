package entities

type FarmEntity struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"size:150;not null"`
	UserID uint
	Batch  []BatchEntity `gorm:"foreignKey:FarmID"`
}
