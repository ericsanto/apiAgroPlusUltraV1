package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type SalePlantingControllerInterface interface {
	PostSalePlanting(c *gin.Context)
	GetAllSalePlanting(c *gin.Context)
	GetSalePlantingByID(c *gin.Context)
	PutSalePlanting(c *gin.Context)
	DeleteSalePlanting(c *gin.Context)
}

type SalePlantingController struct {
	salePlantingService services.SalePlantingServiceInterface
}

func NewSalePlantingController(salePlantingService services.SalePlantingServiceInterface) SalePlantingControllerInterface {
	return &SalePlantingController{salePlantingService: salePlantingService}
}

func (s *SalePlantingController) PostSalePlanting(c *gin.Context) {

	var requestSalePlanting requests.SalePlantingRequest

	if err := c.ShouldBindJSON(&requestSalePlanting); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestSalePlanting)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := s.salePlantingService.PostSalePlanting(requestSalePlanting); err != nil {
		if errors.Is(err, myerror.ErrDuplicateSale) {
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusOK)

}

func (s *SalePlantingController) GetAllSalePlanting(c *gin.Context) {

	responsesSalePlanting, err := s.salePlantingService.GetAllSalePlanting()
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, responsesSalePlanting)
}

func (s *SalePlantingController) GetSalePlantingByID(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "id")

	responseSalePlanting, err := s.salePlantingService.GetSalePlantingByID(id)
	if err != nil {
		if errors.Is(err, myerror.ErrNotFoundSale) {
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, responseSalePlanting)
}

func (s *SalePlantingController) PutSalePlanting(c *gin.Context) {

	var requestSalePlanting requests.SalePlantingRequest

	if err := c.ShouldBindJSON(&requestSalePlanting); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestSalePlanting)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	id := validators.GetAndValidateIdMidlware(c, "id")

	if err := s.salePlantingService.PutSalePlanting(id, requestSalePlanting); err != nil {

		switch {
		case errors.Is(err, myerror.ErrDuplicateSale):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrNotFoundSale):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrViolatedForeingKey):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return
		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}

	}

	c.Status(http.StatusOK)
}

func (s *SalePlantingController) DeleteSalePlanting(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "id")

	if err := s.salePlantingService.DeleteSalePlanting(id); err != nil {
		switch {
		case errors.Is(err, myerror.ErrNotFoundSale):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}
	}

	c.Status(http.StatusNoContent)
}
