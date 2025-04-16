package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BatchController struct {
	batchService *services.BatchService
}

func NewBatchController(batchService *services.BatchService) *BatchController {
	return &BatchController{batchService: batchService}
}

func (b *BatchController) PostBatch(c *gin.Context) {
	var requestBatch requests.BatchRequest

	if err := c.ShouldBindJSON(&requestBatch); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestBatch)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := b.batchService.PostBatchService(requestBatch); err != nil {
		if strings.Contains(err.Error(), "já existe lote cadastrado com esse nome") {
			myerror.HttpErrors(http.StatusConflict, "já existe lote cadastrado com esse nome", c)
			return
		}
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusCreated)

}

func (b *BatchController) GetAllBatch(c *gin.Context) {

	batchs, err := b.batchService.GetAllBatch()
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, batchs)
}

func (b *BatchController) GetBatchFindById(c *gin.Context) {

	validateID := validators.GetAndValidateIdMidlware(c, "validatedID")

	batch, err := b.batchService.GetBatchFindById(validateID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe lote com o ID %d", validateID), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, batch)
}

func (b *BatchController) PutBatch(c *gin.Context) {

	var requestBatch requests.BatchRequest

	if err := c.ShouldBind(&requestBatch); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestBatch)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	validatedID := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := b.batchService.PutBatch(validatedID, requestBatch); err != nil {
		if strings.Contains(err.Error(), "já existe lote cadastrado com esse nome") {
			myerror.HttpErrors(http.StatusConflict, "já existe lote cadastrado com esse nome", c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusOK)
}

func (b *BatchController) DeleteBatch(c *gin.Context) {

	validatedId := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := b.batchService.DeleteBatch(validatedId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe lote com o ID %d", validatedId), c)
			return
		}
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusNoContent)
}
