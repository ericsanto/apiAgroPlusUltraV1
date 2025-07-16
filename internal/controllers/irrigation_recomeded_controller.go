package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type IrrigationRecomendedController struct {
	irrigationRecomendedService *services.IrrigationRecomendedService
}

func NewIrrigationRecomendedController(irrigationRecomendedService *services.IrrigationRecomendedService) *IrrigationRecomendedController {
	return &IrrigationRecomendedController{irrigationRecomendedService: irrigationRecomendedService}
}

func (i *IrrigationRecomendedController) GetAllIrrigationRecomended(c *gin.Context) {

	responseIrrigationRecomended, err := i.irrigationRecomendedService.GetAllIrrigationRecomended()
	if err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "erro no servidor",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, responseIrrigationRecomended)
}

func (i *IrrigationRecomendedController) PostIrrigationRecomended(c *gin.Context) {

	var requestIrrigationRecomended requests.IrrigationRecomendedRequest

	if err := c.ShouldBindBodyWithJSON(&requestIrrigationRecomended); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "body da requisição é inválido",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestIrrigationRecomended)
	if err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "erro interno",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return

	}

	if len(validate) > 0 {
		c.JSON(http.StatusUnprocessableEntity, myerror.ErrorApp{
			Code:      http.StatusUnprocessableEntity,
			Message:   validate,
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	if err := i.irrigationRecomendedService.PostIrrigationRecomended(requestIrrigationRecomended); err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "erro interno",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusCreated)

}

func (i *IrrigationRecomendedController) GetByIdrrigationRecomended(c *gin.Context) {

	val, exists := c.Get("id")
	if !exists {
		return
	}

	id := val.(uint)

	responseIrrigationRecomended, err := i.irrigationRecomendedService.GetByIdIrrigationRecomended(id)
	if err != nil {
		if strings.Contains(err.Error(), "objeto com o id fornecido não existe") {
			c.JSON(http.StatusNotFound, myerror.ErrorApp{
				Code:      http.StatusNotFound,
				Message:   "objeto com o id fornecido não existe",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "erro no servidor",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, responseIrrigationRecomended)
}

func (i *IrrigationRecomendedController) PutIrrigationRecomendedController(c *gin.Context) {

	val, exists := c.Get("id")
	if !exists {
		return
	}

	id := val.(uint)

	var requestIrrigationRecomended requests.IrrigationRecomendedRequest

	if err := c.ShouldBindJSON(&requestIrrigationRecomended); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "body da requisição é inválida",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestIrrigationRecomended)
	if err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "erro no servidor",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	if len(validate) > 0 {
		c.JSON(http.StatusUnprocessableEntity, myerror.ErrorApp{
			Code:      http.StatusUnprocessableEntity,
			Message:   validate,
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	if err := i.irrigationRecomendedService.PutIrrigationRecomended(id, requestIrrigationRecomended); err != nil {
		if strings.Contains(err.Error(), "objeto com o id fornecido não existe") {
			c.JSON(http.StatusNotFound, myerror.ErrorApp{
				Code:      http.StatusNotFound,
				Message:   "objeto com o id fornecido não existe",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "erro no servidor",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (i *IrrigationRecomendedController) DeleteIrrigationRecomendedController(c *gin.Context) {

	val, exists := c.Get("id")
	if !exists {
		return
	}

	id := val.(uint)

	if err := i.irrigationRecomendedService.DeleteIrrigationRecomended(id); err != nil {
		if strings.Contains(err.Error(), "objeto com o id fornecido não existe") {
			c.JSON(http.StatusNotFound, myerror.ErrorApp{
				Code:      http.StatusNotFound,
				Message:   "objeto com o id fornecido não existe",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "erro no servidor",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
