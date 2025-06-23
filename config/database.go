package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

var DB *gorm.DB

func Conect() error {

	dsn := "host=db user=go password=go dbname=go port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	// Migrar (criar tabelas automaticamente)
	db.AutoMigrate(&entities.SoilTypeEntity{}, &entities.AgricultureCultureEntity{}, &entities.TypePestEntity{},
		&entities.PestEntity{}, &entities.PestAgricultureCulture{}, &entities.IrrigationRecomendedEntity{},
		&entities.AgricultureCultureIrrigation{}, &entities.SustainablePestControlEntity{}, &entities.AgricultureCulturePestMethodEntity{},
		&entities.FarmEntity{}, &entities.BatchEntity{},
		&entities.IrrigationTypeEntity{}, &entities.PlantingEntity{}, &entities.ProductionCostEntity{}, &entities.SalePlantingEntity{}, &entities.PerformancePlantingEntity{})

	DB = db

	return nil
}
