package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
		"status":  false,
	})
}

func SendSuccess(ctx *gin.Context, code int, message string, v interface{} )  {
	ctx.Header("Content-type", "application/json")
	
	if code <= 0 {
		code = http.StatusOK
	}

	ctx.JSON(code, gin.H{
		"data": v,
		"message": message,
		"status":  true,
	})
}