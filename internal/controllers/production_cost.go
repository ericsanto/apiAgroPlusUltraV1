package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
)

type ProductionCostController struct {
	productionCostService *services.ProductionCostService
}

func NewProductionCostController(productionCostService *services.ProductionCostService) *ProductionCostController {
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

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

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

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

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

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

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
