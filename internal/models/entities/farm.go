package entities

type FarmEntity struct {
	ID     uint          `gorm:"primaryKey;autoIncrement"`
	Name   string        `gorm:"size:150;not null;uniqueIndex:idx_user_farm_name"`
	UserID uint          `gorm:"uniqueIndex:idx_user_farm_name"`
	Batch  []BatchEntity `gorm:"foreignKey:FarmID"`
}
