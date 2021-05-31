package controllers

import "golangbe/services"

type bookController struct {
	services services.BookServices
}

func InitBookController(services services.BookServices) *bookController {
	return &bookController{services}
}
