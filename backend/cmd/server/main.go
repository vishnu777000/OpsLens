package main

import (
	"github.com/gin-gonic/gin"

	"github.com/vishnu777000/OpsLens/backend/internal/config"
	"github.com/vishnu777000/OpsLens/backend/internal/routes"
)

func main() {
	cfg := config.Load()

	router := gin.Default()

	routes.RegisterRoutes(router)

	if err := router.Run(":" + cfg.Port); err != nil {
		panic(err)
	}
}
