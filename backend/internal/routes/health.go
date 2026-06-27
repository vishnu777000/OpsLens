// internal/routes/health.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishnu777000/OpsLens/backend/internal/handlers"
)

func RegisterHealthRoutes(router *gin.RouterGroup) {
	router.GET("/health", handlers.HealthHandler)
}
