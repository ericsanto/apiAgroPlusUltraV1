package di

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/controllers"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
)

type BatchBuilder struct{}

func NewBatchBuilder() *BatchBuilder {
	return &BatchBuilder{}
}

func (bb *BatchBuilder) Builder() controllers.BatchControllerInterface {

	respositoryBatch := repositories.NewBatchRepository(db.DB)
	serviceBatch := services.NewBatchService(respositoryBatch)
	controllerBatch := controllers.NewBatchController(serviceBatch)

	return controllerBatch
}
