package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
)

type SoilTypeControollerInterface interface {
	GetAllSoilTypes(c *gin.Context)
	GetSoilTypeFindById(c *gin.Context)
	PostSoilType(c *gin.Context)
	PutSoilType(c *gin.Context)
	DeleteSoilType(c *gin.Context)
}

type SoilTypeController struct {
	typeSoilService services.SoilTypeServiceInterface
}

func NewSoilTypeController(typeSoilService services.SoilTypeServiceInterface) SoilTypeControollerInterface {
	return &SoilTypeController{typeSoilService: typeSoilService}
}

func (s *SoilTypeController) GetAllSoilTypes(c *gin.Context) {

	service, err := s.typeSoilService.GetAllSoilType()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Erro": err})
		return
	}

	c.JSON(http.StatusOK, service)

}

func (s *SoilTypeController) GetSoilTypeFindById(c *gin.Context) {

	strIdParam := c.Param("id")
	intIdParam, err := strconv.Atoi(strIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "O id deve ser um número inteiro",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	service, err := s.typeSoilService.GetSoilTypeFindById(uint(intIdParam))

	if err != nil {
		c.JSON(http.StatusNotFound, myerror.ErrorApp{
			Code:      http.StatusNotFound,
			Message:   "Não existe tipo de solo com o id fornecido",
			Timestamp: time.Now().String(),
		})
		return
	}

	c.JSON(http.StatusOK, service)

}

func (s *SoilTypeController) PostSoilType(c *gin.Context) {

	var soilType requests.SoilTypeRequest

	if err := c.ShouldBindJSON(&soilType); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   err.Error(),
			Timestamp: time.Now().String(),
		})
		return
	}

	if err := s.typeSoilService.PostSoilType(soilType); err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			Timestamp: time.Now().String(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (s *SoilTypeController) PutSoilType(c *gin.Context) {

	var soilType requests.SoilTypeRequest

	val, exists := c.Get("validatedID")
	if !exists {
		return
	}

	id := val.(uint)

	if err := c.ShouldBindJSON(&soilType); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   err.Error(),
			Timestamp: time.Now().String(),
		})
		return
	}

	if err := s.typeSoilService.PutSoilType(id, soilType); err != nil {
		c.JSON(http.StatusNotFound, myerror.ErrorApp{
			Code:      http.StatusNotFound,
			Message:   err.Error(),
			Timestamp: time.Now().String(),
		})
		return
	}

	c.Status(http.StatusOK)

}

func (s *SoilTypeController) DeleteSoilType(c *gin.Context) {

	val, exists := c.Get("validatedID")

	if !exists {
		return
	}

	id := val.(uint)

	if err := s.typeSoilService.DeleteTypeSoil(id); err != nil {
		c.JSON(http.StatusNotFound, myerror.ErrorApp{
			Message:   err.Error(),
			Code:      http.StatusNotFound,
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusNoContent)

}
