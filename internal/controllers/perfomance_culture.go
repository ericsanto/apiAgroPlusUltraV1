package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type PerformancePlantingControllerInterface interface {
	PostPerformanceCulture(c *gin.Context)
	GetAllPerformancePlanting(c *gin.Context)
	PutPerformancePlanting(c *gin.Context)
	GetPerformancePlantingByID(c *gin.Context)
	DeletePerformancePlanting(c *gin.Context)
}

type PerformancePlantingController struct {
	servvicePerformancePlanting services.PerformancePlantingServiceInterface
}

func NewPerformancePlantingController(servvicePerformancePlanting services.PerformancePlantingServiceInterface) PerformancePlantingControllerInterface {
	return &PerformancePlantingController{servvicePerformancePlanting: servvicePerformancePlanting}
}

func (p *PerformancePlantingController) PostPerformanceCulture(c *gin.Context) {

	var requestPerformanceCulture requests.PerformancePlantingRequest

	if err := c.ShouldBindJSON(&requestPerformanceCulture); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestPerformanceCulture)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := p.servvicePerformancePlanting.PostPerformancePlanting(requestPerformanceCulture); err != nil {
		switch {
		case errors.Is(err, myerror.ErrDuplicateKey):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrViolatedForeingKey):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return

		case errors.Is(err, myerror.ErrEnumInvalid):
			myerror.HttpErrors(http.StatusUnprocessableEntity, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			fmt.Println(err.Error())
			return
		}

	}

	c.Status(http.StatusCreated)
}

func (p *PerformancePlantingController) GetAllPerformancePlanting(c *gin.Context) {

	performancePlanting, err := p.servvicePerformancePlanting.GetAllPerformancePlanting()
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, performancePlanting)
}

func (p *PerformancePlantingController) PutPerformancePlanting(c *gin.Context) {

	var requestPerformancePlanting requests.PerformancePlantingRequest

	if err := c.ShouldBindJSON(&requestPerformancePlanting); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestPerformancePlanting)

	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := p.servvicePerformancePlanting.PutPerformancePlanting(id, requestPerformancePlanting); err != nil {
		switch {
		case errors.Is(err, myerror.ErrDuplicateKey):
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
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

func (p *PerformancePlantingController) GetPerformancePlantingByID(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	responsePerformancePlanting, err := p.servvicePerformancePlanting.GetPerformancePlantingWithAgricultureCultureAndPlantingEntitiesByI(id)
	if err != nil {
		switch {
		case errors.Is(err, myerror.ErrNotFound):
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return

		default:
			myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
			return
		}

	}

	c.JSON(http.StatusOK, responsePerformancePlanting)
}

func (p *PerformancePlantingController) DeletePerformancePlanting(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "validatedID")

	if err := p.servvicePerformancePlanting.DeletePerformancePlanting(id); err != nil {
		switch {
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
