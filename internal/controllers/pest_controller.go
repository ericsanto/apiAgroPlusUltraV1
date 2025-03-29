package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/ericsanto/apiAgroPlusUltraV1/utils"
	"github.com/gin-gonic/gin"
)


type PestController struct {

  pestService *services.PestService
}

func NewPestController(pestService *services.PestService) *PestController {
  return &PestController{pestService:pestService}
}

func(p *PestController) GetAllPestController(c *gin.Context) {

  pestsResponse, err := p.pestService.GetAllPest()
  if err != nil {
    c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
      Code: http.StatusInternalServerError,
      Message: "Não foi possível buscar pragas",
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }

  c.JSON(http.StatusOK, pestsResponse)
}

func(p *PestController) GetFindByIdPestController(c *gin.Context) {

  val, exists := c.Get("validatedID")
  if !exists {
    return
  }

  id := val.(uint)

  pestResponse, err := p.pestService.GetFindByIdPest(id)
  if err != nil {
    c.JSON(http.StatusNotFound, myerror.ErrorApp{
      Code: http.StatusNotFound,
      Message: "Não existe praga com esse id",
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }

  c.JSON(http.StatusOK, pestResponse)
}

func(p *PestController) PostPestController(c *gin.Context) {

  var requestPest requests.PestRequest

  if err := c.ShouldBindJSON(&requestPest); err != nil {
    c.JSON(http.StatusBadRequest, myerror.ErrorApp{
      Code: http.StatusBadRequest,
      Message: err.Error(),
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }

  val, err := utils.ValidateFieldErrors422UnprocessableEntity(requestPest)
  if err != nil {
    log.Print(err)
    return
  }

  if len(val) > 0 {
    c.JSON(http.StatusUnprocessableEntity, myerror.ErrorApp{
      Code: http.StatusUnprocessableEntity,
      Message: val,
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }

  if err := p.pestService.PostPest(requestPest); err != nil {
    c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
      Code: http.StatusInternalServerError,
      Message: err.Error(),
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }

  c.Status(http.StatusCreated)
}

func(p *PestController) PutPestController(c *gin.Context) {

  val, exists := c.Get("validatedID")
  if !exists {
    return
  }

  id := val.(uint)

  var requesPest requests.PestRequest

  if err := c.ShouldBindJSON(&requesPest); err != nil {
    c.JSON(http.StatusBadRequest, myerror.ErrorApp{
      Code: http.StatusBadRequest,
      Message: err.Error(),
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }

  valid, err := utils.ValidateFieldErrors422UnprocessableEntity(requesPest)
  if err != nil {
    log.Print(err)
    return
  }

  if len(valid) > 0 {
    c.JSON(http.StatusUnprocessableEntity, myerror.ErrorApp{
      Code: http.StatusUnprocessableEntity,
      Message: val,
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }

  if err := p.pestService.PutPest(id, requesPest); err != nil {
    log.Print(err)
    c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
      Code: http.StatusInternalServerError,
      Message: "Erro interno no servidor",
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }

  c.Status(http.StatusOK)
}

func(p *PestController) DeletePestController(c *gin.Context){

  val, exists := c.Get("validatedID")
  if !exists{
    return
  }

  id := val.(uint)

  if err := p.pestService.DeletePest(id); err != nil {
    c.JSON(http.StatusNotFound, myerror.ErrorApp{
      Code: http.StatusNotFound,
      Message: err.Error(),
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }

  c.Status(http.StatusNoContent)
}
