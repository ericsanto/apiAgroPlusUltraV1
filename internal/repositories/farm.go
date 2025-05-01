package repositories

import (
	"fmt"
	"log"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"gorm.io/gorm"
)

type FarmRepository struct {
	db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) *FarmRepository {
	return &FarmRepository{db: db}
}

func (fr *FarmRepository) FindByID(id float64) (*entities.FarmEntity, error) {

	var entityFarm entities.FarmEntity

	if err := fr.db.First(&entityFarm, id).Error; err != nil {
		return nil, fmt.Errorf("fazenda com id %f %w", id, myerror.ErrNotFound)
	}

	return &entityFarm, nil
}

func (fr *FarmRepository) Create(farmEntity entities.FarmEntity) error {

	query := `SELECT  EXISTS(SELECT 1 FROM user_models WHERE id = ?)`

	var exists bool

	if err := fr.db.Raw(query, farmEntity.UserID).Scan(&exists).Error; err != nil {
		log.Println("%w", err)
		return fmt.Errorf("não foi possível verificar a existência do usuário")
	}

	fmt.Println(exists)

	if !exists {
		return fmt.Errorf("usuário com id %d %w", farmEntity.UserID, myerror.ErrNotFound)
	}

	if err := fr.db.Create(&farmEntity).Error; err != nil {
		log.Println(err.Error())
		return fmt.Errorf("erro ao criar fazenda")

	}

	return nil
}
