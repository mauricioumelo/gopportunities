package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricioumelo/gopportunities/schemas"
)

func CreateOpeningHandler(ctx *gin.Context) {
	req := CreateOpeningRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		logger.Errorf("Validation error: %v", err.Error())
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     req.Role,
		Company:  req.Company,
		Location: req.Location,
		Remote:   *req.Remote,
		Link:     req.Link,
		Salary:   req.Salary,
	}

	err := db.Create(&opening)

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

	SendSuccess(ctx, http.StatusCreated, "Success opening create", resp)

}
