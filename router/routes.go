package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mauricioumelo/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	opening := v1.Group("opening/")
	{
		opening.GET("/", handler.ListOpeningsHandler)

		opening.POST("/", handler.CreateOpeningHandler)

		opening.GET("/show", handler.ShowOpeningHandler)

		opening.PUT("/", handler.UpdateOpeningHandler)

		opening.DELETE("/", handler.DeleteOpeningHandler)
	}
}
