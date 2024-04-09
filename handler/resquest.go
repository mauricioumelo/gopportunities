package handler

import "fmt"

func errValueIsRequired(name string, typ string) error {
	return fmt.Errorf("params: %s (type: %s) is required", name, typ)
}

type CreateOpeningRequest struct {
	Role     string `json:"role" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Remote   *bool  `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
}
type ListOpeningRequest struct {
	ID string `uri:"id" binding:"required"`
}
type UpdateOpeningRequest struct {
	Role     string `json:"role" binding:"required"`
	Company  string `json:"company" binding:"required"`
	Location string `json:"location" binding:"required"`
	Remote   *bool  `json:"remote" binding:"required"`
	Link     string `json:"link" binding:"required"`
	Salary   int64  `json:"salary" binding:"required"`
}
