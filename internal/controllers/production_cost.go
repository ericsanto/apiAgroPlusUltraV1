package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type ProductionCostControllerInterface interface {
	GetAllProductionCost(c *gin.Context)
	PostProductionCost(c *gin.Context)
	GetProductionCostByID(c *gin.Context)
	PutProductionCost(c *gin.Context)
	DeleteProductionCost(c *gin.Context)
}

type ProductionCostController struct {
	productionCostService services.ProductionCostServiceInterface
}

func NewProductionCostController(productionCostService services.ProductionCostServiceInterface) ProductionCostControllerInterface {
	return &ProductionCostController{productionCostService: productionCostService}
}

func (p *ProductionCostController) GetAllProductionCost(c *gin.Context) {

	productionCost, err := p.productionCostService.GetAllProductionCost()
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, productionCost)
}

func (p *ProductionCostController) PostProductionCost(c *gin.Context) {

	var requesProductionCost requests.ProductionCostRequest

	if err := c.ShouldBindJSON(&requesProductionCost); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	valid, err := validators.ValidateFieldErrors422UnprocessableEntity(requesProductionCost)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(valid) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, valid, c)
		return
	}

	if err := p.productionCostService.PostProductionCost(requesProductionCost); err != nil {
		if strings.Contains(err.Error(), "plantio com o ID fornecido não existe") {
			myerror.HttpErrors(http.StatusUnprocessableEntity, fmt.Sprintf("plantio com o ID %d não existe", requesProductionCost.PlantingID), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusCreated)
}

func (p *ProductionCostController) GetProductionCostByID(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "id")

	productionCost, err := p.productionCostService.GetAllProductionCostByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "não existe custo com o id") {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe custo com o id %d", id), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, productionCost)
}

func (p *ProductionCostController) PutProductionCost(c *gin.Context) {

	var requesProductionCost requests.ProductionCostRequest

	id := validators.GetAndValidateIdMidlware(c, "id")

	if err := c.ShouldBindJSON(&requesProductionCost); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	valid, err := validators.ValidateFieldErrors422UnprocessableEntity(requesProductionCost)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(valid) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, valid, c)
		return
	}

	if err := p.productionCostService.PutProductionCost(id, requesProductionCost); err != nil {
		if strings.Contains(err.Error(), "plantio com o ID fornecido não existe") {
			myerror.HttpErrors(http.StatusUnprocessableEntity, fmt.Sprintf("plantio com o ID %d não existe", requesProductionCost.PlantingID), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusOK)

}

func (p *ProductionCostController) DeleteProductionCost(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "id")

	if err := p.productionCostService.DeleteProductionCost(id); err != nil {
		if strings.Contains(err.Error(), "não existe custo com o id") {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe custo com o id %d", id), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusNoContent)
}
