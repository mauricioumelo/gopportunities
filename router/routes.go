package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mauricioumelo/gopportunities/handler"
	docs "github.com/mauricioumelo/gopportunities/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler
	handler.InitializeHandler()
	
	basePath := "/api/v1/"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)
	opening := v1.Group("opening")
	{
		opening.GET("/", handler.ListOpeningsHandler)

		opening.POST("/", handler.CreateOpeningHandler)

		opening.GET("/:id", handler.ShowOpeningHandler)

		opening.PUT("/:id", handler.UpdateOpeningHandler)

		opening.DELETE("/:id", handler.DeleteOpeningHandler)
	}

	//Initialize Swagger 
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}