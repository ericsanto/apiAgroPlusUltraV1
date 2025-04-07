package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
)

type AgricultureCultureIrrigationController struct {
	agricultureCultureIrrigationService *services.AgricultureCultureIrrigationService
}

func NewAgricultureCultureIrrigationController(agricultureCultureIrrigationService *services.AgricultureCultureIrrigationService) *AgricultureCultureIrrigationController {
	return &AgricultureCultureIrrigationController{agricultureCultureIrrigationService: agricultureCultureIrrigationService}
}

func (a *AgricultureCultureIrrigationController) GetAgricultureCultureIrrigationFindByIDController(c *gin.Context) {

	val, exists := c.Get("validatedCultureId")
	if !exists {
		return
	}

	culture_id := val.(uint)

	agriculturesCulturesIrrigationsResponse, err := a.agricultureCultureIrrigationService.GetAgricultureCultureIrrigationFindByID(culture_id)
	if err != nil {

		if strings.Contains(err.Error(), "erro ao buscar cultura agrícola. Id não existe no banco de dados") {
			c.JSON(http.StatusNotFound, myerror.ErrorApp{
				Code:      http.StatusNotFound,
				Message:   "erro ao buscar cultura agrícola. Id não existe no banco de dados",
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

	c.JSON(http.StatusOK, agriculturesCulturesIrrigationsResponse)
}

func (a *AgricultureCultureIrrigationController) PostAgricultureCultureIrrigationController(c *gin.Context) {

	var requestAgricultureCultureIrrigation requests.AgricultureCultureIrrigationRequest

	if err := c.ShouldBindJSON(&requestAgricultureCultureIrrigation); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "body da requisição é inválido",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestAgricultureCultureIrrigation)
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

	if err := a.agricultureCultureIrrigationService.PostAgricultureCultureIrrigation(requestAgricultureCultureIrrigation); err != nil {
		if strings.Contains(err.Error(), "erro ao tentar cadastrar. já existe objeto com essa relação de id") {
			c.JSON(http.StatusConflict, myerror.ErrorApp{
				Code:      http.StatusConflict,
				Message:   "erro ao tentar cadastrar. já existe objeto com essa relação de id",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		if strings.Contains(err.Error(), "id não existe") {
			c.JSON(http.StatusNotFound, myerror.ErrorApp{
				Code:      http.StatusNotFound,
				Message:   "id não existe",
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

	c.Status(http.StatusCreated)

}

func (a *AgricultureCultureIrrigationController) PutAgricultureCultureIrrigation(c *gin.Context) {

	val, exist := c.Get("validatedCultureId")

	if !exist {
		return
	}

	cultureId := val.(uint)

	val, exist = c.Get("validatedIrrigationId")

	if !exist {
		return
	}

	irrigationId := val.(uint)

	var requestAgricultureCultureIrrigation requests.AgricultureCultureIrrigationRequest

	if err := c.ShouldBindJSON(&requestAgricultureCultureIrrigation); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "body da requisição é inválido",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	valid, err := validators.ValidateFieldErrors422UnprocessableEntity(requestAgricultureCultureIrrigation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "erro no servidor",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	if len(valid) > 0 {
		c.JSON(http.StatusUnprocessableEntity, myerror.ErrorApp{
			Code:      http.StatusUnprocessableEntity,
			Message:   valid,
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	if err := a.agricultureCultureIrrigationService.PutAgricultureCultureIrrigation(cultureId, irrigationId, requestAgricultureCultureIrrigation); err != nil {
		if strings.Contains(err.Error(), "não existe objeto com id fornecido") {
			c.JSON(http.StatusNotFound, myerror.ErrorApp{
				Code:      http.StatusNotFound,
				Message:   "não existe objeto com id fornecido",
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

func (a *AgricultureCultureIrrigationController) DeleteAgricultureCulturueIrrigation(c *gin.Context) {

	val, exist := c.Get("validatedCultureId")

	if !exist {
		return
	}

	cultureId := val.(uint)

	val, exist = c.Get("validatedIrrigationId")

	if !exist {
		return
	}

	irrigationId := val.(uint)

	if err := a.agricultureCultureIrrigationService.DeleteAgricultureCulturueIrrigation(cultureId, irrigationId); err != nil {
		if strings.Contains(err.Error(), "não existe objeto com id fornecido") {
			c.JSON(http.StatusNotFound, myerror.ErrorApp{
				Code:      http.StatusNotFound,
				Message:   "não existe objeto com id fornecido",
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
