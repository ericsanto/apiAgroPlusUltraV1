package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type BatchControllerInterface interface {
	PostBatch(c *gin.Context)
	GetBatchFindById(c *gin.Context)
	GetAllBatch(c *gin.Context)
	PutBatch(c *gin.Context)
	DeleteBatch(c *gin.Context)
}

type BatchController struct {
	batchService services.BatchServiceInterface
}

func NewBatchController(batchService services.BatchServiceInterface) BatchControllerInterface {
	return &BatchController{batchService: batchService}
}

func (b *BatchController) PostBatch(c *gin.Context) {
	var requestBatch requests.BatchRequest

	val, exitsts := c.Get("userID")

	if !exitsts {
		return
	}

	farmID := validators.GetAndValidateIdMidlware(c, "farmID")

	userID := val.(uint)

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

	if err := b.batchService.PostBatchService(userID, farmID, requestBatch); err != nil {
		if errors.Is(err, myerror.ErrBatchAlreadyExists) {
			myerror.HttpErrors(http.StatusConflict, myerror.ErrBatchAlreadyExists.Error(), c)
			return
		}
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusCreated)

}

func (b *BatchController) GetAllBatch(c *gin.Context) {

	val, exitsts := c.Get("userID")

	if !exitsts {
		return
	}

	farmID := validators.GetAndValidateIdMidlware(c, "farmID")

	userID := val.(uint)

	batchs, err := b.batchService.GetAllBatch(userID, farmID)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, batchs)
}

func (b *BatchController) GetBatchFindById(c *gin.Context) {

	val, exist := c.Get("userID")

	if !exist {
		return
	}

	userID := val.(uint)

	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")

	batch, err := b.batchService.GetBatchFindById(userID, farmID, batchID)
	if err != nil {
		if errors.Is(err, myerror.ErrFarmNotFound) {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe fazenda com o ID %d", farmID), c)
			return
		}

		if errors.Is(err, myerror.ErrNotFound) {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe lote com o ID %d", batchID), c)
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

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")

	if err := b.batchService.PutBatch(userID, farmID, batchID, requestBatch); err != nil {
		if errors.Is(err, myerror.ErrBatchAlreadyExists) {
			myerror.HttpErrors(http.StatusConflict, myerror.ErrBatchAlreadyExists.Error(), c)
			return
		}

		if errors.Is(err, myerror.ErrFarmNotFound) {
			myerror.HttpErrors(http.StatusNotFound, myerror.ErrFarmNotFound.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusOK)
}

func (b *BatchController) DeleteBatch(c *gin.Context) {

	userID := validators.GetAndValidateIdMidlware(c, "userID")
	farmID := validators.GetAndValidateIdMidlware(c, "farmID")
	batchID := validators.GetAndValidateIdMidlware(c, "batchID")

	if err := b.batchService.DeleteBatch(userID, farmID, batchID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("não existe lote com o ID %d", batchID), c)
			return
		}

		if errors.Is(err, myerror.ErrFarmNotFound) {
			myerror.HttpErrors(http.StatusNotFound, myerror.ErrFarmNotFound.Error(), c)
			return
		}
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusNoContent)
}
