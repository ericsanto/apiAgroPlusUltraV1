package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/validators"
	"github.com/gin-gonic/gin"
)

type TypePestController struct {
	typePestService *services.TypePestService
}

func NewTypePestController(typePestService *services.TypePestService) *TypePestController {
	return &TypePestController{typePestService: typePestService}
}

func (t *TypePestController) GetAllTypePestController(c *gin.Context) {

	pests, err := t.typePestService.GetAllTypePest()
	if err != nil {
		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   "Não foi possível buscar pragas",
			Timestamp: time.Now().Format(time.RFC3339),
		})
	}

	c.JSON(http.StatusOK, pests)
}

func (t *TypePestController) GetAllTypePestFindByIdController(c *gin.Context) {

	val, exists := c.Get("validatedID")
	if !exists {
		return
	}

	id := val.(uint)

	pest, err := t.typePestService.GetTypePestFindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, myerror.ErrorApp{
			Code:      http.StatusNotFound,
			Message:   "Não existe praga com esse id",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.JSON(http.StatusOK, pest)
}

func (t *TypePestController) PostTypePestController(c *gin.Context) {

	var pestRequest requests.TypePestRequest

	if err := c.ShouldBindJSON(&pestRequest); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(pestRequest)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, myerror.ErrorApp{
			Code:      http.StatusUnprocessableEntity,
			Message:   validate,
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	if err := t.typePestService.PostTypePest(pestRequest); err != nil {

		c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func (t *TypePestController) PutTypePestController(c *gin.Context) {

	val, exists := c.Get("validatedID")
	if !exists {
		return
	}

	id := val.(uint)

	var pestRequest requests.TypePestRequest

	if err := c.ShouldBindJSON(&pestRequest); err != nil {
		c.JSON(http.StatusBadRequest, myerror.ErrorApp{
			Code:      http.StatusBadRequest,
			Message:   "Body da requisição inválido!",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	validate, err := validators.ValidateFieldErrors422UnprocessableEntity(pestRequest)
	if err != nil {
		log.Println(err.Error())
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

	if err := t.typePestService.PutTypePest(id, pestRequest); err != nil {
		c.JSON(http.StatusNotFound, myerror.ErrorApp{
			Code:      http.StatusNotFound,
			Message:   "Não existe praga com esse id",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusOK)

}

func (t *TypePestController) DeleteTypePestController(c *gin.Context) {

	val, exists := c.Get("validatedID")

	if !exists {
		return
	}

	id := val.(uint)

	if err := t.typePestService.DeleteTypePest(id); err != nil {
		c.JSON(http.StatusNotFound, myerror.ErrorApp{
			Code:      http.StatusNotFound,
			Message:   "Não existe praga com esse id",
			Timestamp: time.Now().Format(time.RFC3339),
		})
		return
	}

	c.Status(http.StatusNoContent)

}
