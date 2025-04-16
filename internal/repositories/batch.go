package repositories

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"gorm.io/gorm"
)

type BatchRepository struct {
	db *gorm.DB
}

func NewBatchRepository(db *gorm.DB) *BatchRepository {
	return &BatchRepository{db: db}
}

func (b *BatchRepository) Create(batchEntity entities.BatchEntity) error {

	errorDatabase := myerror.MessageErrorDuplicateKeyViolatesUniqueConstraint()

	if err := b.db.Create(&batchEntity).Error; err != nil {
		if strings.Contains(err.Error(), errorDatabase) {
			return fmt.Errorf("já existe lote cadastrado com esse nome")
		}
		return fmt.Errorf("erro ao cadastrar objeto")
	}

	return nil
}

func (b *BatchRepository) FindAllBatch() ([]entities.BatchEntity, error) {

	var entityBatch []entities.BatchEntity

	if err := b.db.Find(&entityBatch).Error; err != nil {
		return entityBatch, fmt.Errorf("erro ao buscar lotes cadastrados")
	}

	return entityBatch, nil
}

func (b *BatchRepository) BatchFindById(id uint) (*entities.BatchEntity, error) {

	var entityBatch entities.BatchEntity

	if err := b.db.First(&entityBatch, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("não existe lote com o ID %d. %w", id, gorm.ErrRecordNotFound)
		}
		return nil, fmt.Errorf("erro ao buscar lote %w", err)
	}

	return &entityBatch, nil
}

func (b *BatchRepository) Update(id uint, entityBatch entities.BatchEntity) error {

	_, err := b.BatchFindById(id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	errorDatabase := myerror.MessageErrorDuplicateKeyViolatesUniqueConstraint()

	if err := b.db.Model(entities.BatchEntity{}).Where("id = ?", id).Updates(&entityBatch).Error; err != nil {
		if strings.Contains(err.Error(), errorDatabase) {
			return fmt.Errorf("já existe lote cadastrado com esse nome %w", err)
		}
		return fmt.Errorf("erro ao atualizar objeto: %w", err)
	}

	return nil
}

func (b *BatchRepository) DeleteBatch(id uint) error {

	batch, err := b.BatchFindById(id)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := b.db.Where("id = ?", id).Delete(&batch).Error; err != nil {
		return fmt.Errorf("erro ao deletar lote")
	}

	return nil
}
