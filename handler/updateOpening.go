package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricioumelo/gopportunities/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	req := UpdateOpeningRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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

	opening.Role = req.Role
	opening.Company = req.Company
	opening.Location = req.Location
	opening.Remote = *req.Remote
	opening.Link = req.Link
	opening.Salary = req.Salary

	err := db.Save(&opening)

	if err.Error != nil {
		logger.Errorf("error creating opening: %v", err.Error)
		SendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}
	resp := schemas.OpeningResponse{
		ID:        opening.ID,
		Role:      opening.Role,
		Company:   opening.Company,
		Location:  opening.Location,
		Remote:    opening.Remote,
		Link:      opening.Link,
		Salary:    opening.Salary,
		CreatedAt: opening.CreatedAt,
		UpdatedAt: opening.UpdatedAt,
	}

	if opening.DeletedAt.Valid {
		resp.DeletedAt = &opening.DeletedAt.Time
	}

	SendSuccess(ctx, http.StatusCreated, "Success opening updated", resp)
}
