package config

import (
		"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB


func Conect() error {

  dsn := "host=localhost user=teste password=teste dbname=teste port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}

	// Migrar (criar tabelas automaticamente)
	db.AutoMigrate(&entities.SoilTypeEntity{}, &entities.AgricultureCultureEntity{})

  DB = db

  return nil
}

