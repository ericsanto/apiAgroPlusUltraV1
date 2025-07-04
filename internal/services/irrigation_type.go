package services

import (
	"fmt"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/enums"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/responses"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/repositories"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type IrrigationTypeServiceInterface interface {
	GetAllIrrigationType() ([]responses.IrrigationTypeResponse, error)
	PostirrigationType(requestIrrigationType requests.IrrigationTypeRequest) error
	GetIrrigationTypeByID(id uint) (*responses.IrrigationTypeResponse, error)
	PutIrrigationType(id uint, requestIrrigationType requests.IrrigationTypeRequest) error
	DeleteIrrigationType(id uint) error
}

type IrrigationTypeService struct {
	irrigationTypeRepository repositories.IrrigationTypeRepositoryInterface
}

func NewIrrigationTypeService(irrigationTypeRepository repositories.IrrigationTypeRepositoryInterface) IrrigationTypeServiceInterface {
	return &IrrigationTypeService{irrigationTypeRepository: irrigationTypeRepository}
}

func (it *IrrigationTypeService) GetAllIrrigationType() ([]responses.IrrigationTypeResponse, error) {

	irrigationTypeEntity, err := it.irrigationTypeRepository.FindAllIrrigationType()
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	var irrigationTypeResponseList []responses.IrrigationTypeResponse

	for _, v := range irrigationTypeEntity {
		irrigationType := responses.IrrigationTypeResponse{
			ID:   v.ID,
			Name: v.Name,
		}

		irrigationTypeResponseList = append(irrigationTypeResponseList, irrigationType)
	}

	return irrigationTypeResponseList, nil
}

func (it *IrrigationTypeService) PostirrigationType(requestIrrigationType requests.IrrigationTypeRequest) error {

	entityirrigationType := entities.IrrigationTypeEntity{
		Name: requestIrrigationType.Name,
	}

	if !enums.IsValidateFieldIrrigationTypeEnum(requestIrrigationType.Name) {
		return fmt.Errorf("o tipo de irrigacao %w", myerror.ErrEnumInvalid)
	}

	if err := it.irrigationTypeRepository.CreateIrrigationType(entityirrigationType); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (it *IrrigationTypeService) GetIrrigationTypeByID(id uint) (*responses.IrrigationTypeResponse, error) {

	irrigationTypeEntity, err := it.irrigationTypeRepository.FindIrrigationTypetByID(id)
	if err != nil {
		return nil, fmt.Errorf("erro: %w", err)
	}

	irrigationTypeResponse := responses.IrrigationTypeResponse{
		ID:   irrigationTypeEntity.ID,
		Name: irrigationTypeEntity.Name,
	}

	return &irrigationTypeResponse, nil
}

func (it *IrrigationTypeService) PutIrrigationType(id uint, requestIrrigationType requests.IrrigationTypeRequest) error {

	entityirrigationType := entities.IrrigationTypeEntity{
		Name: requestIrrigationType.Name,
	}

	if err := it.irrigationTypeRepository.UpdateIrrigationType(id, entityirrigationType); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}

func (it *IrrigationTypeService) DeleteIrrigationType(id uint) error {

	if err := it.irrigationTypeRepository.DeleteIrrigationType(id); err != nil {
		return fmt.Errorf("erro: %w", err)
	}

	return nil
}
