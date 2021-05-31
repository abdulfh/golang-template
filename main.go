package main

import (
	"fmt"
	"golangbe/config"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.DatabaseOpen()

	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.GET("/")
	}

	err = router.Run(":8080")

	if err != nil {
		fmt.Println("Error while running server", err.Error())
	}

}
