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

// @Summary Fetch All Book.
// @Tags Books
// @Produce json
// @Success 200 {object} response.ResponseApi
// @Router /api/v1/book [get]
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

// @Summary Create New Book
// @Tags Books
// @Accept  json
// @Produce  json
// @Param user body models.BookRequest true "new book"
// @Success 200 {object} response.ResponseApi
// @Router /api/v1/book [post]
func (controller *bookController) Store(context *gin.Context) {
	dataRequest := models.BookRequest{}
	response := response.ResponseApi{}

	err := context.ShouldBindJSON(&dataRequest)
	if err != nil {
		response.ResponseCode = constants.ERROR_RC410
		response.ResponseDesc = constants.ERROR_RM410
		context.JSON(http.StatusOK, response)
		return
	}

	dataInput := models.Book{}
	dataInput.Title = dataRequest.Title
	dataInput.Description = dataRequest.Description

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
	response.Data = dataInput
	context.JSON(http.StatusOK, response)
}
