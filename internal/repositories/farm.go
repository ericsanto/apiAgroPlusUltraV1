package repositories

import (
	"fmt"
	"log"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type FarmRepositoryInterface interface {
	FindByID(userID, id uint) (*responses.FarmResponse, error)
	FindAll(userID uint) ([]responses.FarmResponse, error)
	Create(farmEntity entities.FarmEntity) error
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

	if err.RowsAffected == 0 {
		return nil, fmt.Errorf("%w %d", myerror.ErrFarmNotFound, id)
	}

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
