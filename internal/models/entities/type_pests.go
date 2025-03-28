package entities


type TypePestEntity struct {

  Id   uint   `gorm:"primaryKey;autoIncrement"`
  Name string `gorm:"size:100; not null;unique"`
  Pests []PestEntity `gorm:"foreingKey:TypePestId"`
}
