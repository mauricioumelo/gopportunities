package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializa router com as configurações defaults do gin
	r := gin.Default()
	
	// Definindo as rotos
	initializeRoutes(r)

	// Rodando a api
	r.Run(":8080")
}
