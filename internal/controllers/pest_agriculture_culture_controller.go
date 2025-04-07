package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
)

type PestAgricultureCultureController struct {
	pestAgricultureCultureService *services.PestAgricultureCultureService
}

func NewPestAgricultureCultureController(pestAgricultureCultureService *services.PestAgricultureCultureService) *PestAgricultureCultureController {
	return &PestAgricultureCultureController{pestAgricultureCultureService: pestAgricultureCultureService}
}

func (p *PestAgricultureCultureController) GetAllAgricultureCultureController(c *gin.Context) {

	rows, err := p.pestAgricultureCultureService.GetAllPestAgricultureCulture()
	if err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, rows)
}

func (p *PestAgricultureCultureController) GetFindByIdAgricultureCultureController(c *gin.Context) {

	pestId, ok := c.Get("validatedPestId")
	if !ok {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "pestId não validado",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	cultureId, ok := c.Get("validatedCultureId")
	if !ok {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "cultureId não validado",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	result, err := p.pestAgricultureCultureService.GetFindByIdPestAgricultureCulture(
		pestId.(uint),
		cultureId.(uint),
	)
	if err != nil {
		c.JSON(http.StatusNotFound, myerror.ErrorApp{
			Code:      http.StatusNotFound,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (p *PestAgricultureCultureController) PostPestAgricultureCultureController(c *gin.Context) {

	var requestPestAgricultureCulture requests.PestAgricultureCultureRequest

	if err := c.ShouldBindJSON(&requestPestAgricultureCulture); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "Body da requisição é inválido",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	valid, err := validators.ValidateFieldErrors422UnprocessableEntity(requestPestAgricultureCulture)
	if err != nil {
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

	if err := p.pestAgricultureCultureService.PostPestAgricultureCulture(requestPestAgricultureCulture); err != nil {
		if strings.Contains(err.Error(), "já estão relacionados") {
			c.JSON(http.StatusConflict, myerror.ErrorApp{
				Code:      http.StatusConflict,
				Message:   err.Error(),
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "Erro no servidor",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (p *PestAgricultureCultureController) PutPestAgricultureCulture(c *gin.Context) {

	val, exists := c.Get("validatedPestId")

	if !exists {
		return
	}

	pestId := val.(uint)

	val, exists = c.Get("validatedCultureId")

	if !exists {
		return
	}

	cultureId := val.(uint)

	var requestPestAgricultureCulture requests.PestAgricultureCultureRequest

	if err := c.ShouldBindJSON(&requestPestAgricultureCulture); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "body da requisição é inválido",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	if err := p.pestAgricultureCultureService.PutPestAgricultureCulture(pestId, cultureId, requestPestAgricultureCulture); err != nil {
		if strings.Contains(err.Error(), "objeto com id fornecido não existe") {
			c.JSON(http.StatusNotFound, myerror.ErrorApp{
				Code:      http.StatusNotFound,
				Message:   fmt.Sprintf("objeto com pestId=%d e cultureId=%d não existe", pestId, cultureId),
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		c.Status(http.StatusInternalServerError)
		return
	}

	valid, err := validators.ValidateFieldErrors422UnprocessableEntity(requestPestAgricultureCulture)

	if err != nil {
		log.Fatal(err)
	}

	if len(valid) > 0 {
		c.JSON(http.StatusUnprocessableEntity, myerror.ErrorApp{
			Code:      http.StatusUnprocessableEntity,
			Message:   valid,
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusOK)

}

func (p *PestAgricultureCultureController) DeletePestAgricultureCultureController(c *gin.Context) {

	val, exists := c.Get("validatedPestId")

	if !exists {
		return
	}

	pestId := val.(uint)

	val, exists = c.Get("validatedCultureId")

	if !exists {
		return
	}

	cultureId := val.(uint)

	if err := p.pestAgricultureCultureService.DeletePestAgricultureCulture(pestId, cultureId); err != nil {
		fmt.Println(err.Error())
		if strings.Contains(err.Error(), "não existe objeto com o id fornecido") {
			c.JSON(http.StatusNotFound, myerror.ErrorApp{
				Code:      http.StatusNotFound,
				Message:   "objeto com o id fornecido não existe",
				Timestamp: time.Now().Format(time.RFC3339),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   fmt.Sprintf("erro interno %s", err),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
