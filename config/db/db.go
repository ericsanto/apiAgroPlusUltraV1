package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
)

var DB *gorm.DB

type Database interface {
	Connect() (*gorm.DB, error)
	Migrate(db *gorm.DB)
}

type GORM struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	SSLMode    string
}

func NewDatabase(dbName, dbUser, dbPassword, dbHost, dbPort, sslMode string) Database {

	return &GORM{DBName: dbName, DBUser: dbUser, DBPassword: dbPassword, DBHost: dbHost, DBPort: dbPort, SSLMode: sslMode}
}

func (g *GORM) Connect() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", g.DBHost, g.DBUser, g.DBPassword, g.DBName,
		g.DBPort, g.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar com banco de dados %w", err)
	}

	DB = db

	return db, nil
}

func (g *GORM) Migrate(db *gorm.DB) {
	// Migrar (criar tabelas automaticamente)
	db.AutoMigrate(&entities.SoilTypeEntity{}, &entities.AgricultureCultureEntity{}, &entities.TypePestEntity{},
		&entities.PestEntity{}, &entities.PestAgricultureCulture{}, &entities.IrrigationRecomendedEntity{},
		&entities.AgricultureCultureIrrigation{}, &entities.SustainablePestControlEntity{}, &entities.AgricultureCulturePestMethodEntity{},
		&entities.FarmEntity{}, &entities.BatchEntity{},
		&entities.IrrigationTypeEntity{}, &entities.PlantingEntity{}, &entities.ProductionCostEntity{}, &entities.SalePlantingEntity{}, &entities.PerformancePlantingEntity{})
}
