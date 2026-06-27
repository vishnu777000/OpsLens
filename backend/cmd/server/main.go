package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnu777000/OpsLens/backend/internal/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterHealthRoutes(router)

	router.Run(":8080")
}
