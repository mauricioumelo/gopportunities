package main

import (
	"github.com/mauricioumelo/gopportunities/config"
	"github.com/mauricioumelo/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("MAIN")
	//Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
