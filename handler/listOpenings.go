package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricioumelo/gopportunities/schemas"
)

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		SendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	SendSuccess(ctx, http.StatusOK, "Success list openings", &openings)
}
