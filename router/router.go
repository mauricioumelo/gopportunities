package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa router com as configurações defaults do gin
	r := gin.Default()
	// Definindo as rotos
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Rodando a api
	r.Run(":8080")
}
