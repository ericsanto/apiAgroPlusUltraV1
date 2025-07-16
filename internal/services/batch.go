package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type BatchServiceInterface interface {
	PostBatchService(userID, farmID uint, requestBatchService requests.BatchRequest) error
	GetAllBatch(userID, farmID uint) ([]responses.BatchResponse, error)
	GetBatchFindById(userID, farmID, batchID uint) (*responses.BatchResponse, error)
	PutBatch(userID, farmID, batchID uint, requestBatch requests.BatchRequest) error
	DeleteBatch(userID, farmID, batchID uint) error
}

type BatchService struct {
	batchRepository repositories.BatchRepositoryInterface
}

func NewBatchService(batchRepository repositories.BatchRepositoryInterface) BatchServiceInterface {
	return &BatchService{batchRepository: batchRepository}
}

func (b *BatchService) PostBatchService(userID, farmID uint, requestBatchService requests.BatchRequest) error {

	batchEntity := entities.BatchEntity{
		Name:   requestBatchService.Name,
		Area:   requestBatchService.Area,
		Unit:   requestBatchService.Unit,
		FarmID: farmID,
	}

	if err := b.batchRepository.Create(userID, farmID, batchEntity); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (b *BatchService) GetAllBatch(userID, farmID uint) ([]responses.BatchResponse, error) {

	var listResponseBatchs []responses.BatchResponse

	batchs, err := b.batchRepository.FindAllBatch(userID, farmID)
	if err != nil {
		return nil, fmt.Errorf("erro: %v", err)
	}

	for _, v := range batchs {
		responseBatch := responses.BatchResponse{
			ID:   v.ID,
			Name: v.Name,
			Area: v.Area,
			Unit: v.Unit,
		}

		listResponseBatchs = append(listResponseBatchs, responseBatch)

	}

	return listResponseBatchs, nil

}

func (b *BatchService) GetBatchFindById(userID, farmID, batchID uint) (*responses.BatchResponse, error) {

	batch, err := b.batchRepository.BatchFindById(userID, farmID, batchID)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	responseBatch := responses.BatchResponse{
		ID:   batch.ID,
		Name: batch.Name,
		Area: batch.Area,
		Unit: batch.Unit,
	}

	return &responseBatch, nil
}

func (b *BatchService) PutBatch(userID, farmID, batchID uint, requestBatch requests.BatchRequest) error {

	entitieBatch := entities.BatchEntity{
		Name: requestBatch.Name,
		Area: requestBatch.Area,
		Unit: requestBatch.Unit,
	}

	if err := b.batchRepository.Update(userID, farmID, batchID, entitieBatch); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (b *BatchService) DeleteBatch(userID, farmID, batchID uint) error {

	if err := b.batchRepository.DeleteBatch(userID, farmID, batchID); err != nil {
		return fmt.Errorf("falha ao deletar: %w", err)
	}

	return nil
}
