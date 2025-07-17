package repositories

import (
	"context"
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
)

type FarmRepositoryInterface interface {
	FindByID(userID, id uint) (*responses.FarmResponse, error)
	FindAll(userID uint) ([]responses.FarmResponse, error)
	Create(ctx context.Context, farmEntity entities.FarmEntity) error
}

type FarmRepository struct {
	db interfaces.GORMRepositoryInterface
}

func NewFarmRepository(db interfaces.GORMRepositoryInterface) FarmRepositoryInterface {
	return &FarmRepository{db: db}
}

func (fr *FarmRepository) FindByID(userID, id uint) (*responses.FarmResponse, error) {

	var entityFarm responses.FarmResponse

	query := `SELECT farm_entities.id AS id,
	farm_entities.name AS name
	FROM farm_entities
	INNER JOIN user_models ON user_models.id = farm_entities.user_id
	WHERE farm_entities.id = ? AND farm_entities.user_id = ?`

	err := fr.db.Raw(query, id, userID).Scan(&entityFarm)

	if err.Error != nil {
		return nil, fmt.Errorf("erro ao buscar fazenda")
	}

	return &entityFarm, nil
}

func (fr *FarmRepository) FindAll(userID uint) ([]responses.FarmResponse, error) {

	var listFarmResponse []responses.FarmResponse

	query := `SELECT farm_entities.id AS id,
	farm_entities.name AS name
	FROM farm_entities
	INNER JOIN user_models ON user_models.id = farm_entities.user_id
	WHERE farm_entities.user_id = ?`

	if err := fr.db.Raw(query, userID).Scan(&listFarmResponse).Error; err != nil {
		return nil, fmt.Errorf("erro ao buscar todas as fazendas")
	}

	return listFarmResponse, nil
}

func (fr *FarmRepository) Create(ctx context.Context, farmEntity entities.FarmEntity) error {

	if err := fr.db.WithContext(ctx).Create(&farmEntity).Error; err != nil {

		return fmt.Errorf("nao foi possivel criar fazenda: ")

	}

	return nil
}
