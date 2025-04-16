package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
)

type BatchService struct {
	batchRepository *repositories.BatchRepository
}

func NewBatchService(batchRepository *repositories.BatchRepository) *BatchService {
	return &BatchService{batchRepository: batchRepository}
}

func (b *BatchService) PostBatchService(requestBatchService requests.BatchRequest) error {

	batchEntity := entities.BatchEntity{
		Name: requestBatchService.Name,
		Area: requestBatchService.Area,
		Unit: requestBatchService.Unit,
	}

	if err := b.batchRepository.Create(batchEntity); err != nil {
		return fmt.Errorf("erro: %v", err)
	}

	return nil
}

func (b *BatchService) GetAllBatch() ([]responses.BatchResponse, error) {

	var listResponseBatchs []responses.BatchResponse

	batchs, err := b.batchRepository.FindAllBatch()
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

func (b *BatchService) GetBatchFindById(id uint) (*responses.BatchResponse, error) {

	batch, err := b.batchRepository.BatchFindById(id)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	responseBatch := responses.BatchResponse{
		Name: batch.Name,
		Area: batch.Area,
		Unit: batch.Unit,
	}

	return &responseBatch, nil
}

func (b *BatchService) PutBatch(id uint, requestBatch requests.BatchRequest) error {

	entitieBatch := entities.BatchEntity{
		Name: requestBatch.Name,
		Area: requestBatch.Area,
		Unit: requestBatch.Unit,
	}

	if err := b.batchRepository.Update(id, entitieBatch); err != nil {
		return fmt.Errorf("falha ao atualizar dados: %w", err)
	}

	return nil
}

func (b *BatchService) DeleteBatch(id uint) error {

	if err := b.batchRepository.DeleteBatch(id); err != nil {
		return fmt.Errorf("falha ao deletar: %w", err)
	}

	return nil
}
