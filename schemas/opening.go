package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
type OpeningRequest struct {
	ID   string `uri:"id" binding:"required"`
	Role string `uri:"role"`
}

type OpeningResponse struct {
	ID        uint       `json:"id"`
	Role      string     `json:"opening-role"`
	Company   string     `json:"company"`
	Location  string     `json:"location"`
	Remote    bool       `json:"remote"`
	Link      string     `json:"link"`
	Salary    int64      `json:"salary"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
