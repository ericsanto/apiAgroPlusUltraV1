package repositories

import (
	"fmt"
	"strings"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories/interfaces"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type BatchRepositoryInterface interface {
	Create(userID, farmID uint, batchEntity entities.BatchEntity) error
	FindAllBatch(userID, farmID uint) ([]entities.BatchEntity, error)
	BatchFindById(userID, farmID, batchID uint) (*entities.BatchEntity, error)
	Update(userID, farmID, batchID uint, entityBatch entities.BatchEntity) error
	DeleteBatch(userID, farmID, batchID uint) error
}

type BatchRepository struct {
	db             interfaces.GORMRepositoryInterface
	farmRepository FarmRepositoryInterface
}

func NewBatchRepository(db interfaces.GORMRepositoryInterface, farmRepository FarmRepositoryInterface) BatchRepositoryInterface {
	return &BatchRepository{db: db, farmRepository: farmRepository}
}

func (b *BatchRepository) Create(userID, farmID uint, batchEntity entities.BatchEntity) error {

	errorDatabase := myerror.MessageErrorDuplicateKeyViolatesUniqueConstraint()

	_, err := b.farmRepository.FindByID(userID, farmID)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := b.db.Create(&batchEntity).Error; err != nil {
		if strings.Contains(err.Error(), errorDatabase) {
			return myerror.ErrBatchAlreadyExists
		}
		return fmt.Errorf("erro ao cadastrar objeto")
	}

	return nil
}

func (b *BatchRepository) FindAllBatch(userID, farmID uint) ([]entities.BatchEntity, error) {

	var entityBatch []entities.BatchEntity

	query := `SELECT batch_entities.id AS id, 
	batch_entities.name AS name,
	batch_entities.area AS area,
	batch_entities.unit AS unit
	FROM batch_entities 
	INNER JOIN farm_entities ON farm_entities.id = batch_entities.farm_id
	INNER JOIN user_models ON user_models.id = farm_entities.user_id
	WHERE farm_entities.id = ? AND user_models.id = ?`

	if err := b.db.Raw(query, farmID, userID).Scan(&entityBatch).Error; err != nil {
		return entityBatch, fmt.Errorf("erro ao buscar lotes cadastrados")
	}

	return entityBatch, nil
}

func (b *BatchRepository) BatchFindById(userID, farmID, batchID uint) (*entities.BatchEntity, error) {

	var entityBatch entities.BatchEntity

	_, err := b.farmRepository.FindByID(userID, farmID)

	if err != nil {
		return nil, err
	}

	query := `SELECT batch_entities.id AS id,
	batch_entities.name AS name,
	batch_entities.area AS area,
	batch_entities.unit AS unit
	FROM batch_entities 
	INNER JOIN farm_entities ON farm_entities.id = batch_entities.farm_id
	INNER JOIN user_models ON user_models.id = farm_entities.user_id
	WHERE batch_entities.id = ? AND farm_entities.id = ? AND user_models.id = ?`

	tx := b.db.Raw(query, batchID, farmID, userID).Scan(&entityBatch)

	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("nao existe lote com esse id %w", myerror.ErrNotFound)
	}

	if tx.Error != nil {
		return nil, fmt.Errorf("erro ao buscar lote")
	}

	return &entityBatch, nil
}

func (b *BatchRepository) Update(userID, farmID, batchID uint, entityBatch entities.BatchEntity) error {

	_, err := b.BatchFindById(userID, farmID, batchID)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	errorDatabase := myerror.MessageErrorDuplicateKeyViolatesUniqueConstraint()

	query := `UPDATE batch_entities
	SET name = ?, area = ?, unit = ?
	FROM farm_entities, user_models
	WHERE batch_entities.id = ? AND farm_entities.id = ? AND user_models.id =  ?`

	if err := b.db.Exec(query, entityBatch.Name, entityBatch.Area, entityBatch.Unit, batchID, farmID, userID).Error; err != nil {
		if strings.Contains(err.Error(), errorDatabase) {
			return myerror.ErrBatchAlreadyExists
		}

		return fmt.Errorf("erro ao atualizar lote")
	}

	return nil
}

func (b *BatchRepository) DeleteBatch(userID, farmID, batchID uint) error {

	batch, err := b.BatchFindById(userID, farmID, batchID)
	if err != nil {
		return err
	}

	if err := b.db.Delete(&batch).Error; err != nil {
		return fmt.Errorf("erro ao deletar batch")
	}

	return nil
}
