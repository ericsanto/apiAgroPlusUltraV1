package config

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		&entities.AgricultureCultureIrrigation{}, &entities.SustainablePestControlEntity{}, &entities.AgricultureCulturePestMethodEntity{}, &entities.BatchEntity{},
		&entities.PlantingEntity{}, &entities.ProductionCostEntity{}, &entities.SalePlantingEntity{}, &entities.PerformancePlantingEntity{}, &entities.FarmEntity{})

	DB = db

	return nil
}
