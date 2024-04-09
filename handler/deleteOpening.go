package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricioumelo/gopportunities/schemas"
)

func DeleteOpeningHandler(ctx *gin.Context) {
	or := ListOpeningRequest{}

	if err := ctx.ShouldBindUri(&or); err != nil {
		SendError(ctx, http.StatusBadRequest, errValueIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, or.ID).Error; err != nil {
		SendError(ctx, http.StatusNotFound, fmt.Errorf("opening with id: %s not found", or.ID).Error())
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		SendError(ctx, http.StatusNotFound, "error delete opening")
		return
	}

	SendSuccess(ctx, http.StatusOK, "Success delete opening", opening)
}
