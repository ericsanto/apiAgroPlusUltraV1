package controllers

import (
	"errors"
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

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")
	plantingID := validators.GetAndValidateIdMidlware(c, "plantingID")

	productionCost, err := p.productionCostService.GetAllProductionCost(batchID, farmID, userID, plantingID)
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

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")
	plantingID := validators.GetAndValidateIdMidlware(c, "plantingID")

	if err := p.productionCostService.PostProductionCost(batchID, farmID, userID, plantingID, requesProductionCost); err != nil {

		switch {
		case strings.Contains(err.Error(), "plantio com o ID fornecido não existe"):
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("plantio com o ID %d não existe", plantingID), c)
			return

		case errors.Is(err, myerror.ErrFarmNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}
	}

	c.Status(http.StatusCreated)
}

func (p *ProductionCostController) GetProductionCostByID(c *gin.Context) {

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")
	plantingID := validators.GetAndValidateIdMidlware(c, "plantingID")
	costID := validators.GetAndValidateIdMidlware(c, "costID")

	productionCost, err := p.productionCostService.GetAllProductionCostByID(batchID, farmID, userID, plantingID, costID)
	if err != nil {
		if strings.Contains(err.Error(), "não existe custo com o id") {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe custo com o id %d", costID), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, productionCost)
}

func (p *ProductionCostController) PutProductionCost(c *gin.Context) {

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

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")
	plantingID := validators.GetAndValidateIdMidlware(c, "plantingID")
	costID := validators.GetAndValidateIdMidlware(c, "costID")

	if err := p.productionCostService.PutProductionCost(batchID, farmID, userID, plantingID, costID, requesProductionCost); err != nil {

		switch {

		case strings.Contains(err.Error(), "não existe custo com o id"):
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe custo com o id %d", costID), c)
			return

		case strings.Contains(err.Error(), "plantio com o ID fornecido não existe"):
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("plantio com o ID %d não existe", plantingID), c)
			return

		case errors.Is(err, myerror.ErrFarmNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}

	}

	c.Status(http.StatusOK)

}

func (p *ProductionCostController) DeleteProductionCost(c *gin.Context) {

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")
	plantingID := validators.GetAndValidateIdMidlware(c, "plantingID")
	costID := validators.GetAndValidateIdMidlware(c, "costID")

	if err := p.productionCostService.DeleteProductionCost(batchID, farmID, userID, plantingID, costID); err != nil {

		switch {

		case strings.Contains(err.Error(), "não existe custo com o id"):
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe custo com o id %d", costID), c)
			return

		case strings.Contains(err.Error(), "plantio com o ID fornecido não existe"):
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("plantio com o ID %d não existe", plantingID), c)
			return

		case errors.Is(err, myerror.ErrFarmNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}

	}

	c.Status(http.StatusNoContent)
}
