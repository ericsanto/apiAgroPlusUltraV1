package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
)

type AgricultureCultureControllerInterface interface {
	GetAllAgriculturesCulturesController(c *gin.Context)
	PostAgricultureCultureController(c *gin.Context)
	PutAgricultureCultureController(c *gin.Context)
	DeleteAgricultureCultureController(c *gin.Context)
}

type AgricultureCultureController struct {
	agricultureCultureService services.AgricultureCultureServiceInterface
}

func NewAgricultureController(agricultureCultureService services.AgricultureCultureServiceInterface) AgricultureCultureControllerInterface {

	return &AgricultureCultureController{agricultureCultureService: agricultureCultureService}
}

func (a *AgricultureCultureController) GetAllAgriculturesCulturesController(c *gin.Context) {

	agricultureCultures, err := a.agricultureCultureService.FindAllAgricultureCultureService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "Internal server error",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		log.Println(err.Error())
		return
	}
	c.JSON(http.StatusOK, agricultureCultures)
}

func (a *AgricultureCultureController) PostAgricultureCultureController(c *gin.Context) {

	var agricultureCultureRequest requests.AgricultureCultureRequest

	if err := c.ShouldBindJSON(&agricultureCultureRequest); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "Invalid request body",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		log.Println(err.Error())
		return
	}

	valid, err := validators.ValidateFieldErrors422UnprocessableEntity(agricultureCultureRequest)

	if err != nil {
		log.Println(err.Error())
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

	if err := a.agricultureCultureService.CreateAgricultureCultureService(agricultureCultureRequest); err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "Internal server error",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusCreated)

}

func (a *AgricultureCultureController) PutAgricultureCultureController(c *gin.Context) {

	val, exists := c.Get("id")
	if !exists {
		return
	}

	id := val.(uint)

	var agricultureCulture requests.AgricultureCultureRequest

	if err := c.ShouldBindJSON(&agricultureCulture); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "Invalid request body",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		log.Println(err.Error())
		return
	}

	valid, err := validators.ValidateFieldErrors422UnprocessableEntity(agricultureCulture)
	if err != nil {
		log.Println(err.Error())
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

	if err := a.agricultureCultureService.PutAgricultureCultureService(id, agricultureCulture); err != nil {
		c.JSON(http.StatusNotFound, myerror.ErrorApp{
			Code:      http.StatusNotFound,
			Message:   "Id not found",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (a *AgricultureCultureController) DeleteAgricultureCultureController(c *gin.Context) {

	val, exists := c.Get("id")
	if !exists {
		return
	}

	id := val.(uint)

	if err := a.agricultureCultureService.DeleteAgricultureCultureService(id); err != nil {
		c.JSON(http.StatusNotFound, myerror.ErrorApp{
			Code:      http.StatusNotFound,
			Message:   "id not found",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
