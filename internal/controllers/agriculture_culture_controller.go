package controllers

import (
	"net/http"
	"time"

	"github.com/ericsanto/apiAgroPlusUltraV1/internal/models/requests"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/services"
	myerror "github.com/ericsanto/apiAgroPlusUltraV1/myError"
	"github.com/gin-gonic/gin"
)


type AgricultureCultureController struct {

  agricultureCultureService *services.AgricultureCultureService
}


func NewAgricultureController(agricultureCultureService *services.AgricultureCultureService) *AgricultureCultureController {

  return &AgricultureCultureController{agricultureCultureService:agricultureCultureService}
}


func(a *AgricultureCultureController) GetAllAgriculturesCultures(c *gin.Context) {

  agricultureCultures, err := a.agricultureCultureService.FindAllAgricultureCulture()

  if err != nil {
    c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
      Code: http.StatusInternalServerError,
      Message: err.Error(),
      Timestamp: time.Now().Format(time.RFC3339),
    })
  }

  c.JSON(http.StatusOK, agricultureCultures)
}

func(a *AgricultureCultureController) PostAgricultureCulture(c *gin.Context) {

  var agricultureCultureRequest requests.AgricultureCultureRequest

  if err := c.ShouldBindJSON(&agricultureCultureRequest); err != nil {
    c.JSON(http.StatusBadRequest, myerror.ErrorApp{
      Code: http.StatusBadRequest,
      Message: err.Error(),
      Timestamp: time.Now().Format(time.RFC3339),
    })
    return
  }



  if err := a.agricultureCultureService.CreateAgricultureCulture(agricultureCultureRequest); err != nil {
    c.JSON(http.StatusInternalServerError, myerror.ErrorApp{
      Code: http.StatusInternalServerError,
      Message: err.Error(),
      Timestamp: time.Now().Format(time.RFC3339),
    }) 
    return
  }

  c.Status(http.StatusCreated)

}
