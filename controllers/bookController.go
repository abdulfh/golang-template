package controllers

import (
	"golangbe/models"
	"golangbe/services"
	"golangbe/utils/constants"
	"golangbe/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookController struct {
	services services.BookServices
}

func InitBookController(services services.BookServices) *bookController {
	return &bookController{services}
}

func (controller *bookController) List(context *gin.Context) {
	response := response.ResponseApi{}

	result, err := controller.services.List()
	if err != nil {
		response.ResponseCode = constants.ERROR_RC511
		response.ResponseDesc = constants.ERROR_RM511
		response.Data = result
		context.JSON(http.StatusOK, response)
		return
	}

	response.ResponseCode = constants.ERROR_RC200
	response.ResponseDesc = constants.ERROR_RM200
	response.Data = result
	context.JSON(http.StatusOK, response)
}

func (controller *bookController) Store(context *gin.Context) {
	dataInput := models.Book{}
	response := response.ResponseApi{}

	err := context.ShouldBindJSON(&dataInput)
	if err != nil {
		response.ResponseCode = constants.ERROR_RC410
		response.ResponseDesc = constants.ERROR_RM410
		context.JSON(http.StatusOK, response)
		return
	}

	result, err := controller.services.Store(dataInput)
	if err != nil {
		response.ResponseCode = constants.ERROR_RC511
		response.ResponseDesc = constants.ERROR_RM511
		response.Data = result
		context.JSON(http.StatusOK, response)
		return
	}

	response.ResponseCode = constants.ERROR_RC200
	response.ResponseDesc = constants.ERROR_RM200
	response.Data = result
	context.JSON(http.StatusOK, response)
}
