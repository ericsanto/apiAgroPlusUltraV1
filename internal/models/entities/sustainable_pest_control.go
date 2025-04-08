package entities

type SustainablePestControlEntity struct {
	Id   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"size=100;not null;unique"`
}
