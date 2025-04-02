package repositories

import (
	"fmt"
	"log"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/entities"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"gorm.io/gorm"
)

type PestRepositoryInterface interface {
	FindAllPestService() ([]entities.PestEntity, error)
	FindByIdPestService(id uint) (entities.TypePestEntity, error)
	CreatePestService(requestPest requests.PestRequest) error
	UpdatePestService(id uint, requesPest requests.PestRequest) error
	DeletePestService(id uint) error
}

type PestRepository struct {
	db *gorm.DB
}

func NewPestRepository(db *gorm.DB) *PestRepository {
	return &PestRepository{db: db}
}

func (p *PestRepository) FindAllPest() ([]entities.PestEntity, error) {

	var entitiesPest []entities.PestEntity

	err := p.db.Find(&entitiesPest)
	if err.Error != nil {
		return entitiesPest, fmt.Errorf("erro ao buscar todas as pragas")
	}

	return entitiesPest, nil
}

func (p *PestRepository) FindByIdPest(id uint) (entities.PestEntity, error) {

	var pestEntity entities.PestEntity

	if err := p.db.First(&pestEntity, id); err.Error != nil {
		fmt.Println(err)
		return pestEntity, fmt.Errorf("erro ao buscar praga com o id fornecido")
	}

	return pestEntity, nil

}

func (p *PestRepository) CreatePest(entityPest entities.PestEntity) error {

	if err := p.db.Create(&entityPest); err.Error != nil {
		log.Print(err)
		return fmt.Errorf("erro ao criar praga: %v", err)
	}

	return nil
}

func (p *PestRepository) UpdatePest(id uint, entityPest entities.PestEntity) error {

	existsPestEntity, err := p.FindByIdPest(id)
	if err != nil {
		return fmt.Errorf("erro ao atualizar praga. Id não existe")
	}

	if err := p.db.Model(&entities.PestEntity{}).Where("id = ?", existsPestEntity.Id).Updates(entityPest); err.Error != nil {
		return fmt.Errorf("praga encontrada no banco de dados. Porém, não foi possível atualizá-la %v", err)
	}

	return nil
}

func (p *PestRepository) DeletePest(id uint) error {

	existsPestEntity, err := p.FindByIdPest(id)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	if err := p.db.Delete(&existsPestEntity); err.Error != nil {
		return fmt.Errorf("praga encontrada no banco de dados. Porém, não foi possível atualizá-la: %v", err)
	}

	return nil
}
