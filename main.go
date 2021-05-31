package main

import (
	"fmt"
	"golangbe/config"
	"golangbe/controllers"
	"golangbe/repository"
	"golangbe/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := config.DatabaseOpen()

	bookRepository := repository.InitBookRepository(db)
	bookServices := services.InitBookServices(bookRepository)
	bookController := controllers.InitBookController(bookServices)

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/book", bookController.List)
		api.POST("/book", bookController.Store)
	}

	err := router.Run(":8080")

	if err != nil {
		fmt.Println("Error while running server", err.Error())
	}

}
