package handler

import (
	"github.com/mauricioumelo/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("HANDLER")
	db = config.GetSQLite()
}
