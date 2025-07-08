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

type SustainablePestControlControllerInterface interface {
	GetAllSustainablePestControl(c *gin.Context)
	PostSustainablePestControl(c *gin.Context)
	GetFindByIdSustainablePestControl(c *gin.Context)
	PutSustainablePestControl(c *gin.Context)
	DeleteSustainablePestControl(c *gin.Context)
}

type SustainablePestControlController struct {
	sustainablePestControlService services.SustainablePestControlServiceInterface
}

func NewSustainablePestControlController(sustainablePestControlService services.SustainablePestControlServiceInterface) SustainablePestControlControllerInterface {
	return &SustainablePestControlController{sustainablePestControlService: sustainablePestControlService}
}

func (s *SustainablePestControlController) GetAllSustainablePestControl(c *gin.Context) {

	responseSustainablesPestControl, err := s.sustainablePestControlService.GetAllSustainablePestControl()
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, responseSustainablesPestControl)
}

func (s *SustainablePestControlController) PostSustainablePestControl(c *gin.Context) {

	var requestSustainablePestControl requests.SustainablePestControlRequest

	if err := c.ShouldBindJSON(&requestSustainablePestControl); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestSustainablePestControl)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := s.sustainablePestControlService.PostSustainablePestControl(requestSustainablePestControl); err != nil {
		if strings.Contains(err.Error(), "objeto já existe com esse nome") {
			myerror.HttpErrors(http.StatusConflict, err.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusCreated)

}

func (s *SustainablePestControlController) GetFindByIdSustainablePestControl(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "id")

	responseSustainablesPestControl, err := s.sustainablePestControlService.GetFindByIdSustainablePestControl(id)
	if err != nil {
		if strings.Contains(err.Error(), "record not found") {
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.JSON(http.StatusOK, responseSustainablesPestControl)
}

func (s *SustainablePestControlController) PutSustainablePestControl(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "id")

	var requestSustainablePestControl requests.SustainablePestControlRequest

	if err := c.ShouldBindJSON(&requestSustainablePestControl); err != nil {
		myerror.HttpErrors(http.StatusBadRequest, "body da requisição é inválido", c)
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(requestSustainablePestControl)
	if err != nil {
		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	if len(validate) > 0 {
		myerror.HttpErrors(http.StatusUnprocessableEntity, validate, c)
		return
	}

	if err := s.sustainablePestControlService.PutSustainablePestControl(id, requestSustainablePestControl); err != nil {
		if strings.Contains(err.Error(), "record not found") {
			myerror.HttpErrors(http.StatusNotFound, err.Error(), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusOK)

}

func (s *SustainablePestControlController) DeleteSustainablePestControl(c *gin.Context) {

	id := validators.GetAndValidateIdMidlware(c, "id")

	if err := s.sustainablePestControlService.DeleteSustainablePestControl(id); err != nil {
		if strings.Contains(err.Error(), fmt.Sprintf("obejto com id %d não existe", id)) {
			myerror.HttpErrors(http.StatusNotFound, fmt.Sprintf("obejto com id %d não existe", id), c)
			return
		}

		myerror.HttpErrors(http.StatusInternalServerError, "erro no servidor", c)
		return
	}

	c.Status(http.StatusNoContent)
}
