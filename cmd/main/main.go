package main

import (
	"github.com/lucasbarroso23/gopportunities/config"
	"github.com/lucasbarroso23/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.NewLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
