package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mauricioumelo/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1/")
	opening := v1.Group("opening")
	{
		opening.GET("/", handler.ListOpeningsHandler)

		opening.POST("/", handler.CreateOpeningHandler)

		opening.GET("/:id", handler.ShowOpeningHandler)

		opening.PUT("/:id", handler.UpdateOpeningHandler)

		opening.DELETE("/:id", handler.DeleteOpeningHandler)
	}
}