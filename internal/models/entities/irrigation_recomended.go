package entities

type IrrigationRecomendedEntity struct {
	Id                       uint                           `gorm:"primaryKey;autoIncrement"`
	PhenologicalPhase        string                         `gorm:"size:100;not null"`
	PhaseDurationDays        int                            `gorm:"not null"`
	IrrigationMax            float32                        `gorm:"not null"`
	IrrigationMin            float32                        `gorm:"not null"`
	Unit                     string                         `gorm:"not null"`
	Description              string                         `gorm:"size:100"`
	AgricultureCultureEntity []AgricultureCultureIrrigation `gorm:"foreignKey:IrrigationRecomendedId"`
}
