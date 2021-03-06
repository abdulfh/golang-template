package routers

import (
	"fmt"
	"golangbe/config"
	"golangbe/controllers"
	"golangbe/docs"
	"golangbe/repository"
	"golangbe/services"
	"golangbe/utils/helper"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type GoRouter struct {
	GinFunc gin.HandlerFunc
	Router  *gin.Engine
}

// Server ..
func Server(listenAddress string) (err error) {
	debugMode := helper.GetEnv("APPS_DEBUG", "debug")
	gin.SetMode(debugMode)
	hrisRouter := GoRouter{}
	hrisRouter.Routers()

	err = http.ListenAndServe(listenAddress, hrisRouter.Router)

	if err != nil {
		fmt.Println("Error : ", err)
		return err
	}
	fmt.Println("Routing successfully: ", listenAddress)

	return err
}

func (goRouter *GoRouter) Routers() {
	db, _ := config.DatabaseOpen()

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	bookRepository := repository.InitBookRepository(db)
	bookServices := services.InitBookServices(bookRepository)
	bookController := controllers.InitBookController(bookServices)

	docs.SwaggerInfo.Title = "Golang Template Swagger"
	docs.SwaggerInfo.Description = "This is a list of sample api for Golang Template."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", helper.GetEnv("SWAGGER_HOST", "localhost"), helper.GetEnv("SERVER_PORT", ":40001"))

	api := router.Group("/api/v1")
	{
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		api.GET("/book", bookController.List)
		api.POST("/book", bookController.Store)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	goRouter.Router = router
}
